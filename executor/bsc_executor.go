package executor

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"

	relayercommon "github.com/bnb-chain/greenfield-relayer/common"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/executor/crosschain"
	"github.com/bnb-chain/greenfield-relayer/executor/greenfieldlightclient"
	"github.com/bnb-chain/greenfield-relayer/logging"
	rtypes "github.com/bnb-chain/greenfield-relayer/types"
)

type BSCClient struct {
	rpcClient             *rpc.Client // for eth_getFinalizedBlock usage, supported by BSC
	ethClient             *ethclient.Client
	crossChainClient      *crosschain.Crosschain
	greenfieldLightClient *greenfieldlightclient.Greenfieldlightclient
	provider              string
	height                uint64
	updatedAt             time.Time
}

func NewBSCClients(config *config.Config) []*BSCClient {
	bscClients := make([]*BSCClient, 0)
	for _, provider := range config.BSCConfig.RPCAddrs {
		rpcClient, err := rpc.DialContext(context.Background(), provider)
		if err != nil {
			panic("new rpc client error")
		}
		ethClient, err := ethclient.Dial(provider)
		if err != nil {
			panic("new eth client error")
		}
		greenfieldLightClient, err := greenfieldlightclient.NewGreenfieldlightclient(
			common.HexToAddress(config.RelayConfig.GreenfieldLightClientContractAddr),
			ethClient)
		if err != nil {
			panic("new crossChain client error")
		}
		crossChainClient, err := crosschain.NewCrosschain(
			common.HexToAddress(config.RelayConfig.CrossChainContractAddr),
			ethClient)
		if err != nil {
			panic("new greenfield light client error")
		}
		bscClients = append(bscClients, &BSCClient{
			rpcClient:             rpcClient,
			ethClient:             ethClient,
			crossChainClient:      crossChainClient,
			greenfieldLightClient: greenfieldLightClient,
			provider:              provider,
			updatedAt:             time.Now(),
		})
	}
	return bscClients
}

type BSCExecutor struct {
	gasPriceMutex      sync.RWMutex
	mutex              sync.RWMutex
	GreenfieldExecutor *GreenfieldExecutor
	clientIdx          int
	bscClients         []*BSCClient
	config             *config.Config
	privateKey         *ecdsa.PrivateKey
	txSender           common.Address
	gasPrice           *big.Int
	relayers           []rtypes.Validator // cached relayers
}

func getBscPrivateKey(cfg *config.BSCConfig) string {
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

func NewBSCExecutor(cfg *config.Config) *BSCExecutor {
	privKey := viper.GetString(config.FlagConfigPrivateKey)
	if privKey == "" {
		privKey = getBscPrivateKey(&cfg.BSCConfig)
	}

	ecdsaPrivKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		panic(err)
	}
	publicKey := ecdsaPrivKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("get public key error")
	}
	txSender := crypto.PubkeyToAddress(*publicKeyECDSA)
	var initGasPrice *big.Int
	if cfg.BSCConfig.GasPrice == 0 {
		initGasPrice = big.NewInt(DefaultGasPrice)
	} else {
		initGasPrice = big.NewInt(int64(cfg.BSCConfig.GasPrice))
	}
	return &BSCExecutor{
		clientIdx:  0,
		bscClients: NewBSCClients(cfg),
		privateKey: ecdsaPrivKey,
		txSender:   txSender,
		config:     cfg,
		gasPrice:   initGasPrice,
	}
}

func (e *BSCExecutor) SetGreenfieldExecutor(ge *GreenfieldExecutor) {
	e.GreenfieldExecutor = ge
}

func (e *BSCExecutor) GetRpcClient() *rpc.Client {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.bscClients[e.clientIdx].rpcClient
}

func (e *BSCExecutor) GetEthClient() *ethclient.Client {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.bscClients[e.clientIdx].ethClient
}

func (e *BSCExecutor) getCrossChainClient() *crosschain.Crosschain {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.bscClients[e.clientIdx].crossChainClient
}

func (e *BSCExecutor) getGreenfieldLightClient() *greenfieldlightclient.Greenfieldlightclient {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.bscClients[e.clientIdx].greenfieldLightClient
}

func (e *BSCExecutor) SwitchClient() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.clientIdx++
	if e.clientIdx >= len(e.bscClients) {
		e.clientIdx = 0
	}
	logging.Logger.Infof("switch to provider: %s", e.config.BSCConfig.RPCAddrs[e.clientIdx])
}

func (e *BSCExecutor) GetLatestFinalizedBlockHeightWithRetry() (latestHeight uint64, err error) {
	return e.getLatestBlockHeightWithRetry(e.GetEthClient(), e.GetRpcClient(), true)
}

func (e *BSCExecutor) GetLatestBlockHeightWithRetry() (latestHeight uint64, err error) {
	return e.getLatestBlockHeightWithRetry(e.GetEthClient(), e.GetRpcClient(), false)
}

func (e *BSCExecutor) getLatestBlockHeightWithRetry(ethClient *ethclient.Client, rpcClient *rpc.Client, finalized bool) (latestHeight uint64, err error) {
	return latestHeight, retry.Do(func() error {
		latestHeight, err = e.getLatestBlockHeight(ethClient, rpcClient, finalized)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			logging.Logger.Errorf("failed to query latest height, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (e *BSCExecutor) getLatestBlockHeight(client *ethclient.Client, rpcClient *rpc.Client, finalized bool) (uint64, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	if finalized {
		return e.getFinalizedBlockHeight(ctxWithTimeout, rpcClient)
	}
	header, err := client.HeaderByNumber(ctxWithTimeout, nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

func (e *BSCExecutor) UpdateClientLoop() {
	ticker := time.NewTicker(SleepSecondForUpdateClient * time.Second)
	for range ticker.C {
		logging.Logger.Infof("start to monitor bsc data-seeds healthy")
		for _, bscClient := range e.bscClients {
			if time.Since(bscClient.updatedAt).Seconds() > DataSeedDenyServiceThreshold {
				msg := fmt.Sprintf("data seed %s is not accessable", bscClient.provider)
				logging.Logger.Error(msg)
				config.SendTelegramMessage(e.config.AlertConfig.Identity, e.config.AlertConfig.TelegramBotId,
					e.config.AlertConfig.TelegramChatId, msg)
			}
			height, err := e.getLatestBlockHeight(bscClient.ethClient, bscClient.rpcClient, true)
			if err != nil {
				logging.Logger.Errorf("get latest block height error, err=%s", err.Error())
				continue
			}
			bscClient.height = height
			bscClient.updatedAt = time.Now()
		}

		highestHeight := uint64(0)
		highestIdx := 0
		for idx := 0; idx < len(e.bscClients); idx++ {
			if e.bscClients[idx].height > highestHeight {
				highestHeight = e.bscClients[idx].height
				highestIdx = idx
			}
		}
		// current client block sync is fall behind, switch to the client with the highest block height
		if e.bscClients[e.clientIdx].height+FallBehindThreshold < highestHeight {
			e.mutex.Lock()
			e.clientIdx = highestIdx
			e.mutex.Unlock()
		}
	}
}

func (e *BSCExecutor) GetBlockHeaderAtHeight(height uint64) (*types.Header, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	header, err := e.GetEthClient().HeaderByNumber(ctx, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return header, nil
}

// GetNextReceiveSequenceForChannelWithRetry gets the next receive sequence for specified channel from BSC
func (e *BSCExecutor) GetNextReceiveSequenceForChannelWithRetry(channelID rtypes.ChannelId) (sequence uint64, err error) {
	return sequence, retry.Do(func() error {
		sequence, err = e.getNextReceiveSequenceForChannel(channelID)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			logging.Logger.Errorf("failed to query receive sequence for channel %d, attempt: %d times, max_attempts: %d", channelID, n+1, relayercommon.RtyAttNum)
		}))
}

func (e *BSCExecutor) getNextReceiveSequenceForChannel(channelID rtypes.ChannelId) (sequence uint64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	callOpts := &bind.CallOpts{
		Pending: true,
		Context: ctx,
	}
	return e.getCrossChainClient().ChannelReceiveSequenceMap(callOpts, uint8(channelID))
}

// GetNextSendSequenceForChannelWithRetry gets the next send oracle sequence from  BSC
func (e *BSCExecutor) GetNextSendSequenceForChannelWithRetry() (sequence uint64, err error) {
	return sequence, retry.Do(func() error {
		sequence, err = e.getNextSendOracleSequence()
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			logging.Logger.Errorf("failed to query send oracle sequence, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (e *BSCExecutor) getNextSendOracleSequence() (sequence uint64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	callOpts := &bind.CallOpts{
		Pending: true,
		Context: ctx,
	}
	sentOracleSeq, err := e.getCrossChainClient().OracleSequence(callOpts)
	if err != nil {
		return 0, err
	}
	return uint64(sentOracleSeq + 1), nil
}

// GetNextDeliveryOracleSequenceWithRetry gets the next delivery Oracle sequence from Greenfield
func (e *BSCExecutor) GetNextDeliveryOracleSequenceWithRetry() (sequence uint64, err error) {
	return sequence, retry.Do(func() error {
		sequence, err = e.getNextDeliveryOracleSequence()
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			logging.Logger.Errorf("failed to query oracle sequence, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (e *BSCExecutor) getNextDeliveryOracleSequence() (uint64, error) {
	sequence, err := e.GreenfieldExecutor.GetNextReceiveOracleSequence()
	if err != nil {
		return 0, err
	}
	return sequence, nil
}

func (e *BSCExecutor) getTransactor(nonce uint64) (*bind.TransactOpts, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(e.privateKey, big.NewInt(int64(e.config.BSCConfig.ChainId)))
	if err != nil {
		return nil, err
	}
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(0)
	txOpts.GasLimit = e.config.BSCConfig.GasLimit
	txOpts.GasPrice = e.getGasPrice()
	return txOpts, nil
}

func (e *BSCExecutor) getGasPrice() *big.Int {
	e.gasPriceMutex.RLock()
	defer e.gasPriceMutex.RUnlock()
	return e.gasPrice
}

func (e *BSCExecutor) SyncTendermintLightBlock(height uint64) (common.Hash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	nonce, err := e.GetEthClient().PendingNonceAt(ctx, e.txSender)
	if err != nil {
		return common.Hash{}, err
	}
	txOpts, err := e.getTransactor(nonce)
	if err != nil {
		return common.Hash{}, err
	}
	lightBlock, err := e.QueryTendermintLightBlockWithRetry(int64(height))
	if err != nil {
		return common.Hash{}, err
	}
	tx, err := e.getGreenfieldLightClient().SyncLightBlock(txOpts, lightBlock, height)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (e *BSCExecutor) QueryTendermintLightBlockWithRetry(height int64) (lightBlock []byte, err error) {
	return lightBlock, retry.Do(func() error {
		lightBlock, err = e.GreenfieldExecutor.QueryTendermintLightBlock(height)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			logging.Logger.Errorf("failed to query tendermint header, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (e *BSCExecutor) QueryLatestTendermintHeaderWithRetry() (lightBlock []byte, err error) {
	latestHeigh, err := e.GreenfieldExecutor.GetLatestBlockHeight()
	if err != nil {
		return nil, err
	}
	return lightBlock, retry.Do(func() error {
		lightBlock, err = e.GreenfieldExecutor.QueryTendermintLightBlock(int64(latestHeigh))
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			logging.Logger.Errorf("failed to query tendermint header, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (e *BSCExecutor) GetNonce() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	return e.GetEthClient().PendingNonceAt(ctx, e.txSender)
}

func (e *BSCExecutor) CallBuildInSystemContract(blsSignature []byte, validatorSet *big.Int, msgBytes []byte, nonce uint64) (common.Hash, error) {
	txOpts, err := e.getTransactor(nonce)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := e.getCrossChainClient().HandlePackage(txOpts, msgBytes, blsSignature, validatorSet)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

// QueryLatestValidators used for gnfd -> bsc
func (e *BSCExecutor) QueryLatestValidators() ([]rtypes.Validator, error) {
	relayerAddresses, err := e.getGreenfieldLightClient().GetRelayers(nil)
	if err != nil {
		return nil, err
	}
	blsKeys, err := e.getGreenfieldLightClient().BlsPubKeys(nil)
	if err != nil {
		return nil, err
	}
	relayers := make([]rtypes.Validator, len(relayerAddresses))
	nextRelayerBtsStartIdx := 0

	for i, addr := range relayerAddresses {
		r := rtypes.Validator{
			RelayerAddress: addr,
			BlsPublicKey:   blsKeys[nextRelayerBtsStartIdx : nextRelayerBtsStartIdx+RelayerBytesLength][:],
		}
		nextRelayerBtsStartIdx = nextRelayerBtsStartIdx + RelayerBytesLength
		relayers[i] = r
	}
	return relayers, nil
}

// QueryCachedLatestValidators Used for gnfd -> bsc
func (e *BSCExecutor) QueryCachedLatestValidators() ([]rtypes.Validator, error) {
	if len(e.relayers) != 0 {
		return e.relayers, nil
	}
	relayers, err := e.QueryLatestValidators()
	if err != nil {
		return nil, err
	}
	return relayers, nil
}

func (e *BSCExecutor) UpdateCachedLatestValidatorsLoop() {
	ticker := time.NewTicker(UpdateCachedValidatorsInterval)
	for range ticker.C {
		relayers, err := e.QueryLatestValidators()
		if err != nil {
			logging.Logger.Errorf("update latest bsc relayers error, err=%s", err)
			continue
		}
		e.relayers = relayers
	}
}

func (e *BSCExecutor) GetLightClientLatestHeight() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	callOpts := &bind.CallOpts{
		Pending: true,
		Context: ctx,
	}
	latestHeight, err := e.getGreenfieldLightClient().GnfdHeight(callOpts)
	if err != nil {
		return 0, err
	}
	return latestHeight, err
}

func (e *BSCExecutor) GetValidatorsBlsPublicKey() ([]string, error) {
	validators, err := e.QueryCachedLatestValidators()
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, v := range validators {
		keys = append(keys, hex.EncodeToString(v.BlsPublicKey[:]))
	}
	return keys, nil
}

func (e *BSCExecutor) GetInturnRelayer() (*rtypes.InturnRelayer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	callOpts := &bind.CallOpts{
		Pending: true,
		Context: ctx,
	}
	r, err := e.getGreenfieldLightClient().GetInturnRelayer(callOpts)
	if err != nil {
		return nil, err
	}

	return &rtypes.InturnRelayer{
		BlsPublicKey: hex.EncodeToString(r.BlsKey),
		Start:        r.Start.Uint64(),
		End:          r.End.Uint64(),
	}, nil
}

func (e *BSCExecutor) getFinalizedBlockHeight(ctx context.Context, rpcClient *rpc.Client) (uint64, error) {
	var head *types.Header
	err := rpcClient.CallContext(ctx, &head, "eth_getFinalizedHeader", e.config.BSCConfig.NumberOfBlocksForFinality)
	if err == nil && head == nil {
		return 0, ethereum.NotFound
	}
	return head.Number.Uint64(), nil
}
