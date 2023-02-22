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
}

// MonitorEventsLoop will monitor cross chain events for every block and persist into DB
func (r *GreenfieldRelayer) MonitorEventsLoop() {
	r.Listener.StartLoop()
}

func (r *GreenfieldRelayer) SignAndBroadcastLoop() {
	r.voteProcessor.SignAndBroadcastLoop()
}

func (r *GreenfieldRelayer) CollectVotesLoop() {
	r.voteProcessor.CollectVotesLoop()
}

func (r *GreenfieldRelayer) AssembleTransactionsLoop() {
	r.greenfieldAssembler.AssembleTransactionsLoop()
}
