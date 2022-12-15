package vote

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/crypto/bls"

	validatorpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1/validator-client"
	"github.com/prysmaticlabs/prysm/validator/accounts/iface"
	"github.com/prysmaticlabs/prysm/validator/accounts/wallet"
	"github.com/prysmaticlabs/prysm/validator/keymanager"
	"inscription-relayer/common"
)

type VoteSigner struct {
	km     *keymanager.IKeymanager
	pubKey [48]byte
}

func NewVoteSigner(blsPasswordPath, blsWalletPath string) (*VoteSigner, error) {
	dirExists, err := wallet.Exists(blsWalletPath)
	if err != nil {
		common.Logger.Error("Check BLS wallet exists error: %v.", err)
		return nil, err
	}
	if !dirExists {
		common.Logger.Error("BLS wallet not exist.")
		return nil, fmt.Errorf("BLS wallet not exist")
	}

	walletPassword, err := os.ReadFile(blsPasswordPath)
	if err != nil {
		common.Logger.Error("Read BLS wallet password error: %v.", err)
		return nil, err
	}
	common.Logger.Info("Read BLS wallet password successfully")

	w, err := wallet.OpenWallet(context.Background(), &wallet.Config{
		WalletDir:      blsWalletPath,
		WalletPassword: string(walletPassword),
	})
	if err != nil {
		common.Logger.Error("Open BLS wallet failed: %v.", err)
		return nil, err
	}
	common.Logger.Info("Open BLS wallet successfully")

	km, err := w.InitializeKeymanager(context.Background(), iface.InitKeymanagerConfig{ListenForChanges: false})
	if err != nil {
		common.Logger.Error("Initialize key manager failed: %v.", err)
		return nil, err
	}
	common.Logger.Info("Initialized keymanager successfully")

	ctx, cancel := context.WithTimeout(context.Background(), VoteSignerTimeout)
	defer cancel()

	pubKeys, err := km.FetchValidatingPublicKeys(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch validating public keys")
	}

	return &VoteSigner{
		km:     &km,
		pubKey: pubKeys[0],
	}, nil
}

// SignVote sign a vote, data is used to signed to generate the signature
func (signer *VoteSigner) SignVote(vote *Vote, data []byte) error {
	pubKey := signer.pubKey
	blsPubKey, err := bls.PublicKeyFromBytes(pubKey[:])
	if err != nil {
		return errors.Wrap(err, "convert public key from bytes to vote failed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), VoteSignerTimeout)
	defer cancel()
	signature, err := (*signer.km).Sign(ctx, &validatorpb.SignRequest{
		PublicKey:   pubKey[:],
		SigningRoot: data[:],
	})

	if err != nil {
		return err
	}

	copy(vote.EventHash[:], data[:])
	copy(vote.PubKey[:], blsPubKey.Marshal()[:])
	copy(vote.Signature[:], signature.Marshal()[:])

	return nil
}
