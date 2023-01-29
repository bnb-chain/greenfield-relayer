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
	go r.MonitorEventsLoop()
	go r.SignAndBroadcastLoop()
	go r.CollectVotesLoop()
	go r.AssembleTransactionsLoop()
	go r.UpdateCachedLatestValidatorsLoop()
}

// MonitorEventsLoop will monitor cross chain events for every block and persist into DB
func (r *InscriptionRelayer) MonitorEventsLoop() {
	r.Listener.Start()
}

func (r *InscriptionRelayer) SignAndBroadcastLoop() {
	r.voteProcessor.SignAndBroadcast()
}

func (r *InscriptionRelayer) CollectVotesLoop() {
	r.voteProcessor.CollectVotes()
}

func (r *InscriptionRelayer) AssembleTransactionsLoop() {
	r.inscriptionAssembler.AssembleTransactionAndSend()
}

func (r *InscriptionRelayer) UpdateCachedLatestValidatorsLoop() {
	r.InscriptionExecutor.UpdateCachedLatestValidators() // cache validators queried from inscription, update it every 1 minute
}
func (r *InscriptionRelayer) UpdateClientLoop() {
	r.InscriptionExecutor.UpdateClients()
}
