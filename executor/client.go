package executor

import (
	"context"
	"sync"

	sdkclient "github.com/bnb-chain/greenfield-go-sdk/client"
	"github.com/bnb-chain/greenfield-go-sdk/types"
	"github.com/bnb-chain/greenfield-relayer/logging"
)

type GreenfieldClient struct {
	sdkclient.Client
	Height int64
}

type GnfdCompositeClients struct {
	clients []*GreenfieldClient
}

func NewGnfdCompositClients(rpcAddrs []string, chainId string, account *types.Account, useWebsocket bool) GnfdCompositeClients {
	clients := make([]*GreenfieldClient, 0)
	for i := 0; i < len(rpcAddrs); i++ {
		sdkClient, err := sdkclient.New(chainId, rpcAddrs[i], sdkclient.Option{DefaultAccount: account, UseWebSocketConn: useWebsocket})
		if err != nil {
			logging.Logger.Errorf("rpc node %s is not available", rpcAddrs[i])
			continue
		}
		clients = append(clients, &GreenfieldClient{
			Client: sdkClient,
		})
	}
	return GnfdCompositeClients{
		clients: clients,
	}
}

func (gc *GnfdCompositeClients) GetClient() *GreenfieldClient {
	wg := new(sync.WaitGroup)
	wg.Add(len(gc.clients))
	clientCh := make(chan *GreenfieldClient)
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

func getClientBlockHeight(clientChan chan *GreenfieldClient, wg *sync.WaitGroup, client *GreenfieldClient) {
	defer wg.Done()
	status, err := client.GetStatus(context.Background())
	if err != nil {
		return
	}
	latestHeight := status.SyncInfo.LatestBlockHeight
	client.Height = latestHeight
	clientChan <- client
}
