package executor

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/avast/retry-go/v4"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cometbft/cometbft/votepool"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/blst"
	"github.com/spf13/viper"

	sdktypes "github.com/bnb-chain/greenfield-go-sdk/types"
	relayercommon "github.com/bnb-chain/greenfield-relayer/common"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/logging"
	"github.com/bnb-chain/greenfield-relayer/types"
	gnfdsdktypes "github.com/bnb-chain/greenfield/sdk/types"
)

type GreenfieldExecutor struct {
	BscExecutor   *BSCExecutor
	gnfdClients   GnfdCompositeClients
	config        *config.Config
	address       string
	validators    []*tmtypes.Validator // used to cache validators
	BlsPrivateKey []byte
	BlsPubKey     []byte
}

func NewGreenfieldExecutor(cfg *config.Config) *GreenfieldExecutor {
	privKey := viper.GetString(config.FlagConfigPrivateKey)
	if privKey == "" {
		privKey = getGreenfieldPrivateKey(&cfg.GreenfieldConfig)
	}
	blsPrivKeyStr := viper.GetString(config.FlagConfigBlsPrivateKey)
	if blsPrivKeyStr == "" {
		blsPrivKeyStr = getGreenfieldBlsPrivateKey(&cfg.GreenfieldConfig)
	}
	blsPrivKeyBts := ethcommon.Hex2Bytes(blsPrivKeyStr)

	blsPrivKey, err := blst.SecretKeyFromBytes(blsPrivKeyBts)
	if err != nil {
		panic(err)
	}
	account, err := sdktypes.NewAccountFromPrivateKey("relayer", privKey)
	if err != nil {
		panic(err)
	}
	clients := NewGnfdCompositClients(
		cfg.GreenfieldConfig.RPCAddrs,
		cfg.GreenfieldConfig.ChainIdString,
		account,
		cfg.GreenfieldConfig.UseWebsocket,
	)
	return &GreenfieldExecutor{
		gnfdClients:   clients,
		address:       account.GetAddress().String(),
		config:        cfg,
		BlsPrivateKey: blsPrivKeyBts,
		BlsPubKey:     blsPrivKey.PublicKey().Marshal(),
	}
}

func (e *GreenfieldExecutor) SetBSCExecutor(be *BSCExecutor) {
	e.BscExecutor = be
}

func getGreenfieldPrivateKey(cfg *config.GreenfieldConfig) string {
	if cfg.KeyType == config.KeyTypeAWSPrivateKey {
		result, err := config.GetSecret(cfg.AWSSecretName, cfg.AWSRegion)
		if err != nil {
			panic(err)
		}
		type AwsPrivateKey struct {
			PrivateKey string `json:"private_key"`
		}
		var awsPrivateKey AwsPrivateKey
		err = json.Unmarshal([]byte(result), &awsPrivateKey)
		if err != nil {
			panic(err)
		}
		return awsPrivateKey.PrivateKey
	}
	return cfg.PrivateKey
}

func getGreenfieldBlsPrivateKey(cfg *config.GreenfieldConfig) string {
	if cfg.KeyType == config.KeyTypeAWSPrivateKey {
		result, err := config.GetSecret(cfg.AWSBlsSecretName, cfg.AWSRegion)
		if err != nil {
			panic(err)
		}
		type AwsPrivateKey struct {
			PrivateKey string `json:"bls_private_key"`
		}
		var awsBlsPrivateKey AwsPrivateKey
		err = json.Unmarshal([]byte(result), &awsBlsPrivateKey)
		if err != nil {
			panic(err)
		}
		return awsBlsPrivateKey.PrivateKey
	}
	return cfg.BlsPrivateKey
}

func (e *GreenfieldExecutor) GetGnfdClient() *GreenfieldClient {
	return e.gnfdClients.GetClient()
}

func (e *GreenfieldExecutor) GetBlockAndBlockResultAtHeight(height int64) (*tmtypes.Block, *ctypes.ResultBlockResults, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	block, err := e.GetGnfdClient().GetBlockByHeight(ctx, height)
	if err != nil {
		return nil, nil, err
	}
	blockResults, err := e.GetGnfdClient().GetBlockResultByHeight(ctx, height)
	if err != nil {
		return nil, nil, err
	}
	return block, blockResults, nil
}

func (e *GreenfieldExecutor) GetLatestBlockHeight() (latestHeight uint64, err error) {
	return uint64(e.gnfdClients.GetClient().Height), nil
}

func (e *GreenfieldExecutor) QueryTendermintLightBlock(height int64) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	validators, err := e.GetGnfdClient().GetValidatorsByHeight(ctx, height)
	if err != nil {
		return nil, err
	}
	commit, err := e.GetGnfdClient().GetCommit(ctx, height)
	if err != nil {
		return nil, err
	}
	validatorSet := tmtypes.NewValidatorSet(validators)
	if err != nil {
		return nil, err
	}
	lightBlock := tmtypes.LightBlock{
		SignedHeader: &commit.SignedHeader,
		ValidatorSet: validatorSet,
	}
	protoBlock, err := lightBlock.ToProto()
	if err != nil {
		return nil, err
	}
	return protoBlock.Marshal()
}

// GetNextDeliverySequenceForChannelWithRetry calls dest chain(BSC) to return a sequence # which should be used.
func (e *GreenfieldExecutor) GetNextDeliverySequenceForChannelWithRetry(channelID types.ChannelId) (sequence uint64, err error) {
	return sequence, retry.Do(func() error {
		sequence, err = e.getNextDeliverySequenceForChannel(channelID)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			logging.Logger.Errorf("failed to query receive sequence for channel %d, attempt: %d times, max_attempts: %d", channelID, n+1, relayercommon.RtyAttNum)
		}))
}

func (e *GreenfieldExecutor) getNextDeliverySequenceForChannel(channelID types.ChannelId) (uint64, error) {
	sequence, err := e.BscExecutor.GetNextReceiveSequenceForChannelWithRetry(channelID)
	if err != nil {
		return 0, err
	}
	return sequence, nil
}

// GetNextSendSequenceForChannelWithRetry gets the next send sequence of a specified channel from Greenfield
func (e *GreenfieldExecutor) GetNextSendSequenceForChannelWithRetry(destChainID sdk.ChainID, channelID types.ChannelId) (sequence uint64, err error) {
	return sequence, retry.Do(func() error {
		sequence, err = e.getNextSendSequenceForChannel(destChainID, channelID)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			logging.Logger.Errorf("failed to query send sequence for channel %d, attempt: %d times, max_attempts: %d", channelID, n+1, relayercommon.RtyAttNum)
		}))
}

func (e *GreenfieldExecutor) getNextSendSequenceForChannel(destChainId sdk.ChainID, channelId types.ChannelId) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	return e.GetGnfdClient().GetChannelSendSequence(
		ctx,
		destChainId,
		uint32(channelId),
	)
}

// GetNextReceiveOracleSequence gets the next receive Oracle sequence from Greenfield
func (e *GreenfieldExecutor) GetNextReceiveOracleSequence(destChainId sdk.ChainID) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	return e.GetGnfdClient().GetChannelReceiveSequence(
		ctx,
		destChainId,
		uint32(relayercommon.OracleChannelId),
	)
}

// GetNextReceiveSequenceForChannel gets the sequence specifically for bsc -> gnfd package's channel from Greenfield
func (e *GreenfieldExecutor) GetNextReceiveSequenceForChannel(destChainId sdk.ChainID, channelId types.ChannelId) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	return e.GetGnfdClient().GetChannelReceiveSequence(
		ctx,
		destChainId,
		uint32(channelId),
	)
}

func (e *GreenfieldExecutor) queryLatestValidators() ([]*tmtypes.Validator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	_, validators, err := e.GetGnfdClient().GetValidatorSet(ctx)
	if err != nil {
		return nil, err
	}
	return validators, nil
}

func (e *GreenfieldExecutor) QueryValidatorsAtHeight(height uint64) ([]*tmtypes.Validator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	return e.GetGnfdClient().GetValidatorsByHeight(ctx, int64(height))
}

func (e *GreenfieldExecutor) QueryCachedLatestValidators() ([]*tmtypes.Validator, error) {
	if len(e.validators) != 0 {
		return e.validators, nil
	}
	validators, err := e.queryLatestValidators()
	if err != nil {
		return nil, err
	}
	return validators, nil
}

func (e *GreenfieldExecutor) UpdateCachedLatestValidatorsLoop() {
	ticker := time.NewTicker(UpdateCachedValidatorsInterval)
	for range ticker.C {
		validators, err := e.queryLatestValidators()
		if err != nil {
			logging.Logger.Errorf("update latest greenfield validators error, err=%s", err)
			continue
		}
		e.validators = validators
	}
}

func (e *GreenfieldExecutor) GetValidatorsBlsPublicKey() ([]string, error) {
	validators, err := e.QueryCachedLatestValidators()
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, v := range validators {
		keys = append(keys, hex.EncodeToString(v.BlsKey))
	}
	return keys, nil
}

func (e *GreenfieldExecutor) GetNonce() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	acc, err := e.GetGnfdClient().GetAccount(ctx, e.address)
	if err != nil {
		return 0, err
	}
	return acc.GetSequence(), nil
}

func (e *GreenfieldExecutor) GetNonceOnNextBlock() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	err := e.GetGnfdClient().WaitForNextBlock(ctx)
	if err != nil {
		return 0, err
	}
	return e.GetNonce()
}

func (e *GreenfieldExecutor) ClaimPackages(client *GreenfieldClient, payloadBts []byte, aggregatedSig []byte, voteAddressSet []uint64, claimTs int64, oracleSeq uint64, nonce uint64) (string, error) {
	msg := oracletypes.NewMsgClaim(
		e.address,
		e.getSrcChainId(),
		e.getDestChainId(),
		oracleSeq,
		uint64(claimTs),
		payloadBts,
		voteAddressSet,
		aggregatedSig)
	gasLimit, feeAmount, err := e.getGasLimitAndFeeAmount(msg)
	if err != nil {
		return "", err
	}
	txOpt := gnfdsdktypes.TxOption{
		NoSimulate: true,
		GasLimit:   uint64(gasLimit),
		FeeAmount:  sdk.NewCoins(sdk.NewCoin(gnfdsdktypes.Denom, sdk.NewInt(feeAmount))),
		Nonce:      nonce,
	}
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	resp, err := client.BroadcastTx(ctx, []sdk.Msg{msg}, &txOpt)
	if err != nil {
		return "", err
	}
	txRes := resp.TxResponse
	if txRes.Codespace == oracletypes.ModuleName && txRes.Code == oracletypes.ErrInvalidReceiveSequence.ABCICode() {
		return "", oracletypes.ErrInvalidReceiveSequence
	}

	if txRes.Codespace == sdkErrors.RootCodespace && txRes.Code == sdkErrors.ErrWrongSequence.ABCICode() {
		return "", sdkErrors.ErrWrongSequence
	}

	if txRes.Codespace == sdkErrors.RootCodespace && txRes.Code == sdkErrors.ErrTxInMempoolCache.ABCICode() {
		return "", sdkErrors.ErrTxInMempoolCache
	}

	if txRes.Code != 0 {
		return "", fmt.Errorf("claim error, code=%d, log=%s", txRes.Code, txRes.RawLog)
	}
	return txRes.TxHash, nil
}

func (e *GreenfieldExecutor) GetInturnRelayer(srcChain oracletypes.ClaimSrcChain) (*oracletypes.QueryInturnRelayerResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	return e.GetGnfdClient().GetInturnRelayer(ctx, &oracletypes.QueryInturnRelayerRequest{
		ClaimSrcChain: srcChain,
	})
}

func (e *GreenfieldExecutor) QueryVotesByEventHashAndType(eventHash []byte, eventType votepool.EventType) ([]*votepool.Vote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	votes, err := e.gnfdClients.GetClient().QueryVote(ctx, int(eventType), eventHash)
	if err != nil {
		return nil, err
	}
	return votes.Votes, nil
}

func (e *GreenfieldExecutor) BroadcastVote(v *votepool.Vote) error {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	return e.gnfdClients.GetClient().BroadcastVote(ctx, *v)
}

func (e *GreenfieldExecutor) getDestChainId() uint32 {
	return uint32(e.config.GreenfieldConfig.ChainId)
}

func (e *GreenfieldExecutor) getSrcChainId() uint32 {
	return uint32(e.config.BSCConfig.ChainId)
}

func (e *GreenfieldExecutor) getGasLimitAndFeeAmount(msg *oracletypes.MsgClaim) (gasLimit int64, feeAmount int64, err error) {
	bz, err := msg.Marshal()
	if err != nil {
		return
	}
	if len(bz)+EstimatedTxExtraMetaSize >= MaxTxSizeForFixGasLimit {
		gasLimit = GasLimitRatio * int64(len(bz)+EstimatedTxExtraMetaSize)
		feeAmount = gasLimit * GnfdGasPrice
		return
	}
	return e.config.GreenfieldConfig.GasLimit, e.config.GreenfieldConfig.FeeAmount, nil
}

func (e *GreenfieldExecutor) GetCrossTxPack(destChainID sdk.ChainID, channelID types.ChannelId, sequence uint64) (pack []byte, err error) {
	return pack, retry.Do(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
		defer cancel()
		pack, err = e.GetGnfdClient().GetCrossChainPackage(ctx, destChainID, uint32(channelID), sequence)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			logging.Logger.Errorf("failed to query crosschain tx for channel %d, seq %d, attempt: %d times, max_attempts: %d", channelID, n+1, relayercommon.RtyAttNum)
		}))
}
