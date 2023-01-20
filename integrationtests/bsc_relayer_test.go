package integrationtests

import (
	"encoding/hex"
	"math/big"
	"sort"
	"testing"
	"time"

	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/util"
	"github.com/bnb-chain/inscription-relayer/vote"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prysmaticlabs/prysm/crypto/bls/blst"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/votepool"
)

const (
	// Change your relayers Bls private key when integration test using local env
	Relayer1HexBlsPrivKey = "2969268e6722a8e16579e7a3380f83a2dd0b15478a2994cb0ac6480e1aead999" // for test only
	Relayer2HexBlsPrivKey = "6f235c2c0d91ecdf961f4409061a785d456b9bc4b398e2a0940378397772cb0b"
)

func TestClaimPackagesSucceed(t *testing.T) {
	app := InitTestApp()
	go app.BSCRelayer.SignAndBroadcast()
	go app.BSCRelayer.CollectVotes()
	go app.BSCRelayer.AssemblePackages()

	inscriptionExecutor := app.BSCRelayer.InscriptionExecutor
	daoManager := app.BSCRelayer.Listener.DaoManager

	// Given: Prepare cross-chain packages to be sent. Define the channel id, oracle sequence and package sequence are
	// retrieved from destination chain(Inscription),
	channelId := uint8(1)
	oracleSeq, err := inscriptionExecutor.GetNextReceiveOracleSequence()
	require.NoError(t, err)
	packageStartSeq, err := inscriptionExecutor.GetNextReceiveSequenceForChannel(relayercommon.ChannelId(channelId))
	require.NoError(t, err)

	relayPkgs := make([]*model.BscRelayPackage, 0)
	ts := time.Now().Unix()

	packagesSize := uint64(2)
	endSeq := packageStartSeq + (packagesSize - 1)

	// TODO Use cli to send cross-chain tx from BSC so that there is no need to mock votes
	txIndex := 0
	for i := packageStartSeq; i <= endSeq; i++ {
		relayPkg := model.BscRelayPackage{}
		relayPkg.ChannelId = channelId
		relayPkg.OracleSequence = oracleSeq
		relayPkg.PackageSequence = i
		relayPkg.PayLoad = hex.EncodeToString(getPayload(uint64(ts)))
		relayPkg.Height = 1
		relayPkg.TxHash = "testHash"
		relayPkg.TxIndex = uint(txIndex)
		relayPkg.Status = db.Saved
		relayPkg.TxTime = ts
		relayPkg.UpdatedTime = ts
		relayPkgs = append(relayPkgs, &relayPkg)
		txIndex++
	}

	// When: Save packages with status 'Saved' into Database, there are processes to aggregate packages to submit vote to
	// Inscription Votepool, gathering votes from votepool, and assembler them to claim in Inscription.
	err = daoManager.BSCDao.SaveBatchPackages(relayPkgs)
	require.NoError(t, err)

	// This is needed in local testing, if move to use testnet, can trigger transaction use cli.
	go broadcastVotesFromOtherRelayers(daoManager, app.BSCRelayer.VotePoolExecutor, oracleSeq)

	// The first in-turn relayer has 40 seconds relaying window, so that need to wait for a  while if current one is not the first in-turn.
	// sleep for all processes have done their work, if there are more validators, might need to set this larger due to
	// it takes time for current relayer to be in-turn.
	time.Sleep(60 * time.Second)

	// Then:  the oracle sequence is filled, sequences for cross-chain packages are all filled
	pkgs, err := daoManager.BSCDao.GetPackagesByStatus(db.Delivered)
	require.NoError(t, err)
	sort.Slice(pkgs, func(i, j int) bool {
		return pkgs[i].PackageSequence < pkgs[j].PackageSequence
	})
	require.EqualValues(t, packageStartSeq, pkgs[0].PackageSequence)
	require.EqualValues(t, endSeq, pkgs[packagesSize-1].PackageSequence)

	nextSeq, err := inscriptionExecutor.GetNextReceiveSequenceForChannel(relayercommon.ChannelId(channelId))
	require.NoError(t, err)
	require.EqualValues(t, endSeq+1, nextSeq)
	nextOracleSeq, err := inscriptionExecutor.GetNextReceiveOracleSequence()
	require.NoError(t, err)
	require.EqualValues(t, oracleSeq+1, nextOracleSeq)
}

func getPayload(ts uint64) []byte {
	payloadHeader := sdk.EncodePackageHeader(sdk.PackageHeader{
		PackageType:   sdk.SynCrossChainPackageType,
		Timestamp:     ts,
		RelayerFee:    big.NewInt(1),
		AckRelayerFee: big.NewInt(1),
	})
	payloadHeader = append(payloadHeader, []byte("test payload")...)
	return payloadHeader
}

func broadcastVotesFromOtherRelayers(daoManager *dao.DaoManager,
	votePoolExecutor *vote.VotePoolExecutor, oracleSeq uint64) {
	var vote *model.Vote
	// retry to query vote from local DB
	for {
		localVote, err := daoManager.VoteDao.GetVoteByChannelIdAndSequenceAndPubKey(uint8(relayercommon.OracleChannelId), oracleSeq, hex.EncodeToString(util.GetBlsPubKeyFromPrivKeyStr(GetTestConfig().VotePoolConfig.BlsPrivateKey)))
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		vote = localVote
		break
	}

	secretKey1, err := blst.SecretKeyFromBytes(common.Hex2Bytes(Relayer1HexBlsPrivKey))
	if err != nil {
		panic(err)
	}
	pubKey1 := secretKey1.PublicKey()
	sign1 := secretKey1.Sign(vote.EventHash[:]).Marshal()

	mockVoteFromRelayer1 := &votepool.Vote{
		PubKey:    pubKey1.Marshal(),
		Signature: sign1,
		EventType: 2,
		EventHash: vote.EventHash[:],
	}

	secretKey2, err := blst.SecretKeyFromBytes(common.Hex2Bytes(Relayer2HexBlsPrivKey))
	if err != nil {
		panic(err)
	}
	pubKey2 := secretKey2.PublicKey()
	sign2 := secretKey2.Sign(vote.EventHash[:]).Marshal()

	mockVoteFromRelayer2 := &votepool.Vote{
		PubKey:    pubKey2.Marshal(),
		Signature: sign2,
		EventType: 2,
		EventHash: vote.EventHash[:],
	}
	_ = votePoolExecutor.BroadcastVote(mockVoteFromRelayer1)
	_ = votePoolExecutor.BroadcastVote(mockVoteFromRelayer2)
}
