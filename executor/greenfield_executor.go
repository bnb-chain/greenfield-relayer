package executor

import (
	"context"
	"encoding/hex"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/bnb-chain/greenfield-relayer/types"
	"sync"
	"time"

	"github.com/bnb-chain/greenfield-relayer/logging"

	"github.com/avast/retry-go/v4"
	clitx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	crosschaintypes "github.com/cosmos/cosmos-sdk/x/crosschain/types"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	libclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
	tmtypes "github.com/tendermint/tendermint/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	relayercommon "github.com/bnb-chain/greenfield-relayer/common"
	"github.com/bnb-chain/greenfield-relayer/config"
)

type GreenfieldClient struct {
	rpcClient        rpcclient.Client
	txClient         tx.ServiceClient
	authClient       authtypes.QueryClient
	crossChainClient crosschaintypes.QueryClient
	Provider         string
	Height           uint64
	UpdatedAt        time.Time
}

type GreenfieldExecutor struct {
	mutex             sync.RWMutex
	BscExecutor       *BSCExecutor
	clientIdx         int
	greenfieldClients []*GreenfieldClient
	config            *config.Config
	privateKey        *ethsecp256k1.PrivKey
	address           string
	validators        []*tmtypes.Validator // used to cache validators
}

func grpcConn(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	return conn
}

func NewRpcClient(addr string) *http.HTTP {
	httpClient, err := libclient.DefaultHTTPClient(addr)
	if err != nil {
		panic(err)
	}
	rpcClient, err := http.NewWithClient(addr, "/websocket", httpClient)
	if err != nil {
		panic(err)
	}
	return rpcClient
}

func getGreenfieldPrivateKey(cfg *config.GreenfieldConfig) *ethsecp256k1.PrivKey {
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
	privKey, err := HexToEthSecp256k1PrivKey(privateKey)
	if err != nil {
		panic(err)
	}
	return privKey
}

func initGreenfieldClients(rpcAddrs, grpcAddrs []string) []*GreenfieldClient {
	greenfieldClients := make([]*GreenfieldClient, 0)

	for i := 0; i < len(rpcAddrs); i++ {
		conn := grpcConn(grpcAddrs[i])
		greenfieldClients = append(greenfieldClients, &GreenfieldClient{
			txClient:         tx.NewServiceClient(conn),
			authClient:       authtypes.NewQueryClient(conn),
			crossChainClient: crosschaintypes.NewQueryClient(conn),
			rpcClient:        NewRpcClient(rpcAddrs[i]),
			Provider:         rpcAddrs[i],
			UpdatedAt:        time.Now(),
		})
	}
	return greenfieldClients
}

func NewGreenfieldExecutor(cfg *config.Config) *GreenfieldExecutor {
	privKey := getGreenfieldPrivateKey(&cfg.GreenfieldConfig)
	return &GreenfieldExecutor{
		clientIdx:         0,
		greenfieldClients: initGreenfieldClients(cfg.GreenfieldConfig.RPCAddrs, cfg.GreenfieldConfig.GRPCAddrs),
		privateKey:        privKey,
		address:           privKey.PubKey().Address().String(),
		config:            cfg,
	}
}

func (e *GreenfieldExecutor) SetBSCExecutor(be *BSCExecutor) {
	e.BscExecutor = be
}

func (e *GreenfieldExecutor) getRpcClient() rpcclient.Client {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.greenfieldClients[e.clientIdx].rpcClient
}

func (e *GreenfieldExecutor) getTxClient() tx.ServiceClient {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.greenfieldClients[e.clientIdx].txClient
}

func (e *GreenfieldExecutor) getAuthClient() authtypes.QueryClient {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.greenfieldClients[e.clientIdx].authClient
}

func (e *GreenfieldExecutor) getCrossChainClient() crosschaintypes.QueryClient {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.greenfieldClients[e.clientIdx].crossChainClient
}

func (e *GreenfieldExecutor) GetBlockResultAtHeight(height int64) (*ctypes.ResultBlockResults, error) {
	blockResults, err := e.getRpcClient().BlockResults(context.Background(), &height)
	if err != nil {
		return nil, err
	}
	return blockResults, nil
}

func (e *GreenfieldExecutor) GetBlockAtHeight(height int64) (*tmtypes.Block, error) {
	block, err := e.getRpcClient().Block(context.Background(), &height)
	if err != nil {
		return nil, err
	}
	return block.Block, nil
}

func (e *GreenfieldExecutor) GetLatestBlockHeightWithRetry() (latestHeight uint64, err error) {
	return e.getLatestBlockHeightWithRetry(e.getRpcClient())
}

func (e *GreenfieldExecutor) getLatestBlockHeightWithRetry(client rpcclient.Client) (latestHeight uint64, err error) {
	return latestHeight, retry.Do(func() error {
		latestHeightQueryCtx, cancelLatestHeightQueryCtx := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelLatestHeightQueryCtx()
		var err error
		latestHeight, err = e.getLatestBlockHeight(latestHeightQueryCtx, client)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			logging.Logger.Infof("failed to query latest height, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (e *GreenfieldExecutor) getLatestBlockHeight(ctx context.Context, client rpcclient.Client) (uint64, error) {
	status, err := client.Status(ctx)
	if err != nil {
		return 0, err
	}
	return uint64(status.SyncInfo.LatestBlockHeight), nil
}

func (e *GreenfieldExecutor) UpdateClientLoop() {
	ticker := time.NewTicker(SleepSecondForUpdateClient * time.Second)
	for range ticker.C {
		logging.Logger.Infof("start to monitor greenfield data-seeds healthy")
		for _, greenfieldClient := range e.greenfieldClients {
			if time.Since(greenfieldClient.UpdatedAt).Seconds() > DataSeedDenyServiceThreshold {
				msg := fmt.Sprintf("data seed %s is not accessable", greenfieldClient.Provider)
				logging.Logger.Error(msg)
				config.SendTelegramMessage(e.config.AlertConfig.Identity, e.config.AlertConfig.TelegramBotId,
					e.config.AlertConfig.TelegramChatId, msg)
			}
			height, err := e.getLatestBlockHeightWithRetry(greenfieldClient.rpcClient)
			if err != nil {
				logging.Logger.Errorf("get latest block height error, err=%s", err.Error())
				continue
			}
			greenfieldClient.Height = height
			greenfieldClient.UpdatedAt = time.Now()
		}
		highestHeight := uint64(0)
		highestIdx := 0
		for idx := 0; idx < len(e.greenfieldClients); idx++ {
			if e.greenfieldClients[idx].Height > highestHeight {
				highestHeight = e.greenfieldClients[idx].Height
				highestIdx = idx
			}
		}
		// current GreenfieldClient block sync is fall behind, switch to the GreenfieldClient with the highest block height
		if e.greenfieldClients[e.clientIdx].Height+FallBehindThreshold < highestHeight {
			e.mutex.Lock()
			e.clientIdx = highestIdx
			e.mutex.Unlock()
		}
	}
}

func (e *GreenfieldExecutor) QueryTendermintLightBlock(height int64) ([]byte, error) {
	validators, err := e.getRpcClient().Validators(context.Background(), &height, nil, nil)
	commit, err := e.getRpcClient().Commit(context.Background(), &height)
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

// GetNextDeliverySequenceForChannel call dest chain(BSC) to return a sequence# which should be used.
func (e *GreenfieldExecutor) GetNextDeliverySequenceForChannel(channelID types.ChannelId) (uint64, error) {
	sequence, err := e.BscExecutor.GetNextReceiveSequenceForChannel(channelID)
	if err != nil {
		return 0, err
	}
	return sequence, nil
}

func (e *GreenfieldExecutor) GetNextReceiveOracleSequence() (uint64, error) {
	res, err := e.getCrossChainClient().ReceiveSequence(
		context.Background(),
		&crosschaintypes.QueryReceiveSequenceRequest{ChannelId: uint32(relayercommon.OracleChannelId)},
	)
	if err != nil {
		return 0, err
	}
	return res.Sequence, nil
}

// GetNextReceiveSequenceForChannel gets the sequence specifically for cross-chain package's channel
func (e *GreenfieldExecutor) GetNextReceiveSequenceForChannel(channelId types.ChannelId) (uint64, error) {
	res, err := e.getCrossChainClient().ReceiveSequence(
		context.Background(),
		&crosschaintypes.QueryReceiveSequenceRequest{ChannelId: uint32(channelId)},
	)
	if err != nil {
		return 0, err
	}
	return res.Sequence, nil
}

func (e *GreenfieldExecutor) queryLatestValidators() ([]*tmtypes.Validator, error) {
	validators, err := e.getRpcClient().Validators(context.Background(), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return validators.Validators, nil
}

func (e *GreenfieldExecutor) QueryValidatorsAtHeight(height uint64) ([]*tmtypes.Validator, error) {
	atHeight := int64(height)
	validators, err := e.getRpcClient().Validators(context.Background(), &atHeight, nil, nil)
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

func (e *GreenfieldExecutor) GetAccount(address string) (authtypes.AccountI, error) {
	authRes, err := e.getAuthClient().Account(context.Background(), &authtypes.QueryAccountRequest{Address: address})
	if err != nil {
		return nil, err
	}
	var account authtypes.AccountI
	if err := Cdc().InterfaceRegistry().UnpackAny(authRes.Account, &account); err != nil {
		return nil, err
	}
	return account, nil
}

func (e *GreenfieldExecutor) ClaimPackages(payloadBts []byte, aggregatedSig []byte, voteAddressSet []uint64, claimTs int64) (string, error) {
	txConfig := authtx.NewTxConfig(Cdc(), authtx.DefaultSignModes)
	txBuilder := txConfig.NewTxBuilder()
	seq, err := e.GetNextReceiveOracleSequence()
	if err != nil {
		return "", err
	}
	msgClaim := &oracletypes.MsgClaim{}
	msgClaim.FromAddress = e.address
	msgClaim.Payload = payloadBts
	msgClaim.VoteAddressSet = voteAddressSet
	msgClaim.Sequence = seq
	msgClaim.AggSignature = aggregatedSig
	msgClaim.DestChainId = e.getDestChainId()
	msgClaim.SrcChainId = e.getSrcChainId()
	msgClaim.Timestamp = uint64(claimTs)
	err = txBuilder.SetMsgs(msgClaim)

	if err != nil {
		return "", err
	}
	txBuilder.SetGasLimit(e.config.GreenfieldConfig.GasLimit)

	acct, err := e.GetAccount(e.address)
	if err != nil {
		return "", err
	}
	accountNum := acct.GetAccountNumber()
	accountSeq := acct.GetSequence()

	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	sig := signing.SignatureV2{
		PubKey: e.privateKey.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  signing.SignMode_SIGN_MODE_EIP_712,
			Signature: nil,
		},
		Sequence: accountSeq,
	}

	err = txBuilder.SetSignatures(sig)
	if err != nil {
		return "", err
	}

	// Second round: all signer infos are set, so each signer can sign.
	sig = signing.SignatureV2{}

	signerData := xauthsigning.SignerData{
		ChainID:       e.config.GreenfieldConfig.ChainIdString,
		AccountNumber: accountNum,
		Sequence:      accountSeq,
	}

	sig, err = clitx.SignWithPrivKey(signing.SignMode_SIGN_MODE_EIP_712,
		signerData,
		txBuilder,
		e.privateKey,
		txConfig,
		accountSeq,
	)
	if err != nil {
		return "", err
	}

	err = txBuilder.SetSignatures(sig)
	if err != nil {
		return "", err
	}

	txBytes, err := txConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return "", err
	}
	// Broadcast transaction
	txRes, err := e.getTxClient().BroadcastTx(
		context.Background(),
		&tx.BroadcastTxRequest{
			Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		})
	if err != nil {
		return "", err
	}
	if txRes.TxResponse.Code != 0 {
		return "", fmt.Errorf("claim error, code=%d, log=%s", txRes.TxResponse.Code, txRes.TxResponse.RawLog)
	}
	return txRes.TxResponse.TxHash, nil
}

func (e *GreenfieldExecutor) getDestChainId() uint32 {
	return uint32(e.config.GreenfieldConfig.ChainId)
}

func (e *GreenfieldExecutor) getSrcChainId() uint32 {
	return uint32(e.config.BSCConfig.ChainId)
}
