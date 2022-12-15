package executor

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum"
	ethereumcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	relayercommon "inscription-relayer/common"
	"inscription-relayer/config"
	"inscription-relayer/db/dao"
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
	mutex               sync.RWMutex
	daoManager          *dao.DaoManager
	InscriptionExecutor *InscriptionExecutor
	clientIdx           int
	bscClients          []*BSCClient
	config              *config.Config
	privateKey          *ecdsa.PrivateKey
	TxSender            common.Address
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

	return &BSCExecutor{
		daoManager: dao,
		clientIdx:  0,
		bscClients: initBSCClients(cfg.BSCConfig.RPCAddrs),

		privateKey: privKey,
		TxSender:   txSender,
		config:     cfg,
	}, nil
}

func (executor *BSCExecutor) SetInscriptionExecutor(e *InscriptionExecutor) {
	executor.InscriptionExecutor = e
}

func (executor *BSCExecutor) GetClient() *ethclient.Client {
	executor.mutex.RLock()
	defer executor.mutex.RUnlock()
	return executor.bscClients[executor.clientIdx].BSCClient
}

func (executor *BSCExecutor) SwitchClient() {
	executor.mutex.Lock()
	defer executor.mutex.Unlock()
	executor.clientIdx++
	if executor.clientIdx >= len(executor.bscClients) {
		executor.clientIdx = 0
	}
	relayercommon.Logger.Infof("Switch to provider: %s", executor.config.BSCConfig.RPCAddrs[executor.clientIdx])
}

func (executor *BSCExecutor) GetLogsFromHeader(header *types.Header) ([]types.Log, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := executor.GetClient()
	topics := [][]ethereumcommon.Hash{{CrossChainPackageEventHash}}
	blockHash := header.Hash()
	logs, err := client.FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: []ethereumcommon.Address{executor.config.BSCConfig.BSCCrossChainContractAddress},
	})
	if err != nil {
		return nil, err
	}
	return logs, nil

}

func (executor *BSCExecutor) GetLatestBlockHeightWithRetry() (latestHeight uint64, err error) {
	return executor.getLatestBlockHeightWithRetry(executor.GetClient())
}

func (executor *BSCExecutor) getLatestBlockHeightWithRetry(client *ethclient.Client) (latestHeight uint64, err error) {
	return latestHeight, retry.Do(func() error {
		var err error
		latestHeight, err = executor.GetLatestBlockHeight(client)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			relayercommon.Logger.Infof("Failed to query latest height, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (executor *BSCExecutor) GetLatestBlockHeight(client *ethclient.Client) (uint64, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	block, err := client.BlockByNumber(ctxWithTimeout, nil)
	if err != nil {
		return 0, err
	}
	return block.Number().Uint64(), nil
}

func (executor *BSCExecutor) UpdateClients() {
	for {
		relayercommon.Logger.Infof("Start to monitor bsc data-seeds healthy")
		for _, bscClient := range executor.bscClients {
			if time.Since(bscClient.UpdatedAt).Seconds() > DataSeedDenyServiceThreshold {
				msg := fmt.Sprintf("data seed %s is not accessable", bscClient.Provider)
				relayercommon.Logger.Error(msg)
				config.SendTelegramMessage(executor.config.AlertConfig.Identity, executor.config.AlertConfig.TelegramBotId,
					executor.config.AlertConfig.TelegramChatId, msg)
			}
			height, err := executor.GetLatestBlockHeight(bscClient.BSCClient)
			if err != nil {
				relayercommon.Logger.Errorf("get latest block height error, err=%s", err.Error())
				continue
			}
			bscClient.Height = height
			bscClient.UpdatedAt = time.Now()
		}

		highestHeight := uint64(0)
		highestIdx := 0
		for idx := 0; idx < len(executor.bscClients); idx++ {
			if executor.bscClients[idx].Height > highestHeight {
				highestHeight = executor.bscClients[idx].Height
				highestIdx = idx
			}
		}
		// current client block sync is fall behind, switch to the client with the highest block height
		if executor.bscClients[executor.clientIdx].Height+FallBehindThreshold < highestHeight {
			executor.mutex.Lock()
			executor.clientIdx = highestIdx
			executor.mutex.Unlock()
		}
		time.Sleep(SleepSecondForUpdateClient * time.Second)
	}
}

func (executor *BSCExecutor) GetBlockHeaderAtHeight(height uint64) (*types.Header, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	header, err := executor.GetClient().HeaderByNumber(ctxWithTimeout, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return header, nil
}

// TODO
func (executor *BSCExecutor) GetNextSequence(channelID relayercommon.ChannelId) (uint64, error) {
	//crossChainInstance, err := crosschain.NewCrosschain(crossChainContractAddr, executor.GetClient())
	//if err != nil {
	//	return 0, err
	//}
	//
	//callOpts, err := executor.getCallOpts()
	//if err != nil {
	//	return 0, err
	//}
	//
	//return crossChainInstance.ChannelReceiveSequenceMap(callOpts, uint8(channelID))
	return 0, nil
}

func (executor *BSCExecutor) GetNextDeliverySequenceForChannel(channelId relayercommon.ChannelId) (uint64, error) {
	sequence, err := executor.InscriptionExecutor.GetNextSequence(channelId)
	if err != nil {
		return 0, err
	}
	return sequence, nil
}
