package relayer

import (
	"github.com/bnb-chain/inscription-relayer/assembler"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/listener"
	"github.com/bnb-chain/inscription-relayer/vote"
)

type BSCRelayer struct {
	Listener            *listener.BSCListener
	InscriptionExecutor *executor.InscriptionExecutor
	bscExecutor         *executor.BSCExecutor
	VotePoolExecutor    *vote.VotePoolExecutor
	voteProcessor       *vote.BSCVoteProcessor
	assembler           *assembler.BSCAssembler
}

func NewBSCRelayer(listener *listener.BSCListener, inscriptionExecutor *executor.InscriptionExecutor, bscExecutor *executor.BSCExecutor,
	votePoolExecutor *vote.VotePoolExecutor, voteProcessor *vote.BSCVoteProcessor, bscAssembler *assembler.BSCAssembler,
) *BSCRelayer {
	return &BSCRelayer{
		Listener:            listener,
		InscriptionExecutor: inscriptionExecutor,
		bscExecutor:         bscExecutor,
		VotePoolExecutor:    votePoolExecutor,
		voteProcessor:       voteProcessor,
		assembler:           bscAssembler,
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
	r.Listener.Start()
}

func (r *BSCRelayer) SignAndBroadcastVoteLoop() {
	r.voteProcessor.SignAndBroadcast()
}

func (r *BSCRelayer) CollectVotesLoop() {
	r.voteProcessor.CollectVotes()
}

func (r *BSCRelayer) AssemblePackagesLoop() {
	r.assembler.AssemblePackagesAndClaim()
}

func (r *BSCRelayer) UpdateCachedLatestValidatorsLoop() {
	r.bscExecutor.UpdateCachedLatestValidators() // cache validators queried from inscription, update it every 1 minute
}

func (r *BSCRelayer) UpdateClientLoop() {
	r.bscExecutor.UpdateClients()
}
