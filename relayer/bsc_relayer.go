package relayer

import (
	"github.com/bnb-chain/greenfield-relayer/assembler"
	"github.com/bnb-chain/greenfield-relayer/executor"
	"github.com/bnb-chain/greenfield-relayer/listener"
	"github.com/bnb-chain/greenfield-relayer/vote"
)

type BSCRelayer struct {
	Listener           *listener.BSCListener
	GreenfieldExecutor *executor.GreenfieldExecutor
	bscExecutor        *executor.BSCExecutor
	VotePoolExecutor   *vote.VotePoolExecutor
	voteProcessor      *vote.BSCVoteProcessor
	assembler          *assembler.BSCAssembler
}

func NewBSCRelayer(listener *listener.BSCListener, greenfieldExecutor *executor.GreenfieldExecutor, bscExecutor *executor.BSCExecutor,
	votePoolExecutor *vote.VotePoolExecutor, voteProcessor *vote.BSCVoteProcessor, bscAssembler *assembler.BSCAssembler,
) *BSCRelayer {
	return &BSCRelayer{
		Listener:           listener,
		GreenfieldExecutor: greenfieldExecutor,
		bscExecutor:        bscExecutor,
		VotePoolExecutor:   votePoolExecutor,
		voteProcessor:      voteProcessor,
		assembler:          bscAssembler,
	}
}

func (r *BSCRelayer) Start() {
	go r.MonitorEventsLoop()
	go r.SignAndBroadcastVoteLoop()
	go r.CollectVotesLoop()
	go r.AssemblePackagesLoop()
	go r.UpdateCachedLatestValidatorsLoop()
	go r.UpdateClientLoop()
}

// MonitorEventsLoop will monitor cross chain events for every block and persist into DB
func (r *BSCRelayer) MonitorEventsLoop() {
	r.Listener.StartLoop()
}

func (r *BSCRelayer) SignAndBroadcastVoteLoop() {
	r.voteProcessor.SignAndBroadcastVoteLoop()
}

func (r *BSCRelayer) CollectVotesLoop() {
	r.voteProcessor.CollectVotesLoop()
}

func (r *BSCRelayer) AssemblePackagesLoop() {
	r.assembler.AssemblePackagesAndClaimLoop()
}

func (r *BSCRelayer) UpdateCachedLatestValidatorsLoop() {
	r.bscExecutor.UpdateCachedLatestValidatorsLoop() // cache validators queried from greenfield, update it every 1 minute
}

func (r *BSCRelayer) UpdateClientLoop() {
	r.bscExecutor.UpdateClientLoop()
}
