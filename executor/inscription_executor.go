package executor

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"inscription-relayer/assembler"
	relayercommon "inscription-relayer/common"
	"inscription-relayer/config"
	"inscription-relayer/db/dao"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/tx"
	chantypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	bts "github.com/tendermint/tendermint/libs/bytes"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	libclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
	tmtypes "github.com/tendermint/tendermint/types"
)

type InscriptionClient struct {
	rpcClient    rpcclient.Client
	grpcTxClient tx.ServiceClient
	chanClient   chantypes.QueryClient
	Provider     string
	Height       uint64
	UpdatedAt    time.Time
}

type InscriptionExecutor struct {
	mutex              sync.RWMutex
	daoManager         *dao.DaoManager
	BSCExecutor        *BSCExecutor
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

func initInscriptionClients(providers []string) []*InscriptionClient {
	inscriptionClients := make([]*InscriptionClient, 0)
	for _, provider := range providers {
		rpcClient, err := NewRpcClient(provider)
		if err != nil {
			panic("new RPC client error")
		}
		conn, err := grpcConn(provider)
		if err != nil {
			panic("new GRPC connection error")
		}
		grpcTxClient := tx.NewServiceClient(conn)
		chanClient := chantypes.NewQueryClient(conn)

		inscriptionClients = append(inscriptionClients, &InscriptionClient{
			grpcTxClient: grpcTxClient,
			chanClient:   chanClient,
			rpcClient:    rpcClient,
			Provider:     provider,
			UpdatedAt:    time.Now(),
		})
	}
	return inscriptionClients
}

func NewInscriptionExecutor(cfg *config.Config, dao *dao.DaoManager) (*InscriptionExecutor, error) {
	privKey, err := getInscriptionPrivateKey(&cfg.InscriptionConfig)
	if err != nil {
		return nil, err
	}
	return &InscriptionExecutor{
		daoManager:         dao,
		clientIdx:          0,
		inscriptionClients: initInscriptionClients(cfg.InscriptionConfig.RPCAddrs),
		privateKey:         privKey,
		config:             cfg,
	}, nil
}

func (executor *InscriptionExecutor) SetBSCExecutor(bscExecutor *BSCExecutor) {
	executor.BSCExecutor = bscExecutor
}

func (executor *InscriptionExecutor) GetRpcClient() rpcclient.Client {
	executor.mutex.RLock()
	defer executor.mutex.RUnlock()
	return executor.inscriptionClients[executor.clientIdx].rpcClient
}

func (executor *InscriptionExecutor) GetGrpcTxClient() tx.ServiceClient {
	executor.mutex.RLock()
	defer executor.mutex.RUnlock()
	return executor.inscriptionClients[executor.clientIdx].grpcTxClient
}

func (executor *InscriptionExecutor) GetChanClient() chantypes.QueryClient {
	executor.mutex.RLock()
	defer executor.mutex.RUnlock()
	return executor.inscriptionClients[executor.clientIdx].chanClient
}

func (executor *InscriptionExecutor) GetBlockResultAtHeight(height int64) (*ctypes.ResultBlockResults, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	blockResults, err := executor.GetRpcClient().BlockResults(ctx, &height)
	if err != nil {
		return nil, err
	}
	return blockResults, nil
}

func (executor *InscriptionExecutor) GetBlockAtHeight(height int64) (*tmtypes.Block, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	block, err := executor.GetRpcClient().Block(ctx, &height)
	if err != nil {
		return nil, err
	}
	return block.Block, nil
}

func (executor *InscriptionExecutor) GetLatestBlockHeightWithRetry() (latestHeight uint64, err error) {
	return executor.getLatestBlockHeightWithRetry(executor.GetRpcClient())
}

func (executor *InscriptionExecutor) getLatestBlockHeightWithRetry(client rpcclient.Client) (latestHeight uint64, err error) {
	return latestHeight, retry.Do(func() error {
		latestHeightQueryCtx, cancelLatestHeightQueryCtx := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelLatestHeightQueryCtx()
		var err error
		latestHeight, err = executor.GetLatestBlockHeight(latestHeightQueryCtx, client)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			relayercommon.Logger.Infof("Failed to query latest height, attempt: %d times, max_attempts: %d", n+1, relayercommon.RtyAttNum)
		}))
}

func (executor *InscriptionExecutor) GetLatestBlockHeight(ctx context.Context, client rpcclient.Client) (uint64, error) {
	status, err := client.Status(ctx)
	if err != nil {
		return 0, err
	}
	return uint64(status.SyncInfo.LatestBlockHeight), nil
}

func (executor *InscriptionExecutor) UpdateClients() {
	for {
		relayercommon.Logger.Infof("Start to monitor inscription data-seeds healthy")
		for _, inscriptionClient := range executor.inscriptionClients {
			if time.Since(inscriptionClient.UpdatedAt).Seconds() > DataSeedDenyServiceThreshold {
				msg := fmt.Sprintf("data seed %s is not accessable", inscriptionClient.Provider)
				relayercommon.Logger.Error(msg)
				config.SendTelegramMessage(executor.config.AlertConfig.Identity, executor.config.AlertConfig.TelegramBotId,
					executor.config.AlertConfig.TelegramChatId, msg)
			}
			height, err := executor.getLatestBlockHeightWithRetry(inscriptionClient.rpcClient)
			if err != nil {
				relayercommon.Logger.Errorf("get latest block height error, err=%s", err.Error())
				continue
			}
			inscriptionClient.Height = height
			inscriptionClient.UpdatedAt = time.Now()
		}

		highestHeight := uint64(0)
		highestIdx := 0
		for idx := 0; idx < len(executor.inscriptionClients); idx++ {
			if executor.inscriptionClients[idx].Height > highestHeight {
				highestHeight = executor.inscriptionClients[idx].Height
				highestIdx = idx
			}
		}
		// current InscriptionClient block sync is fall behind, switch to the InscriptionClient with highest block height
		if executor.inscriptionClients[executor.clientIdx].Height+FallBehindThreshold < highestHeight {
			executor.mutex.Lock()
			executor.clientIdx = highestIdx
			executor.mutex.Unlock()
		}
		time.Sleep(SleepSecondForUpdateClient * time.Second)
	}
}

func (executor *InscriptionExecutor) MonitorValidatorSetChange(height int64, preValidatorsHash bts.HexBytes) (bool, bts.HexBytes, error) {
	validatorSetChanged := false

	block, err := executor.GetRpcClient().Block(context.Background(), &height)
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

func (executor *InscriptionExecutor) QueryTendermintHeader(height int64) (*relayercommon.Header, error) {
	nextHeight := height + 1

	commit, err := executor.GetRpcClient().Commit(context.Background(), &height)
	if err != nil {
		return nil, err
	}

	validators, err := executor.GetRpcClient().Validators(context.Background(), &height, nil, nil)
	if err != nil {
		return nil, err
	}

	nextvalidators, err := executor.GetRpcClient().Validators(context.Background(), &nextHeight, nil, nil)
	if err != nil {
		return nil, err
	}

	header := &relayercommon.Header{
		SignedHeader:     commit.SignedHeader,
		ValidatorSet:     tmtypes.NewValidatorSet(validators.Validators),
		NextValidatorSet: tmtypes.NewValidatorSet(nextvalidators.Validators),
	}

	return header, nil
}

// GetNextDeliverySequenceForChannel call dest chain(BSC) to return a sequence# which should be used.
func (executor *InscriptionExecutor) GetNextDeliverySequenceForChannel(channelID relayercommon.ChannelId) (uint64, error) {
	sequence, err := executor.BSCExecutor.GetNextSequence(channelID)
	if err != nil {
		return 0, err
	}
	return sequence, nil
}

func (executor *InscriptionExecutor) GetNextSequence(channelId relayercommon.ChannelId) (uint64, error) {
	path := fmt.Sprintf("/store/%s/%s", SequenceStoreName, "key")
	key := BuildChannelSequenceKey(relayercommon.ChainId(executor.config.InscriptionConfig.ChainId), channelId)
	response, err := executor.GetRpcClient().ABCIQuery(context.Background(), path, key)
	if err != nil {
		return 0, err
	}
	if response.Response.Value == nil {
		return 0, nil
	}
	return binary.BigEndian.Uint64(response.Response.Value), nil
}

// ClaimPackages TODO use inscription-cosmos-sdk to claim a transaction
func (executor *InscriptionExecutor) ClaimPackages(m *assembler.MsgClaim) (*assembler.MsgClaimResponse, error) {
	//
	//interfaceRegistry := cdctypes.NewInterfaceRegistry()
	//cdc := codec.NewProtoCodec(interfaceRegistry)
	//txConfig := authtx.NewTxConfig(cdc, authtx.DefaultSignModes)
	//txBuilder := txConfig.NewTxBuilder()
	//
	//privKey := executor.privateKey
	//
	//txBuilder.SetMsgs()
	//
	//err := txBuilder.SetMsgs(msg)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var sigsV2 []signing.SignatureV2
	//accountNum := executor.config.InscriptionConfig.AccountNum
	//accountSeq := executor.config.InscriptionConfig.AccountSequence
	//
	//// First round: we gather all the signer infos. We use the "set empty
	//// signature" hack to do that.
	//sigV2 := signing.SignatureV2{
	//	PubKey: privKey.PubKey(),
	//	Data: &signing.SingleSignatureData{
	//		SignMode:  txConfig.SignModeHandler().DefaultMode(),
	//		Signature: nil,
	//	},
	//	Sequence: accountSeq,
	//}
	//sigsV2 = append(sigsV2, sigV2)
	//
	//err = txBuilder.SetSignatures(sigsV2...)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Second round: all signer infos are set, so each signer can sign.
	//sigsV2 = []signing.SignatureV2{}
	//
	//signerData := xauthsigning.SignerData{
	//	ChainID:       strconv.Itoa(int(executor.config.InscriptionConfig.ChainId)),
	//	AccountNumber: accountNum,
	//	Sequence:      accountSeq,
	//}
	//
	//sigV2, err = clitx.SignWithPrivKey(txConfig.SignModeHandler().DefaultMode(), signerData, txBuilder, privKey, txConfig, accountSeq)
	//
	//sigsV2 = append(sigsV2, sigV2)
	//
	//err = txBuilder.SetSignatures(sigsV2...)
	//if err != nil {
	//	return nil, err
	//}
	//
	////
	//txBytes, err := txConfig.TxEncoder()(txBuilder.GetTx())
	//
	////Broadcast transaction
	//txRes, err := executor.GetGrpcTxClient().BroadcastTx(
	//	context.Background(),
	//	&tx.BroadcastTxRequest{
	//		Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
	//		TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
	//	})
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}
