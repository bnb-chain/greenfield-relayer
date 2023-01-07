package relayer

import (
	"github.com/bnb-chain/inscription-relayer/assembler"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/listener"
	"github.com/bnb-chain/inscription-relayer/vote"
)

type InscriptionRelayer struct {
	listener             *listener.InscriptionListener
	inscriptionExecutor  *executor.InscriptionExecutor
	bscExecutor          *executor.BSCExecutor
	votePoolExecutor     *vote.VotePoolExecutor
	voteProcessor        *vote.InscriptionVoteProcessor
	inscriptionAssembler *assembler.InscriptionAssembler
}

func NewInscriptionRelayer(listener *listener.InscriptionListener, inscriptionExecutor *executor.InscriptionExecutor, bscExecutor *executor.BSCExecutor,
	votePoolExecutor *vote.VotePoolExecutor, voteProcessor *vote.InscriptionVoteProcessor, inscriptionAssembler *assembler.InscriptionAssembler) *InscriptionRelayer {
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
	r.voteProcessor.SignAndBroadcast()
}

func (r *InscriptionRelayer) collectVotes() {
	r.voteProcessor.CollectVotes()
}

func (r *InscriptionRelayer) assembleTransactions() {
	r.inscriptionAssembler.AssembleTransactionAndSend()
}
