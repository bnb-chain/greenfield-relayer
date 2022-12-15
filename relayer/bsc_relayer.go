package relayer

import (
	"inscription-relayer/assembler"
	"inscription-relayer/common"
	"inscription-relayer/executor"
	"inscription-relayer/listener"
	"inscription-relayer/vote"
)

type BSCRelayer struct {
	listener            *listener.BSCListener
	inscriptionExecutor *executor.InscriptionExecutor
	bscExecutor         *executor.BSCExecutor
	votePoolExecutor    *executor.VotePoolExecutor
	voteProcessor       *vote.InscriptionVoteProcessor
	bscAssembler        *assembler.BSCAssembler
}

func NewBSCRelayer(listener *listener.BSCListener, inscriptionExecutor *executor.InscriptionExecutor, bscExecutor *executor.BSCExecutor,
	votePoolExecutor *executor.VotePoolExecutor, voteProcessor *vote.InscriptionVoteProcessor, bscAssembler *assembler.BSCAssembler) *BSCRelayer {
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
		return
	}
}

func (r *BSCRelayer) collectVotes() {
	err := r.voteProcessor.CollectVotes()
	if err != nil {
		return
	}
}

func (r *BSCRelayer) assemblePackages() {
	err := r.bscAssembler.AssemblePackagesAndClaim()
	if err != nil {
		common.Logger.Errorf("Entering err when assemblePackages, err %s ", err)
		return
	}
}
