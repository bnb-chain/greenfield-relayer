package relayer

import (
	"github.com/bnb-chain/inscription-relayer/assembler"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/listener"
	"github.com/bnb-chain/inscription-relayer/vote"
)

type BSCRelayer struct {
	listener            *listener.BSCListener
	inscriptionExecutor *executor.InscriptionExecutor
	bscExecutor         *executor.BSCExecutor
	votePoolExecutor    *vote.VotePoolExecutor
	voteProcessor       *vote.BSCVoteProcessor
	bscAssembler        *assembler.BSCAssembler
}

func NewBSCRelayer(listener *listener.BSCListener, inscriptionExecutor *executor.InscriptionExecutor, bscExecutor *executor.BSCExecutor,
	votePoolExecutor *vote.VotePoolExecutor, voteProcessor *vote.BSCVoteProcessor, bscAssembler *assembler.BSCAssembler,
) *BSCRelayer {
	return &BSCRelayer{
		listener:            listener,
		inscriptionExecutor: inscriptionExecutor,
		bscExecutor:         bscExecutor,
		votePoolExecutor:    votePoolExecutor,
		voteProcessor:       voteProcessor,
		bscAssembler:        bscAssembler,
	}
}

func (r *BSCRelayer) Start() {
	go r.monitorCrossChainEvents()
	go r.signAndBroadcast()
	go r.collectVotes()
	go r.assemblePackages()
}

// monitorCrossChainEvents will monitor cross chain events for every block and persist into DB
func (r *BSCRelayer) monitorCrossChainEvents() {
	r.listener.Start()
}

func (r *BSCRelayer) signAndBroadcast() {
	r.voteProcessor.SignAndBroadcast()
}

func (r *BSCRelayer) collectVotes() {
	r.voteProcessor.CollectVotes()
}

func (r *BSCRelayer) assemblePackages() {
	r.bscAssembler.AssemblePackagesAndClaim()
}
