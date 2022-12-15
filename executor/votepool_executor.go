package executor

import (
	"context"
	"encoding/hex"
	"github.com/tendermint/tendermint/types"
	"inscription-relayer/config"
	"inscription-relayer/vote"
)

type VotePoolExecutor struct {
	config              *config.Config
	inscriptionExecutor *InscriptionExecutor
	bscExecutor         *BSCExecutor
}

func NewVotePoolExecutor(cfg *config.Config, inscriptionExecutor *InscriptionExecutor,
	bscExecutor *BSCExecutor) *VotePoolExecutor {
	return &VotePoolExecutor{
		config:              cfg,
		inscriptionExecutor: inscriptionExecutor,
		bscExecutor:         bscExecutor,
	}
}

func (executor *VotePoolExecutor) QueryVotes(eventHash []byte, ty vote.EventType) ([]*vote.Vote, error) {
	//rpc.NewRPCFunc(QueryVote, "event_type,event_hash")
	return nil, nil

}

func (executor *VotePoolExecutor) SubmitVote(v *vote.Vote) error {
	//rpc.NewRPCFunc(BroadcastVote, "vote", v)
	return nil
}

func (executor *VotePoolExecutor) QueryValidators() ([]*types.Validator, error) {
	latestHeight, err := executor.inscriptionExecutor.GetLatestBlockHeightWithRetry()
	height := int64(latestHeight)
	validatorsResult, err := executor.inscriptionExecutor.GetRpcClient().Validators(context.Background(), &height, nil, nil)
	if err != nil {
		return nil, err
	}
	return validatorsResult.Validators, nil
}

func (executor *VotePoolExecutor) GetValidatorsAddresses() ([]string, error) {
	validators, err := executor.QueryValidators()
	if err != nil {
		return nil, err
	}
	var addresses []string
	for _, v := range validators {
		addresses = append(addresses, v.Address.String())
	}
	return addresses, nil
}

// TODO Assume its bls public key for now, need to use inscription-tendermint repo for new field relayer_bls_key and relayer_bls_address
func (executor *VotePoolExecutor) GetValidatorsPublicKey() ([]string, error) {
	validators, err := executor.QueryValidators()
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, v := range validators {
		keys = append(keys, hex.EncodeToString(v.PubKey.Bytes()))
	}
	return keys, nil
}
