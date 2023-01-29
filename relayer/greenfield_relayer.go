package relayer

import (
	"github.com/bnb-chain/greenfield-relayer/assembler"
	"github.com/bnb-chain/greenfield-relayer/executor"
	"github.com/bnb-chain/greenfield-relayer/listener"
	"github.com/bnb-chain/greenfield-relayer/vote"
)

type GreenfieldRelayer struct {
	Listener            *listener.GreenfieldListener
	GreenfieldExecutor  *executor.GreenfieldExecutor
	bscExecutor         *executor.BSCExecutor
	votePoolExecutor    *vote.VotePoolExecutor
	voteProcessor       *vote.GreenfieldVoteProcessor
	greenfieldAssembler *assembler.GreenfieldAssembler
}

func NewGreenfieldRelayer(listener *listener.GreenfieldListener, greenfieldExecutor *executor.GreenfieldExecutor, bscExecutor *executor.BSCExecutor,
	votePoolExecutor *vote.VotePoolExecutor, voteProcessor *vote.GreenfieldVoteProcessor, greenfieldAssembler *assembler.GreenfieldAssembler,
) *GreenfieldRelayer {
	return &GreenfieldRelayer{
		Listener:            listener,
		GreenfieldExecutor:  greenfieldExecutor,
		bscExecutor:         bscExecutor,
		votePoolExecutor:    votePoolExecutor,
		voteProcessor:       voteProcessor,
		greenfieldAssembler: greenfieldAssembler,
	}
}

func (r *GreenfieldRelayer) Start() {
	go r.MonitorEventsLoop()
	go r.SignAndBroadcastLoop()
	go r.CollectVotesLoop()
	go r.AssembleTransactionsLoop()
	go r.UpdateCachedLatestValidatorsLoop()
}

// MonitorEventsLoop will monitor cross chain events for every block and persist into DB
func (r *GreenfieldRelayer) MonitorEventsLoop() {
	r.Listener.Start()
}

func (r *GreenfieldRelayer) SignAndBroadcastLoop() {
	r.voteProcessor.SignAndBroadcast()
}

func (r *GreenfieldRelayer) CollectVotesLoop() {
	r.voteProcessor.CollectVotes()
}

func (r *GreenfieldRelayer) AssembleTransactionsLoop() {
	r.greenfieldAssembler.AssembleTransactionAndSend()
}

func (r *GreenfieldRelayer) UpdateCachedLatestValidatorsLoop() {
	r.GreenfieldExecutor.UpdateCachedLatestValidators() // cache validators queried from greenfield, update it every 1 minute
}
func (r *GreenfieldRelayer) UpdateClientLoop() {
	r.GreenfieldExecutor.UpdateClients()
}
