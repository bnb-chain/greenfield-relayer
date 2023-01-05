package vote

import (
	"context"
	"encoding/hex"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/rpc/jsonrpc/client"
	"github.com/tendermint/tendermint/votepool"
	"inscription-relayer/config"
	"inscription-relayer/executor"
)

// when query vote, do we need to connect to all rpc endpoint
type VotePoolExecutor struct {
	client              *client.Client
	config              *config.Config
	inscriptionExecutor *executor.InscriptionExecutor
}

func NewVotePoolExecutor(cfg *config.Config, inscriptionExecutor *executor.InscriptionExecutor) *VotePoolExecutor {
	cli, err := client.New("http://127.0.0.1:26657")
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
	queryMap["event_type"] = int(eventType)
	queryMap["event_hash"] = eventHash
	var queryVote coretypes.ResultQueryVote
	_, err := e.client.Call(context.Background(), "query_vote", queryMap, &queryVote)
	if err != nil {
		return nil, err
	}
	return queryVote.Votes, nil
}

func (e *VotePoolExecutor) BroadcastVote(v *votepool.Vote) error {
	broadcastMap := make(map[string]interface{})
	broadcastMap["vote"] = *v
	var broadcastVote coretypes.ResultBroadcastVote

	_, err := e.client.Call(context.Background(), "broadcast_vote", broadcastMap, &broadcastVote)
	if err != nil {
		return err
	}

	return nil
}

func (e *VotePoolExecutor) GetValidatorsAddresses() ([]string, error) {
	validators, err := e.inscriptionExecutor.QueryLatestValidators()
	if err != nil {
		return nil, err
	}
	var addresses []string
	for _, v := range validators {
		addresses = append(addresses, v.RelayerAddress)
	}
	return addresses, nil
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
