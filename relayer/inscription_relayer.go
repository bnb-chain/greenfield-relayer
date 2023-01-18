package relayer

import (
	"github.com/bnb-chain/inscription-relayer/assembler"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/listener"
	"github.com/bnb-chain/inscription-relayer/vote"
)

type InscriptionRelayer struct {
	Listener             *listener.InscriptionListener
	InscriptionExecutor  *executor.InscriptionExecutor
	bscExecutor          *executor.BSCExecutor
	votePoolExecutor     *vote.VotePoolExecutor
	voteProcessor        *vote.InscriptionVoteProcessor
	inscriptionAssembler *assembler.InscriptionAssembler
}

func NewInscriptionRelayer(listener *listener.InscriptionListener, inscriptionExecutor *executor.InscriptionExecutor, bscExecutor *executor.BSCExecutor,
	votePoolExecutor *vote.VotePoolExecutor, voteProcessor *vote.InscriptionVoteProcessor, inscriptionAssembler *assembler.InscriptionAssembler,
) *InscriptionRelayer {
	return &InscriptionRelayer{
		Listener:             listener,
		InscriptionExecutor:  inscriptionExecutor,
		bscExecutor:          bscExecutor,
		votePoolExecutor:     votePoolExecutor,
		voteProcessor:        voteProcessor,
		inscriptionAssembler: inscriptionAssembler,
	}
}

func (r *InscriptionRelayer) Start() {
	go r.MonitorCrossChainEvents()
	go r.SignAndBroadcast()
	go r.CollectVotes()
	go r.AssembleTransactions()
	go r.UpdateCachedLatestValidators()
}

// MonitorCrossChainEvents will monitor cross chain events for every block and persist into DB
func (r *InscriptionRelayer) MonitorCrossChainEvents() {
	r.Listener.Start()
}

func (r *InscriptionRelayer) SignAndBroadcast() {
	r.voteProcessor.SignAndBroadcast()
}

func (r *InscriptionRelayer) CollectVotes() {
	r.voteProcessor.CollectVotes()
}

func (r *InscriptionRelayer) AssembleTransactions() {
	r.inscriptionAssembler.AssembleTransactionAndSend()
}

func (r *InscriptionRelayer) UpdateCachedLatestValidators() {
	r.InscriptionExecutor.UpdateCachedLatestValidators() // cache validators queried from inscription, update it every 1 minute
}
