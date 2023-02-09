package vote

import (
	"context"

	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/rpc/jsonrpc/client"
	"github.com/tendermint/tendermint/votepool"

	"github.com/bnb-chain/greenfield-relayer/config"
)

type VotePoolExecutor struct {
	client *client.Client
	cfg    *config.Config
}

func NewVotePoolExecutor(cfg *config.Config) *VotePoolExecutor {
	cli, err := client.New(cfg.VotePoolConfig.RPCAddr)
	if err != nil {
		panic(err)
	}
	return &VotePoolExecutor{
		client: cli,
		cfg:    cfg,
	}
}

func (e *VotePoolExecutor) QueryVotesByEventHashAndType(eventHash []byte, eventType votepool.EventType) ([]*votepool.Vote, error) {
	queryMap := make(map[string]interface{})
	queryMap[VotePoolQueryParameterEventType] = int(eventType)
	queryMap[VotePoolQueryParameterEventHash] = eventHash
	var queryVote coretypes.ResultQueryVote
	_, err := e.client.Call(context.Background(), VotePoolQueryMethodName, queryMap, &queryVote)
	if err != nil {
		return nil, err
	}
	return queryVote.Votes, nil
}

func (e *VotePoolExecutor) BroadcastVote(v *votepool.Vote) error {
	broadcastMap := make(map[string]interface{})
	broadcastMap[VotePoolBroadcastParameterKey] = *v
	_, err := e.client.Call(context.Background(), VotePoolBroadcastMethodName, broadcastMap, &coretypes.ResultBroadcastVote{})
	if err != nil {
		return err
	}
	return nil
}

func (e *VotePoolExecutor) GetBlsPrivateKey() string {
	return e.cfg.VotePoolConfig.BlsPrivateKey
}
