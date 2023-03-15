package executor

import (
	"context"
	"encoding/hex"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/avast/retry-go/v4"
	sdkclient "github.com/bnb-chain/greenfield-go-sdk/client/chain"
	sdkkeys "github.com/bnb-chain/greenfield-go-sdk/keys"
	sdktypes "github.com/bnb-chain/greenfield-go-sdk/types"
	relayercommon "github.com/bnb-chain/greenfield-relayer/common"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/logging"
	"github.com/bnb-chain/greenfield-relayer/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	crosschaintypes "github.com/cosmos/cosmos-sdk/x/crosschain/types"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	"github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/votepool"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type GreenfieldExecutor struct {
	BscExecutor *BSCExecutor
	gnfdClients *sdkclient.GnfdCompositeClients
	config      *config.Config
	address     string
	validators  []*tmtypes.Validator // used to cache validators
	cdc         *codec.ProtoCodec
}

func NewGreenfieldExecutor(cfg *config.Config) *GreenfieldExecutor {
	privKey := getGreenfieldPrivateKey(&cfg.GreenfieldConfig)

	km, err := sdkkeys.NewPrivateKeyManager(privKey)
	if err != nil {
		panic(err)
	}

	clients := sdkclient.NewGnfdCompositClients(
		cfg.GreenfieldConfig.GRPCAddrs,
		cfg.GreenfieldConfig.RPCAddrs,
		cfg.GreenfieldConfig.ChainIdString,
		sdkclient.WithKeyManager(km),
		sdkclient.WithGrpcDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	return &GreenfieldExecutor{
		gnfdClients: clients,
		address:     km.GetAddr().String(),
		config:      cfg,
		cdc:         Cdc(),
	}
}

func (e *GreenfieldExecutor) SetBSCExecutor(be *BSCExecutor) {
	e.BscExecutor = be
}

func getGreenfieldPrivateKey(cfg *config.GreenfieldConfig) string {
	var privateKey string
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
		privateKey = awsPrivateKey.PrivateKey
	} else {
		privateKey = cfg.PrivateKey
	}
	return privateKey
}

func (e *GreenfieldExecutor) getRpcClient() (client.Client, error) {
	client := e.gnfdClients.GetClient()
	return client.TendermintClient.RpcClient.TmClient, nil
}

func (e *GreenfieldExecutor) getGnfdClient() (*sdkclient.GreenfieldClient, error) {
	client := e.gnfdClients.GetClient()
	return client.GreenfieldClient, nil
}

func (e *GreenfieldExecutor) GetBlockAndBlockResultAtHeight(height int64) (*tmtypes.Block, *ctypes.ResultBlockResults, error) {
	client, err := e.getRpcClient()
	if err != nil {
		return nil, nil, err
	}
	block, err := client.Block(context.Background(), &height)
	if err != nil {
		return nil, nil, err
	}
	blockResults, err := client.BlockResults(context.Background(), &height)
	if err != nil {
		return nil, nil, err
	}
	return block.Block, blockResults, nil
}

func (e *GreenfieldExecutor) GetLatestBlockHeight() (latestHeight uint64, err error) {
	client := e.gnfdClients.GetClient()
	return uint64(client.Height), nil
}

func (e *GreenfieldExecutor) QueryTendermintLightBlock(height int64) ([]byte, error) {
	client, err := e.getRpcClient()
	if err != nil {
		return nil, err
	}
	validators, err := client.Validators(context.Background(), &height, nil, nil)
	if err != nil {
		return nil, err
	}
	commit, err := client.Commit(context.Background(), &height)
	if err != nil {
		return nil, err
	}
	validatorSet := tmtypes.NewValidatorSet(validators.Validators)
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
			logging.Logger.Infof("failed to query sequence for channel %d, attempt: %d times, max_attempts: %d", channelID, n+1, relayercommon.RtyAttNum)
		}))
}

func (e *GreenfieldExecutor) getNextDeliverySequenceForChannel(channelID types.ChannelId) (uint64, error) {
	sequence, err := e.BscExecutor.GetNextReceiveSequenceForChannelWithRetry(channelID)
	if err != nil {
		return 0, err
	}
	return sequence, nil
}

func (e *GreenfieldExecutor) GetNextReceiveOracleSequence() (uint64, error) {
	gnfdClient, err := e.getGnfdClient()
	if err != nil {
		return 0, err
	}
	res, err := gnfdClient.CrosschainQueryClient.ReceiveSequence(
		context.Background(),
		&crosschaintypes.QueryReceiveSequenceRequest{ChannelId: uint32(relayercommon.OracleChannelId)},
	)
	if err != nil {
		return 0, err
	}
	return res.Sequence, nil
}

// GetNextReceiveSequenceForChannel gets the sequence specifically for bsc -> gnfd package's channel
func (e *GreenfieldExecutor) GetNextReceiveSequenceForChannel(channelId types.ChannelId) (uint64, error) {
	gnfdClient, err := e.getGnfdClient()
	if err != nil {
		return 0, err
	}
	res, err := gnfdClient.ReceiveSequence(
		context.Background(),
		&crosschaintypes.QueryReceiveSequenceRequest{ChannelId: uint32(channelId)},
	)
	if err != nil {
		return 0, err
	}
	return res.Sequence, nil
}

func (e *GreenfieldExecutor) queryLatestValidators() ([]*tmtypes.Validator, error) {
	client, err := e.getRpcClient()
	if err != nil {
		return nil, err
	}
	validators, err := client.Validators(context.Background(), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return validators.Validators, nil
}

func (e *GreenfieldExecutor) QueryValidatorsAtHeight(height uint64) ([]*tmtypes.Validator, error) {
	client, err := e.getRpcClient()
	if err != nil {
		return nil, err
	}
	h := int64(height)
	validators, err := client.Validators(context.Background(), &h, nil, nil)
	if err != nil {
		return nil, err
	}
	return validators.Validators, nil
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
		keys = append(keys, hex.EncodeToString(v.RelayerBlsKey))
	}
	return keys, nil
}

func (e *GreenfieldExecutor) GetNonce() (uint64, error) {
	gnfdClient, err := e.getGnfdClient()
	if err != nil {
		return 0, err
	}
	return gnfdClient.GetNonce()
}

func (e *GreenfieldExecutor) ClaimPackages(payloadBts []byte, aggregatedSig []byte, voteAddressSet []uint64, claimTs int64, oracleSeq uint64, nonce uint64) (string, error) {
	gnfdClient, err := e.getGnfdClient()
	if err != nil {
		return "", err
	}
	msgClaim := oracletypes.NewMsgClaim(
		e.address,
		e.getSrcChainId(),
		e.getDestChainId(),
		oracleSeq,
		uint64(claimTs),
		payloadBts,
		voteAddressSet,
		aggregatedSig,
	)
	txRes, err := gnfdClient.BroadcastTx(
		[]sdk.Msg{msgClaim},
		&sdktypes.TxOption{
			Nonce: nonce,
		},
	)
	if err != nil {
		return "", err
	}
	if txRes.TxResponse.Code != 0 {
		return "", fmt.Errorf("claim error, code=%d, log=%s", txRes.TxResponse.Code, txRes.TxResponse.RawLog)
	}
	return txRes.TxResponse.TxHash, nil
}

func (e *GreenfieldExecutor) GetInturnRelayer() (*oracletypes.QueryInturnRelayerResponse, error) {
	gnfdClient, err := e.getGnfdClient()
	if err != nil {
		return nil, err
	}
	return gnfdClient.OracleQueryClient.InturnRelayer(context.Background(), &oracletypes.QueryInturnRelayerRequest{})
}

func (e *GreenfieldExecutor) QueryVotesByEventHashAndType(eventHash []byte, eventType votepool.EventType) ([]*votepool.Vote, error) {
	client := e.gnfdClients.GetClient()
	queryMap := make(map[string]interface{})
	queryMap[VotePoolQueryParameterEventType] = int(eventType)
	queryMap[VotePoolQueryParameterEventHash] = eventHash
	var queryVote ctypes.ResultQueryVote
	_, err := client.JsonRpcClient.Call(context.Background(), VotePoolQueryMethodName, queryMap, &queryVote)
	if err != nil {
		return nil, err
	}
	return queryVote.Votes, nil
}

func (e *GreenfieldExecutor) BroadcastVote(v *votepool.Vote) error {
	client := e.gnfdClients.GetClient()
	broadcastMap := make(map[string]interface{})
	broadcastMap[VotePoolBroadcastParameterKey] = *v
	_, err := client.JsonRpcClient.Call(context.Background(), VotePoolBroadcastMethodName, broadcastMap, &ctypes.ResultBroadcastVote{})
	if err != nil {
		return err
	}
	return nil
}

func (e *GreenfieldExecutor) getDestChainId() uint32 {
	return uint32(e.config.GreenfieldConfig.ChainId)
}

func (e *GreenfieldExecutor) getSrcChainId() uint32 {
	return uint32(e.config.BSCConfig.ChainId)
}
