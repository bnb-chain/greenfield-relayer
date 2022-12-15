package relayer

import (
	"inscription-relayer/assembler"
	"inscription-relayer/common"
	"inscription-relayer/executor"
	"inscription-relayer/listener"
	"inscription-relayer/vote"
)

type InscriptionRelayer struct {
	listener             *listener.InscriptionListener
	inscriptionExecutor  *executor.InscriptionExecutor
	bscExecutor          *executor.BSCExecutor
	votePoolExecutor     *executor.VotePoolExecutor
	voteProcessor        *vote.InscriptionVoteProcessor
	inscriptionAssembler *assembler.InscriptionAssembler
}

func NewInscriptionRelayer(listener *listener.InscriptionListener, inscriptionExecutor *executor.InscriptionExecutor, bscExecutor *executor.BSCExecutor,
	votePoolExecutor *executor.VotePoolExecutor, voteProcessor *vote.InscriptionVoteProcessor, inscriptionAssembler *assembler.InscriptionAssembler) *InscriptionRelayer {
	return &InscriptionRelayer{
		listener:             listener,
		inscriptionExecutor:  inscriptionExecutor,
		bscExecutor:          bscExecutor,
		votePoolExecutor:     votePoolExecutor,
		voteProcessor:        voteProcessor,
		inscriptionAssembler: inscriptionAssembler,
	}
}

func (r *InscriptionRelayer) Start() {
	go r.monitorCrossChainEvents()
	go r.signAndBroadcast()
	go r.collectVotes()
	go r.assembleTransactions()
}

// monitorCrossChainEvents will monitor cross chain events for every block and persist into DB
func (r *InscriptionRelayer) monitorCrossChainEvents() {
	r.listener.Start()
}

func (r *InscriptionRelayer) signAndBroadcast() {
	err := r.voteProcessor.SignAndBroadcast()
	if err != nil {
		common.Logger.Errorf("Entering err when broadcastUnVotes, err %s ", err)
		return
	}
}

func (r *InscriptionRelayer) collectVotes() {
	err := r.voteProcessor.CollectVotes()
	if err != nil {
		common.Logger.Errorf("Entering err when collectVotes, err %s ", err)
		return
	}
}

func (r *InscriptionRelayer) assembleTransactions() {
	err := r.inscriptionAssembler.AssembleTransactionAndSend()
	if err != nil {
		common.Logger.Errorf("Entering err when assemblePackages, err %s ", err)
		return
	}
}
