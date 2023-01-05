package executor

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"google.golang.org/grpc"
	relayercommon "inscription-relayer/common"
	"inscription-relayer/config"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bts "github.com/tendermint/tendermint/libs/bytes"
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
	privateKey         *secp256k1.PrivKey
}

func grpcConn(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewRpcClient(addr string) (*rpchttp.HTTP, error) {
	httpClient, err := libclient.DefaultHTTPClient(addr)
	if err != nil {
		return nil, err
	}
	rpcClient, err := rpchttp.NewWithClient(addr, "/websocket", httpClient)
	if err != nil {
		return nil, err
	}
	return rpcClient, nil
}

func getInscriptionPrivateKey(cfg *config.InscriptionConfig) (*secp256k1.PrivKey, error) {
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
	privKey := secp256k1.PrivKey{Key: []byte(privateKey)}
	return &privKey, nil
}

func initInscriptionClients(rpcAddrs, grpcAddrs []string) []*InscriptionClient {
	inscriptionClients := make([]*InscriptionClient, 0)

	for i := 0; i < len(rpcAddrs); i++ {
		rpcClient, err := NewRpcClient(rpcAddrs[i])
		if err != nil {
			panic(err)
		}
		conn, err := grpcConn(grpcAddrs[i])
		if err != nil {
			panic(err)
		}

		inscriptionClients = append(inscriptionClients, &InscriptionClient{
			txClient:           tx.NewServiceClient(conn),
			stakingQueryClient: stakingtypes.NewQueryClient(conn),
			authClient:         authtypes.NewQueryClient(conn),
			rpcClient:          rpcClient,
			Provider:           rpcAddrs[i],
			UpdatedAt:          time.Now(),
		})
	}

	return inscriptionClients
}

func NewInscriptionExecutor(cfg *config.Config) (*InscriptionExecutor, error) {
	privKey, err := getInscriptionPrivateKey(&cfg.InscriptionConfig)
	if err != nil {
		return nil, err
	}
	return &InscriptionExecutor{
		clientIdx:          0,
		inscriptionClients: initInscriptionClients(cfg.InscriptionConfig.RPCAddrs, cfg.InscriptionConfig.GRPCAddrs),
		privateKey:         privKey,
		config:             cfg,
	}, nil
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

func (e *InscriptionExecutor) GetBlockResultAtHeight(height int64) (*ctypes.ResultBlockResults, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	blockResults, err := e.getRpcClient().BlockResults(ctx, &height)
	if err != nil {
		return nil, err
	}
	return blockResults, nil
}

func (e *InscriptionExecutor) GetBlockAtHeight(height int64) (*tmtypes.Block, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	block, err := e.getRpcClient().Block(ctx, &height)
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
		latestHeight, err = e.GetLatestBlockHeight(latestHeightQueryCtx, client)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			relayercommon.Logger.Infof("Failed to query latest height, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (e *InscriptionExecutor) GetLatestBlockHeight(ctx context.Context, client rpcclient.Client) (uint64, error) {
	status, err := client.Status(ctx)
	if err != nil {
		return 0, err
	}
	return uint64(status.SyncInfo.LatestBlockHeight), nil
}

func (e *InscriptionExecutor) UpdateClients() {
	for {
		relayercommon.Logger.Infof("Start to monitor inscription data-seeds healthy")
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
		// current InscriptionClient block sync is fall behind, switch to the InscriptionClient with highest block height
		if e.inscriptionClients[e.clientIdx].Height+FallBehindThreshold < highestHeight {
			e.mutex.Lock()
			e.clientIdx = highestIdx
			e.mutex.Unlock()
		}
		time.Sleep(SleepSecondForUpdateClient * time.Second)
	}
}

func (e *InscriptionExecutor) MonitorValidatorSetChange(height int64, preValidatorsHash bts.HexBytes) (bool, bts.HexBytes, error) {
	validatorSetChanged := false

	block, err := e.getRpcClient().Block(context.Background(), &height)
	if err != nil {
		return false, nil, err
	}

	var curValidatorsHash bts.HexBytes
	if preValidatorsHash != nil {
		if !bytes.Equal(block.Block.Header.ValidatorsHash, preValidatorsHash) ||
			!bytes.Equal(block.Block.Header.ValidatorsHash, block.Block.Header.NextValidatorsHash) {
			validatorSetChanged = true
			curValidatorsHash = block.Block.Header.ValidatorsHash
		} else {
			curValidatorsHash = preValidatorsHash
		}
	}

	return validatorSetChanged, curValidatorsHash, nil
}

func (e *InscriptionExecutor) QueryTendermintHeader(height int64) (*relayercommon.Header, error) {

	commit, err := e.getRpcClient().Commit(context.Background(), &height)
	if err != nil {
		return nil, err
	}

	validators, err := e.QueryLatestValidators()
	if err != nil {
		return nil, err
	}

	var blsPubKeysBts []byte
	var relayerAddrs []string
	for _, v := range validators {
		blsPubKeysBts = append(blsPubKeysBts, v.RelayerBlsKey...)
		relayerAddrs = append(relayerAddrs, v.RelayerAddress)
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
	sequence, err := e.BscExecutor.GetNextSequence(channelID)
	if err != nil {
		return 0, err
	}
	return sequence, nil
}

func (e *InscriptionExecutor) GetNextOracleSequence() (uint64, error) {
	path := fmt.Sprintf("/store/%s/%s", SequenceStoreName, "key")
	key := BuildChannelSequenceKey(relayercommon.ChainId(e.config.BSCConfig.ChainId), 0x00)
	response, err := e.getRpcClient().ABCIQuery(context.Background(), path, key)
	if err != nil {
		return 0, err
	}
	if response.Response.Value == nil {
		return 0, nil
	}
	return binary.BigEndian.Uint64(response.Response.Value), nil
}

func (e *InscriptionExecutor) QueryLatestValidators() ([]stakingtypes.Validator, error) {
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
	hist := result.Hist
	return hist.Valset, nil
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
