package wallet

import (
	"fmt"
	"github.com/hashicorp/vault/shamir"
	"github.com/tyler-smith/go-bip39"
)

// GenMnemonic 生成熵128位的助记词
func GenMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	return mnemonic, err
}

// IsValidMnemonic 判断助记词是否有效
func IsValidMnemonic(mnemonic string) bool {
	_, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	return err == nil
}

// SplitToShares shamir分割私钥
func SplitToShares(privateKey string, minimumShares int, totalShares int) ([][]byte, error) {
	if privateKey == "" {
		return nil, fmt.Errorf("private key cannot be nil")
	}
	if minimumShares < 1 || minimumShares > totalShares {
		return nil, fmt.Errorf("invalid share count: minimum %d, total %d", minimumShares, totalShares)
	}

	shares, err := shamir.Split([]byte(privateKey), totalShares, minimumShares)
	if err != nil {
		return nil, fmt.Errorf("failed to split private key: %v", err)
	}
	return shares, nil
}

// CombineShares shamir合并私钥
func CombineShares(shares [][]byte) ([]byte, error) {
	return shamir.Combine(shares)
}
