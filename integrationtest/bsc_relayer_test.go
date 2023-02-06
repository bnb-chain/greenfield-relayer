package integrationtest

import (
	"encoding/hex"
	"github.com/bnb-chain/greenfield-relayer/types"
	"math/big"
	"sort"
	"testing"
	"time"

	"github.com/bnb-chain/greenfield-relayer/db"
	"github.com/bnb-chain/greenfield-relayer/db/model"
	"github.com/bnb-chain/greenfield-relayer/vote"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	// Change your relayers Bls private key when integration test using local env
	Relayer1HexBlsPrivKey = "3a1055a667eddef7405a554f2994aedea43c8258712a013b3a61532e8cd0f032" // for test only
	Relayer2HexBlsPrivKey = "139ace9a52fa78b9f4dc2f151231225f9503d60f3aefeef89a1ee82d6d48ef9a"
)

func TestClaimPackagesSucceed(t *testing.T) {
	app := InitTestApp()
	go app.BSCRelayer.SignAndBroadcastVoteLoop()
	go app.BSCRelayer.CollectVotesLoop()
	go app.BSCRelayer.AssemblePackagesLoop()

	greenfieldExecutor := app.BSCRelayer.GreenfieldExecutor
	daoManager := app.BSCRelayer.Listener.DaoManager

	// Given: Prepare cross-chain packages to be sent. Define the channel id, oracle sequence and package sequence are
	// retrieved from destination chain(Greenfield),
	channelId := uint8(1)
	oracleSeq, err := greenfieldExecutor.GetNextReceiveOracleSequence()
	require.NoError(t, err)
	packageStartSeq, err := greenfieldExecutor.GetNextReceiveSequenceForChannel(types.ChannelId(channelId))
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
	// Greenfield Votepool, gathering votes from votepool, and assembler them to claim in Greenfield.
	err = daoManager.BSCDao.SaveBatchPackages(relayPkgs)
	require.NoError(t, err)

	// This is needed in local testing, if move to use testnet, can trigger transaction use cli.
	go vote.BroadcastVotesFromOtherRelayers(
		[]string{Relayer1HexBlsPrivKey, Relayer2HexBlsPrivKey},
		daoManager, app.BSCRelayer.VotePoolExecutor, channelId, oracleSeq, 2)

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

	nextSeq, err := greenfieldExecutor.GetNextReceiveSequenceForChannel(types.ChannelId(channelId))
	require.NoError(t, err)
	require.EqualValues(t, endSeq+1, nextSeq)
	nextOracleSeq, err := greenfieldExecutor.GetNextReceiveOracleSequence()
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
