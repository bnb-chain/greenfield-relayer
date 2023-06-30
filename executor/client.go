package executor

import (
	"context"
	"sync"

	jsonrpcclient "github.com/cometbft/cometbft/rpc/jsonrpc/client"

	sdkclient "github.com/bnb-chain/greenfield-go-sdk/client"
	"github.com/bnb-chain/greenfield-go-sdk/types"
	"github.com/bnb-chain/greenfield-relayer/logging"
	"github.com/bnb-chain/greenfield/sdk/client"
)

type JsonRpcClient = *jsonrpcclient.Client

type GnfdCompositeClient struct {
	sdkclient.Client
	client.TendermintClient
	JsonRpcClient
	Height int64
}

type GnfdCompositeClients struct {
	clients []*GnfdCompositeClient
}

func NewGnfdCompositClients(rpcAddrs []string, chainId string, account *types.Account) GnfdCompositeClients {
	clients := make([]*GnfdCompositeClient, 0)
	for i := 0; i < len(rpcAddrs); i++ {

		sdkClient, err := sdkclient.New(chainId, rpcAddrs[i], sdkclient.Option{DefaultAccount: account})
		if err != nil {
			logging.Logger.Errorf("rpc node %s is not available, err: %v", rpcAddrs[i], err)
			continue
		}
		jsonRpcClient, err := jsonrpcclient.New(rpcAddrs[i])
		if err != nil {
			logging.Logger.Errorf("jsonrpc node %s is not available, err: %v", rpcAddrs[i], err)
			continue
		}
		clients = append(clients, &GnfdCompositeClient{
			Client:           sdkClient,
			TendermintClient: client.NewTendermintClient(rpcAddrs[i]),
			JsonRpcClient:    jsonRpcClient,
		})
	}
	return GnfdCompositeClients{
		clients: clients,
	}
}

func (gc *GnfdCompositeClients) GetClient() *GnfdCompositeClient {
	wg := new(sync.WaitGroup)
	wg.Add(len(gc.clients))
	clientCh := make(chan *GnfdCompositeClient)
	waitCh := make(chan struct{})
	go func() {
		for _, c := range gc.clients {
			go getClientBlockHeight(clientCh, wg, c)
		}
		wg.Wait()
		close(waitCh)
	}()
	var maxHeight int64
	maxHeightClient := gc.clients[0]
	for {
		select {
		case c := <-clientCh:
			if c.Height > maxHeight {
				maxHeight = c.Height
				maxHeightClient = c
			}
		case <-waitCh:
			return maxHeightClient
		}
	}
}

func getClientBlockHeight(clientChan chan *GnfdCompositeClient, wg *sync.WaitGroup, client *GnfdCompositeClient) {
	defer wg.Done()
	status, err := client.TmClient.Status(context.Background())
	if err != nil {
		return
	}
	latestHeight := status.SyncInfo.LatestBlockHeight
	client.Height = latestHeight
	clientChan <- client
}
