package tron

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shengdoushi/base58"
)

type Util struct {
}

func NewUtil() *Util {
	return &Util{}
}

func (u *Util) EncodePublicKeyToAddress(publicKey string) (string, error) {
	return EncodePublicKeyToAddress(publicKey)
}

func (u *Util) DecodeAddressToPublicKey(address string) (string, error) {
	return "", errors.New("eth cannot support decode address to public key")
}

// Check if address is 40 hexadecimal characters
func (u *Util) IsValidAddress(address string) bool {
	return IsValidAddress(address)
}

func EncodePublicKeyToAddress(publicKey string) (string, error) {
	bytes, err := types.HexDecodeString(publicKey)
	if err != nil {
		return "", err
	}
	publicKeyECDSA, err := crypto.UnmarshalPubkey(bytes)
	if err != nil {
		return "", err
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address.String(), nil
}

func DecodeAddressToPublicKey(address string) (string, error) {
	return "", errors.New("eth cannot support decode address to public key")
}

// IsValidAddress Check if address is 34 hexadecimal characters
func IsValidAddress(address string) bool {
	if len(address) != 34 {
		return false
	}
	if string(address[0:1]) != "T" {
		return false
	}
	_, err := DecodeCheck(address)
	if err != nil {
		return false
	}
	return true
}

func DecodeCheck(input string) ([]byte, error) {
	var tronAlphabet = base58.NewAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	decodeCheck, err := base58.Decode(input, tronAlphabet)

	if err != nil {
		return nil, err
	}
	if len(decodeCheck) < 4 {
		return nil, fmt.Errorf("addres base58 not check ok")
	}

	decodeData := decodeCheck[:len(decodeCheck)-4]

	h256h0 := sha256.New()
	h256h0.Write(decodeData)
	h0 := h256h0.Sum(nil)

	h256h1 := sha256.New()
	h256h1.Write(h0)
	h1 := h256h1.Sum(nil)

	if h1[0] == decodeCheck[len(decodeData)] &&
		h1[1] == decodeCheck[len(decodeData)+1] &&
		h1[2] == decodeCheck[len(decodeData)+2] &&
		h1[3] == decodeCheck[len(decodeData)+3] {
		return decodeData, nil
	}

	return nil, fmt.Errorf("addres hash not check ok")
}
