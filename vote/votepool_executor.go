package vote

import (
	"context"
	"encoding/hex"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/executor"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/rpc/jsonrpc/client"
	"github.com/tendermint/tendermint/votepool"
)

type VotePoolExecutor struct {
	client              *client.Client
	config              *config.Config
	inscriptionExecutor *executor.InscriptionExecutor
}

func NewVotePoolExecutor(cfg *config.Config, inscriptionExecutor *executor.InscriptionExecutor) *VotePoolExecutor {
	cli, err := client.New(cfg.VotePoolConfig.RPCAddr)
	if err != nil {
		panic(err)
	}
	return &VotePoolExecutor{
		client:              cli,
		config:              cfg,
		inscriptionExecutor: inscriptionExecutor,
	}
}

func (e *VotePoolExecutor) QueryVotes(eventHash []byte, eventType votepool.EventType) ([]*votepool.Vote, error) {
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
	var broadcastVote coretypes.ResultBroadcastVote

	_, err := e.client.Call(context.Background(), VotePoolBroadcastMethodName, broadcastMap, &broadcastVote)
	if err != nil {
		return err
	}
	return nil
}

func (e *VotePoolExecutor) GetValidatorsBlsPublicKey() ([]string, error) {
	validators, err := e.inscriptionExecutor.QueryLatestValidators()
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, v := range validators {
		keys = append(keys, hex.EncodeToString(v.GetRelayerBlsKey()))
	}
	return keys, nil
}
