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
	go r.MonitorCrossChainEvents()
	if r.InscriptionExecutor.IsValidator() {
		go r.SignAndBroadcast()
		go r.CollectVotes()
		go r.AssemblePackages()
	}
	go r.UpdateCachedLatestValidators()
}

// MonitorCrossChainEvents will monitor cross chain events for every block and persist into DB
func (r *BSCRelayer) MonitorCrossChainEvents() {
	r.Listener.Start()
}

func (r *BSCRelayer) SignAndBroadcast() {
	r.voteProcessor.SignAndBroadcast()
}

func (r *BSCRelayer) CollectVotes() {
	r.voteProcessor.CollectVotes()
}

func (r *BSCRelayer) AssemblePackages() {
	r.assembler.AssemblePackagesAndClaim()
}

func (r *BSCRelayer) UpdateCachedLatestValidators() {
	r.bscExecutor.UpdateCachedLatestValidators() // cache validators queried from inscription, update it every 1 minute
}
