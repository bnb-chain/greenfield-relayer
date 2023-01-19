package executor

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/executor/crosschain"
	"github.com/bnb-chain/inscription-relayer/executor/inscriptionlightclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BSCClient struct {
	rpcClient              *ethclient.Client
	crossChainClient       *crosschain.Crosschain
	inscriptionLightClient *inscriptionlightclient.Inscriptionlightclient
	provider               string
	height                 uint64
	updatedAt              time.Time
}

type BSCExecutor struct {
	gasPriceMutex       sync.RWMutex
	mutex               sync.RWMutex
	InscriptionExecutor *InscriptionExecutor
	clientIdx           int
	bscClients          []*BSCClient
	config              *config.Config
	privateKey          *ecdsa.PrivateKey
	txSender            common.Address
	gasPrice            *big.Int
	relayers            []Validator // cached relayers
}

func initBSCClients(providers []string) []*BSCClient {
	bscClients := make([]*BSCClient, 0)
	for _, provider := range providers {
		rpcClient, err := ethclient.Dial(provider)
		if err != nil {
			panic("new eth client error")
		}
		inscriptionLightClient, err := inscriptionlightclient.NewInscriptionlightclient(InscriptionLightClientContractAddr, rpcClient)
		if err != nil {
			panic("new crossChain client error")
		}
		crossChainClient, err := crosschain.NewCrosschain(CrossChainContractAddr, rpcClient)
		if err != nil {
			panic("new inscription light client error")
		}
		bscClients = append(bscClients, &BSCClient{
			rpcClient:              rpcClient,
			crossChainClient:       crossChainClient,
			inscriptionLightClient: inscriptionLightClient,
			provider:               provider,
			updatedAt:              time.Now(),
		})
	}
	return bscClients
}

func getBscPrivateKey(cfg *config.BSCConfig) *ecdsa.PrivateKey {
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

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}
	return privKey
}

func NewBSCExecutor(cfg *config.Config) *BSCExecutor {
	privKey := getBscPrivateKey(&cfg.BSCConfig)
	publicKey := privKey.Public()
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
		bscClients: initBSCClients(cfg.BSCConfig.RPCAddrs),
		privateKey: privKey,
		txSender:   txSender,
		config:     cfg,
		gasPrice:   initGasPrice,
	}
}

func (e *BSCExecutor) SetInscriptionExecutor(insE *InscriptionExecutor) {
	e.InscriptionExecutor = insE
}

func (e *BSCExecutor) GetRpcClient() *ethclient.Client {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.bscClients[e.clientIdx].rpcClient
}

func (e *BSCExecutor) getCrossChainClient() *crosschain.Crosschain {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.bscClients[e.clientIdx].crossChainClient
}

func (e *BSCExecutor) getInscriptionLightClient() *inscriptionlightclient.Inscriptionlightclient {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.bscClients[e.clientIdx].inscriptionLightClient
}

func (e *BSCExecutor) SwitchClient() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.clientIdx++
	if e.clientIdx >= len(e.bscClients) {
		e.clientIdx = 0
	}
	relayercommon.Logger.Infof("switch to provider: %s", e.config.BSCConfig.RPCAddrs[e.clientIdx])
}

func (e *BSCExecutor) GetLatestBlockHeightWithRetry() (latestHeight uint64, err error) {
	return e.getLatestBlockHeightWithRetry(e.GetRpcClient())
}

func (e *BSCExecutor) getLatestBlockHeightWithRetry(client *ethclient.Client) (latestHeight uint64, err error) {
	return latestHeight, retry.Do(func() error {
		latestHeight, err = e.getLatestBlockHeight(client)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			relayercommon.Logger.Infof("failed to query latest height, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (e *BSCExecutor) getLatestBlockHeight(client *ethclient.Client) (uint64, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	block, err := client.BlockByNumber(ctxWithTimeout, nil)
	if err != nil {
		return 0, err
	}
	return block.Number().Uint64(), nil
}

func (e *BSCExecutor) UpdateClients() {
	for {
		relayercommon.Logger.Infof("start to monitor bsc data-seeds healthy")
		for _, bscClient := range e.bscClients {
			if time.Since(bscClient.updatedAt).Seconds() > DataSeedDenyServiceThreshold {
				msg := fmt.Sprintf("data seed %s is not accessable", bscClient.provider)
				relayercommon.Logger.Error(msg)
				config.SendTelegramMessage(e.config.AlertConfig.Identity, e.config.AlertConfig.TelegramBotId,
					e.config.AlertConfig.TelegramChatId, msg)
			}
			height, err := e.getLatestBlockHeight(bscClient.rpcClient)
			if err != nil {
				relayercommon.Logger.Errorf("get latest block height error, err=%s", err.Error())
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
		time.Sleep(SleepSecondForUpdateClient * time.Second)
	}
}

func (e *BSCExecutor) GetBlockHeaderAtHeight(height uint64) (*types.Header, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	header, err := e.GetRpcClient().HeaderByNumber(ctxWithTimeout, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return header, nil
}

func (e *BSCExecutor) GetNextReceiveSequenceForChannel(channelID relayercommon.ChannelId) (uint64, error) {
	crossChainInstance, err := crosschain.NewCrosschain(CrossChainContractAddr, e.GetRpcClient())
	if err != nil {
		return 0, err
	}
	callOpts := &bind.CallOpts{
		Pending: true,
		Context: context.Background(),
	}
	if err != nil {
		return 0, err
	}
	return crossChainInstance.ChannelReceiveSequenceMap(callOpts, uint8(channelID))
}

func (e *BSCExecutor) GetNextDeliveryOracleSequence() (uint64, error) {
	sequence, err := e.InscriptionExecutor.GetNextReceiveOracleSequence()
	if err != nil {
		return 0, err
	}
	return sequence, nil
}

func (e *BSCExecutor) getTransactor(nonce uint64) (*bind.TransactOpts, error) {
	txOpts := bind.NewKeyedTransactor(e.privateKey)
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

func (e *BSCExecutor) SyncTendermintLightClientHeader(height uint64) (common.Hash, error) {
	nonce, err := e.GetRpcClient().PendingNonceAt(context.Background(), e.txSender)
	if err != nil {
		return common.Hash{}, err
	}
	txOpts, err := e.getTransactor(nonce)
	if err != nil {
		return common.Hash{}, err
	}
	header, err := e.QueryTendermintHeaderWithRetry(int64(height))
	if err != nil {
		return common.Hash{}, err
	}
	headerBytes, err := header.SignedHeader.ToProto().Marshal()
	if err != nil {
		return common.Hash{}, err
	}
	tx, err := e.getInscriptionLightClient().SyncTendermintHeader(txOpts, headerBytes, height, header.BlsPubKeys, header.Relayers)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (e *BSCExecutor) QueryTendermintHeaderWithRetry(height int64) (header *relayercommon.Header, err error) {
	return header, retry.Do(func() error {
		header, err = e.InscriptionExecutor.QueryTendermintHeader(height)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			relayercommon.Logger.Infof("failed to query tendermint header, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (e *BSCExecutor) QueryLatestTendermintHeaderWithRetry() (header *relayercommon.Header, err error) {
	latestHeigh, err := e.InscriptionExecutor.GetLatestBlockHeightWithRetry()
	if err != nil {
		return nil, err
	}
	return header, retry.Do(func() error {
		header, err = e.InscriptionExecutor.QueryTendermintHeader(int64(latestHeigh))
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			relayercommon.Logger.Infof("failed to query tendermint header, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (e *BSCExecutor) CallBuildInSystemContract(blsSignature []byte, validatorSet *big.Int, msgBytes []byte) (common.Hash, error) {
	fmt.Printf("bls sig hex: %s", hex.EncodeToString(blsSignature))
	fmt.Printf("msg byte to hex : %s", hex.EncodeToString(msgBytes))
	fmt.Printf("validatorSet : %s", validatorSet.String())

	nonce, err := e.GetRpcClient().PendingNonceAt(context.Background(), e.txSender)
	if err != nil {
		return common.Hash{}, err
	}
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

func (e *BSCExecutor) QueryLatestValidators() ([]Validator, error) {
	latestHeight, err := e.GetLatestBlockHeightWithRetry()
	if err != nil {
		return nil, err
	}
	callOpts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(latestHeight)),
		Context:     context.Background(),
	}
	relayerAddresses, err := e.getInscriptionLightClient().GetRelayers(callOpts)
	if err != nil {
		return nil, err
	}
	blsKeys, err := e.getInscriptionLightClient().BlsPubKeys(callOpts)
	if err != nil {
		return nil, err
	}

	relayercommon.Logger.Infof("Queried relayers from BSC at height %d", latestHeight)

	relayers := make([]Validator, len(relayerAddresses))
	nextRelayerBtsStartIdx := 0

	for i, addr := range relayerAddresses {
		r := Validator{
			RelayerAddress: addr,
			BlsPublicKey:   blsKeys[nextRelayerBtsStartIdx : nextRelayerBtsStartIdx+RelayerBytesLength][:],
		}
		nextRelayerBtsStartIdx = nextRelayerBtsStartIdx + RelayerBytesLength
		relayers[i] = r
	}
	return relayers, nil
}

func (e *BSCExecutor) QueryCachedLatestValidators() ([]Validator, error) {
	if len(e.relayers) != 0 {
		return e.relayers, nil
	}
	relayers, err := e.QueryLatestValidators()
	if err != nil {
		return nil, err
	}
	return relayers, nil
}

func (e *BSCExecutor) UpdateCachedLatestValidators() {
	ticker := time.NewTicker(UpdateCachedValidatorsInterval)
	for {
		relayers, err := e.QueryLatestValidators()
		if err != nil {
			relayercommon.Logger.Errorf("update latest bsc relayers error, err=%s", err)
			<-ticker.C
			continue
		}
		e.relayers = relayers
		<-ticker.C
	}
}

func (e *BSCExecutor) GetLightClientLatestHeight() (uint64, error) {
	callOpts := &bind.CallOpts{
		Pending: true,
		Context: context.Background(),
	}
	latestHeight, err := e.getInscriptionLightClient().InsHeight(callOpts)
	if err != nil {
		return 0, err
	}
	return latestHeight, err
}

func (e *BSCExecutor) GetValidatorsBlsPublicKey() ([]string, error) {
	validators, err := e.QueryLatestValidators()
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, v := range validators {
		keys = append(keys, hex.EncodeToString(v.BlsPublicKey[:]))
	}
	return keys, nil
}
