package relayer

import (
	"fmt"
	"inscription-relayer/assembler"
	"inscription-relayer/executor"
	"inscription-relayer/listener"
	"inscription-relayer/vote"
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
	votePoolExecutor *vote.VotePoolExecutor, voteProcessor *vote.BSCVoteProcessor, bscAssembler *assembler.BSCAssembler) *BSCRelayer {
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
	err := r.voteProcessor.SignAndBroadcast()
	if err != nil {
		panic(fmt.Sprintf("Entering err when signAndBroadcast, err %s ", err))
	}
}

func (r *BSCRelayer) collectVotes() {
	err := r.voteProcessor.CollectVotes()
	if err != nil {
		panic(fmt.Sprintf("Entering err when collectVotes, err %s ", err))
	}
}

func (r *BSCRelayer) assemblePackages() {
	r.bscAssembler.AssemblePackagesAndClaim()
}
