package executor

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/avast/retry-go/v4"
	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/executor/crosschain"
	"github.com/bnb-chain/inscription-relayer/executor/tendermintlightclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BSCClient struct {
	BSCClient *ethclient.Client
	Provider  string
	Height    uint64
	UpdatedAt time.Time
}

type BSCExecutor struct {
	gasPriceMutex       sync.RWMutex
	mutex               sync.RWMutex
	InscriptionExecutor *InscriptionExecutor
	daoManager          *dao.DaoManager
	clientIdx           int
	bscClients          []*BSCClient
	config              *config.Config
	privateKey          *ecdsa.PrivateKey
	TxSender            common.Address
	gasPrice            *big.Int
}

func initBSCClients(providers []string) []*BSCClient {
	bscClients := make([]*BSCClient, 0)
	for _, provider := range providers {
		client, err := ethclient.Dial(provider)
		if err != nil {
			panic("new eth client error")
		}
		bscClients = append(bscClients, &BSCClient{
			BSCClient: client,
			Provider:  provider,
			UpdatedAt: time.Now(),
		})
	}

	return bscClients
}

func getBscPrivateKey(cfg *config.BSCConfig) (*ecdsa.PrivateKey, error) {
	var privateKey string
	if cfg.KeyType == config.KeyTypeAWSPrivateKey {
		result, err := config.GetSecret(cfg.AWSSecretName, cfg.AWSRegion)
		if err != nil {
			return nil, err
		}
		type AwsPrivateKey struct {
			PrivateKey string `json:"private_key"`
		}
		var awsPrivateKey AwsPrivateKey
		err = json.Unmarshal([]byte(result), &awsPrivateKey)
		if err != nil {
			return nil, err
		}
		privateKey = awsPrivateKey.PrivateKey
	} else {
		privateKey = cfg.PrivateKey
	}

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

func NewBSCExecutor(cfg *config.Config, dao *dao.DaoManager) (*BSCExecutor, error) {
	privKey, err := getBscPrivateKey(&cfg.BSCConfig)
	if err != nil {
		return nil, err
	}
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("get public key error")
	}
	txSender := crypto.PubkeyToAddress(*publicKeyECDSA)

	var initGasPrice *big.Int
	if cfg.BSCConfig.GasPrice == 0 {
		initGasPrice = big.NewInt(DefaultGasPrice)
	} else {
		initGasPrice = big.NewInt(int64(cfg.BSCConfig.GasPrice))
	}

	return &BSCExecutor{
		daoManager: dao,
		clientIdx:  0,
		bscClients: initBSCClients(cfg.BSCConfig.RPCAddrs),
		privateKey: privKey,
		TxSender:   txSender,
		config:     cfg,
		gasPrice:   initGasPrice,
	}, nil
}

func (e *BSCExecutor) SetInscriptionExecutor(insE *InscriptionExecutor) {
	e.InscriptionExecutor = insE
}

func (e *BSCExecutor) GetClient() *ethclient.Client {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.bscClients[e.clientIdx].BSCClient
}

func (e *BSCExecutor) SwitchClient() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.clientIdx++
	if e.clientIdx >= len(e.bscClients) {
		e.clientIdx = 0
	}
	relayercommon.Logger.Infof("Switch to provider: %s", e.config.BSCConfig.RPCAddrs[e.clientIdx])
}

func (e *BSCExecutor) GetLatestBlockHeightWithRetry() (latestHeight uint64, err error) {
	return e.getLatestBlockHeightWithRetry(e.GetClient())
}

func (e *BSCExecutor) getLatestBlockHeightWithRetry(client *ethclient.Client) (latestHeight uint64, err error) {
	return latestHeight, retry.Do(func() error {
		latestHeight, err = e.GetLatestBlockHeight(client)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			relayercommon.Logger.Infof("Failed to query latest height, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (e *BSCExecutor) GetLatestBlockHeight(client *ethclient.Client) (uint64, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	block, err := client.BlockByNumber(ctxWithTimeout, nil)
	if err != nil {
		return 0, err
	}
	return block.Number().Uint64(), nil
}

func (e *BSCExecutor) UpdateClients() {
	for {
		relayercommon.Logger.Infof("Start to monitor bsc data-seeds healthy")
		for _, bscClient := range e.bscClients {
			if time.Since(bscClient.UpdatedAt).Seconds() > DataSeedDenyServiceThreshold {
				msg := fmt.Sprintf("data seed %s is not accessable", bscClient.Provider)
				relayercommon.Logger.Error(msg)
				config.SendTelegramMessage(e.config.AlertConfig.Identity, e.config.AlertConfig.TelegramBotId,
					e.config.AlertConfig.TelegramChatId, msg)
			}
			height, err := e.GetLatestBlockHeight(bscClient.BSCClient)
			if err != nil {
				relayercommon.Logger.Errorf("get latest block height error, err=%s", err.Error())
				continue
			}
			bscClient.Height = height
			bscClient.UpdatedAt = time.Now()
		}

		highestHeight := uint64(0)
		highestIdx := 0
		for idx := 0; idx < len(e.bscClients); idx++ {
			if e.bscClients[idx].Height > highestHeight {
				highestHeight = e.bscClients[idx].Height
				highestIdx = idx
			}
		}
		// current client block sync is fall behind, switch to the client with the highest block height
		if e.bscClients[e.clientIdx].Height+FallBehindThreshold < highestHeight {
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
	header, err := e.GetClient().HeaderByNumber(ctxWithTimeout, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return header, nil
}

func (e *BSCExecutor) GetNextSequence(channelID relayercommon.ChannelId) (uint64, error) {
	crossChainInstance, err := crosschain.NewCrosschain(crossChainContractAddr, e.GetClient())
	if err != nil {
		return 0, err
	}

	callOpts, err := e.getCallOpts()
	if err != nil {
		return 0, err
	}

	return crossChainInstance.ChannelReceiveSequenceMap(callOpts, uint8(channelID))
}

func (e *BSCExecutor) GetNextDeliveryOracleSequence() (uint64, error) {
	sequence, err := e.InscriptionExecutor.GetNextOracleSequence()
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
	txOpts.GasPrice = e.GetGasPrice()
	return txOpts, nil
}

func (e *BSCExecutor) GetGasPrice() *big.Int {
	e.gasPriceMutex.RLock()
	defer e.gasPriceMutex.RUnlock()
	return e.gasPrice
}

func (e *BSCExecutor) getCallOpts() (*bind.CallOpts, error) {
	callOpts := &bind.CallOpts{
		Pending: true,
		Context: context.Background(),
	}
	return callOpts, nil
}

func (e *BSCExecutor) SyncTendermintLightClientHeader(height uint64) (common.Hash, error) {
	nonce, err := e.GetClient().PendingNonceAt(context.Background(), e.TxSender)
	if err != nil {
		return common.Hash{}, err
	}
	txOpts, err := e.getTransactor(nonce)
	if err != nil {
		return common.Hash{}, err
	}

	instance, err := tendermintlightclient.NewTendermintlightclient(tendermintLightClientContractAddr, e.GetClient())
	if err != nil {
		return common.Hash{}, err
	}

tryAgain:
	header, err := e.InscriptionExecutor.QueryTendermintHeader(int64(height))
	if err != nil {
		if isHeaderNonExistingErr(err) {
			goto tryAgain
		} else {
			return common.Hash{}, err
		}
	}

	headerBytes, err := header.SignedHeader.ToProto().Marshal()
	if err != nil {
		return common.Hash{}, err
	}
	tx, err := instance.SyncTendermintHeader(txOpts, headerBytes, height, header.BlsPubKeys, header.Relayers)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (e *BSCExecutor) CallBuildInSystemContract(channelID int8, blsSignature []byte, sequence uint64, validatorSet *big.Int,
	msgBytes []byte, nonce uint64) (common.Hash, error) {

	txOpts, err := e.getTransactor(nonce)
	if err != nil {
		return common.Hash{}, err
	}

	crossChainInstance, err := crosschain.NewCrosschain(crossChainContractAddr, e.GetClient())
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := crossChainInstance.HandlePackage(txOpts, msgBytes, blsSignature, validatorSet, sequence, uint8(channelID))
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}
