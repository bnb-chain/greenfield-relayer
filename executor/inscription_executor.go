package executor

import (
	"context"
	"encoding/hex"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/bnb-chain/inscription-relayer/util"
	"sync"
	"time"

	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	crosschainypes "github.com/cosmos/cosmos-sdk/x/crosschain/types"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	"google.golang.org/grpc"

	"github.com/avast/retry-go/v4"
	clitx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/tx"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	libclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
	tmtypes "github.com/tendermint/tendermint/types"
)

type InscriptionClient struct {
	rpcClient          rpcclient.Client
	txClient           tx.ServiceClient
	stakingQueryClient stakingtypes.QueryClient
	authClient         authtypes.QueryClient
	crossChainClient   crosschainypes.QueryClient
	Provider           string
	Height             uint64
	UpdatedAt          time.Time
}

type InscriptionExecutor struct {
	mutex              sync.RWMutex
	BscExecutor        *BSCExecutor
	clientIdx          int
	inscriptionClients []*InscriptionClient
	config             *config.Config
	privateKey         *ethsecp256k1.PrivKey
	address            string
	validators         []stakingtypes.Validator // used to cache validators
}

func grpcConn(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	return conn
}

func NewRpcClient(addr string) *rpchttp.HTTP {
	httpClient, err := libclient.DefaultHTTPClient(addr)
	if err != nil {
		panic(err)
	}
	rpcClient, err := rpchttp.NewWithClient(addr, "/websocket", httpClient)
	if err != nil {
		panic(err)
	}
	return rpcClient
}

func getInscriptionPrivateKey(cfg *config.InscriptionConfig) *ethsecp256k1.PrivKey {
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

func initInscriptionClients(rpcAddrs, grpcAddrs []string) []*InscriptionClient {
	inscriptionClients := make([]*InscriptionClient, 0)

	for i := 0; i < len(rpcAddrs); i++ {
		conn := grpcConn(grpcAddrs[i])
		inscriptionClients = append(inscriptionClients, &InscriptionClient{
			txClient:           tx.NewServiceClient(conn),
			stakingQueryClient: stakingtypes.NewQueryClient(conn),
			authClient:         authtypes.NewQueryClient(conn),
			crossChainClient:   crosschainypes.NewQueryClient(conn),
			rpcClient:          NewRpcClient(rpcAddrs[i]),
			Provider:           rpcAddrs[i],
			UpdatedAt:          time.Now(),
		})
	}
	return inscriptionClients
}

func NewInscriptionExecutor(cfg *config.Config) *InscriptionExecutor {
	privKey := getInscriptionPrivateKey(&cfg.InscriptionConfig)
	return &InscriptionExecutor{
		clientIdx:          0,
		inscriptionClients: initInscriptionClients(cfg.InscriptionConfig.RPCAddrs, cfg.InscriptionConfig.GRPCAddrs),
		privateKey:         privKey,
		address:            privKey.PubKey().Address().String(),
		config:             cfg,
	}
}

func (e *InscriptionExecutor) SetBSCExecutor(bscE *BSCExecutor) {
	e.BscExecutor = bscE
}

func (e *InscriptionExecutor) getRpcClient() rpcclient.Client {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.inscriptionClients[e.clientIdx].rpcClient
}

func (e *InscriptionExecutor) getTxClient() tx.ServiceClient {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.inscriptionClients[e.clientIdx].txClient
}

func (e *InscriptionExecutor) getStakingClient() stakingtypes.QueryClient {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.inscriptionClients[e.clientIdx].stakingQueryClient
}

func (e *InscriptionExecutor) getAuthClient() authtypes.QueryClient {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.inscriptionClients[e.clientIdx].authClient
}

func (e *InscriptionExecutor) getCrossChainClient() crosschainypes.QueryClient {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.inscriptionClients[e.clientIdx].crossChainClient
}

func (e *InscriptionExecutor) GetBlockResultAtHeight(height int64) (*ctypes.ResultBlockResults, error) {
	blockResults, err := e.getRpcClient().BlockResults(context.Background(), &height)
	if err != nil {
		return nil, err
	}
	return blockResults, nil
}

func (e *InscriptionExecutor) GetBlockAtHeight(height int64) (*tmtypes.Block, error) {
	block, err := e.getRpcClient().Block(context.Background(), &height)
	if err != nil {
		return nil, err
	}
	return block.Block, nil
}

func (e *InscriptionExecutor) GetLatestBlockHeightWithRetry() (latestHeight uint64, err error) {
	return e.getLatestBlockHeightWithRetry(e.getRpcClient())
}

func (e *InscriptionExecutor) getLatestBlockHeightWithRetry(client rpcclient.Client) (latestHeight uint64, err error) {
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
			relayercommon.Logger.Infof("failed to query latest height, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (e *InscriptionExecutor) getLatestBlockHeight(ctx context.Context, client rpcclient.Client) (uint64, error) {
	status, err := client.Status(ctx)
	if err != nil {
		return 0, err
	}
	return uint64(status.SyncInfo.LatestBlockHeight), nil
}

func (e *InscriptionExecutor) UpdateClients() {
	for {
		relayercommon.Logger.Infof("start to monitor inscription data-seeds healthy")
		for _, inscriptionClient := range e.inscriptionClients {
			if time.Since(inscriptionClient.UpdatedAt).Seconds() > DataSeedDenyServiceThreshold {
				msg := fmt.Sprintf("data seed %s is not accessable", inscriptionClient.Provider)
				relayercommon.Logger.Error(msg)
				config.SendTelegramMessage(e.config.AlertConfig.Identity, e.config.AlertConfig.TelegramBotId,
					e.config.AlertConfig.TelegramChatId, msg)
			}
			height, err := e.getLatestBlockHeightWithRetry(inscriptionClient.rpcClient)
			if err != nil {
				relayercommon.Logger.Errorf("get latest block height error, err=%s", err.Error())
				continue
			}
			inscriptionClient.Height = height
			inscriptionClient.UpdatedAt = time.Now()
		}
		highestHeight := uint64(0)
		highestIdx := 0
		for idx := 0; idx < len(e.inscriptionClients); idx++ {
			if e.inscriptionClients[idx].Height > highestHeight {
				highestHeight = e.inscriptionClients[idx].Height
				highestIdx = idx
			}
		}
		// current InscriptionClient block sync is fall behind, switch to the InscriptionClient with the highest block height
		if e.inscriptionClients[e.clientIdx].Height+FallBehindThreshold < highestHeight {
			e.mutex.Lock()
			e.clientIdx = highestIdx
			e.mutex.Unlock()
		}
		time.Sleep(SleepSecondForUpdateClient * time.Second)
	}
}

func (e *InscriptionExecutor) QueryTendermintHeader(height int64) (*relayercommon.Header, error) {
	commit, err := e.getRpcClient().Commit(context.Background(), &height)
	if err != nil {
		return nil, err
	}
	validators, err := e.QueryValidatorsAtHeight(uint64(height))
	if err != nil {
		return nil, err
	}

	var blsPubKeysBts []byte
	var relayerAddrs []common.Address
	for _, v := range validators {
		blsPubKeysBts = append(blsPubKeysBts, v.RelayerBlsKey...)
		relayerAddrs = append(relayerAddrs, common.HexToAddress(v.RelayerAddress))
	}

	header := &relayercommon.Header{
		SignedHeader: commit.SignedHeader,
		Height:       uint64(height),
		BlsPubKeys:   blsPubKeysBts,
		Relayers:     relayerAddrs,
	}

	return header, nil
}

// GetNextDeliverySequenceForChannel call dest chain(BSC) to return a sequence# which should be used.
func (e *InscriptionExecutor) GetNextDeliverySequenceForChannel(channelID relayercommon.ChannelId) (uint64, error) {
	sequence, err := e.BscExecutor.GetNextReceiveSequenceForChannel(channelID)
	if err != nil {
		return 0, err
	}
	return sequence, nil
}

func (e *InscriptionExecutor) GetNextReceiveOracleSequence() (uint64, error) {
	res, err := e.getCrossChainClient().ReceiveSequence(
		context.Background(),
		&crosschainypes.QueryReceiveSequenceRequest{ChannelId: uint32(relayercommon.OracleChannelId)},
	)
	if err != nil {
		return 0, err
	}
	return res.Sequence, nil
}

// GetNextReceiveSequenceForChannel gets the sequence specifically for cross-chain package's channel
func (e *InscriptionExecutor) GetNextReceiveSequenceForChannel(channelId relayercommon.ChannelId) (uint64, error) {
	res, err := e.getCrossChainClient().ReceiveSequence(
		context.Background(),
		&crosschainypes.QueryReceiveSequenceRequest{ChannelId: uint32(channelId)},
	)
	if err != nil {
		return 0, err
	}
	return res.Sequence, nil
}

func (e *InscriptionExecutor) queryLatestValidators() ([]stakingtypes.Validator, error) {
	height, err := e.GetLatestBlockHeightWithRetry()
	if err != nil {
		return nil, err
	}

	result, err := e.QueryValidatorsAtHeight(height)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (e *InscriptionExecutor) QueryValidatorsAtHeight(height uint64) ([]stakingtypes.Validator, error) {
	result, err := e.getStakingClient().HistoricalInfo(context.Background(), &stakingtypes.QueryHistoricalInfoRequest{Height: int64(height)})
	if err != nil {
		return nil, err
	}
	relayercommon.Logger.Infof("queried validators from inscription at height %d", height)
	hist := result.Hist
	return hist.Valset, nil
}

func (e *InscriptionExecutor) QueryCachedLatestValidators() ([]stakingtypes.Validator, error) {
	if len(e.validators) != 0 {
		return e.validators, nil
	}
	validators, err := e.queryLatestValidators()
	if err != nil {
		return nil, err
	}
	return validators, nil
}

func (e *InscriptionExecutor) UpdateCachedLatestValidators() {
	ticker := time.NewTicker(UpdateCachedValidatorsInterval)
	for {
		validators, err := e.queryLatestValidators()
		if err != nil {
			relayercommon.Logger.Errorf("update latest inscription validators error, err=%s", err)
			<-ticker.C
			continue
		}
		e.validators = validators
		<-ticker.C
	}
}

func (e *InscriptionExecutor) GetValidatorsBlsPublicKey() ([]string, error) {
	validators, err := e.QueryCachedLatestValidators()
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, v := range validators {
		keys = append(keys, hex.EncodeToString(v.GetRelayerBlsKey()))
	}
	return keys, nil
}

func (e *InscriptionExecutor) GetAccount(address string) (authtypes.AccountI, error) {
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

func (e *InscriptionExecutor) ClaimPackages(payloadBts []byte, aggregatedSig []byte, voteAddressSet []uint64, claimTs int64) (string, error) {
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
	txBuilder.SetGasLimit(e.config.InscriptionConfig.GasLimit)

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
		ChainID:       e.config.InscriptionConfig.ChainIdString,
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

func (e *InscriptionExecutor) getDestChainId() uint32 {
	return uint32(e.config.InscriptionConfig.ChainId)
}

func (e *InscriptionExecutor) getSrcChainId() uint32 {
	return uint32(e.config.BSCConfig.ChainId)
}

func (e *InscriptionExecutor) IsValidator() bool {
	relayerBlsPubKeys, err := e.GetValidatorsBlsPublicKey()
	if err != nil {
		panic(err)
	}
	relayerPubKey := util.GetBlsPubKeyFromPrivKeyStr(e.config.VotePoolConfig.BlsPrivateKey)
	relayerIdx := util.IndexOf(hex.EncodeToString(relayerPubKey), relayerBlsPubKeys)
	return relayerIdx != -1
}
