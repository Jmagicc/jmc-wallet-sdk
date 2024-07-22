package tron

import (
	"encoding/hex"
	"github.com/Jmagicc/jmc-wallet-sdk/core/base"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/accounts"
	addr "github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/tyler-smith/go-bip39"
)

type Account struct {
	*Util
	privateKey *btcec.PrivateKey
	address    string
}

// NewAccountWithMnemonic creates a new account with a given mnemonic
func NewAccountWithMnemonic(mnemonic string) (*Account, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, err
	}

	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}
	path, err := accounts.ParseDerivationPath("m/44'/195'/0'/0/0")
	if err != nil {
		return nil, err
	}

	key := masterKey
	for _, n := range path {
		key, err = key.DeriveNonStandard(n)
		if err != nil {
			return nil, err
		}
	}

	privateKey, err := key.ECPrivKey()
	if err != nil {
		return nil, err
	}

	address := addr.PubkeyToAddress(privateKey.ToECDSA().PublicKey)

	return &Account{
		Util:       NewUtil(),
		privateKey: privateKey,
		address:    address.String(),
	}, nil
}

// AccountWithPrivateKey creates a new account with a given private key
func AccountWithPrivateKey(privatekey string) (*Account, error) {
	priData, err := types.HexDecodeString(privatekey)
	if err != nil {
		return nil, err
	}

	privateKey, _ := btcec.PrivKeyFromBytes(priData)

	address := addr.PubkeyToAddress(privateKey.ToECDSA().PublicKey)

	return &Account{
		Util:       NewUtil(),
		privateKey: privateKey,
		address:    address.String(),
	}, nil
}

func (a Account) PrivateKey() ([]byte, error) {
	return a.privateKey.Serialize(), nil
}

func (a Account) PrivateKeyHex() (string, error) {
	return hex.EncodeToString(a.privateKey.Serialize()), nil
}

func (a Account) PublicKey() []byte {
	return a.privateKey.PubKey().SerializeCompressed()
}

func (a Account) PublicKeyHex() string {
	return hex.EncodeToString(a.privateKey.PubKey().SerializeCompressed())
}

func (a Account) Address() string {
	return a.address
}

func (a Account) Sign(message []byte, password string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (a Account) SignHex(messageHex string, password string) (*base.OptionalString, error) {
	//TODO implement me
	panic("implement me")
}
