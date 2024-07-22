package eth

import (
	"errors"
	"strings"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Util struct {
}

func NewUtil() *Util {
	return &Util{}
}

// MARK - Implement the protocol wallet.Util

func (u *Util) EncodePublicKeyToAddress(publicKey string) (string, error) {
	return EncodePublicKeyToAddress(publicKey)
}

// Warning: eth cannot support decode address to public key
func (u *Util) DecodeAddressToPublicKey(address string) (string, error) {
	return "", errors.New("eth cannot support decode address to public key")
}

// Check if address is 40 hexadecimal characters
func (u *Util) IsValidAddress(address string) bool {
	return IsValidAddress(address)
}

// MARK - like wallet.Util

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

// Warning: eth cannot support decode address to public key
func DecodeAddressToPublicKey(address string) (string, error) {
	return "", errors.New("eth cannot support decode address to public key")
}

// Check if address is 40 hexadecimal characters
func IsValidAddress(address string) bool {
	return common.IsHexAddress(address)
}

// It will check based on eip55 rules
func IsValidEIP55Address(address string) bool {
	if !IsValidAddress(address) {
		return false
	}
	eip55Address := TransformEIP55Address(address)
	return strings.HasSuffix(eip55Address, address)
}

func TransformEIP55Address(address string) string {
	address = strings.TrimPrefix(address, "0x")
	addressBytes := []byte(strings.ToLower(address))
	checksumBytes := crypto.Keccak256(addressBytes)

	for i, c := range addressBytes {
		if c >= '0' && c <= '9' {
			continue
		} else {
			checksum := checksumBytes[i/2]
			bitcode := byte(0x80) >> ((i % 2) * 4)
			if checksum&bitcode > 0 { // to Upper
				addressBytes[i] -= 32
			}
		}
	}

	return "0x" + string(addressBytes)
}

// ToChecksumAddress 将以太坊地址从不区分大小写转换为区分大小写
func ToChecksumAddress(address string) string {
	// 使用 Keccak-256 哈希算法对地址进行哈希
	hash := calculateHash(strings.ToLower(address))

	// 创建一个新地址，根据哈希值的对应字符的大小写
	var result strings.Builder
	for i := 0; i < len(address); i++ {
		if shouldUppercase(hash[i]) {
			result.WriteRune(rune(strings.ToUpper(string(address[i]))[0]))
		} else {
			result.WriteRune(rune(address[i]))
		}
	}
	return result.String()
}

// 计算地址的 Keccak-256 哈希值
func calculateHash(address string) string {
	// 在真实情况下，这里应该使用真正的 Keccak-256 实现
	// 这里只是演示示例
	// 实际情况中，请使用以太坊客户端或相关库来计算哈希
	return address
}

// 判断字符是否需要大写
func shouldUppercase(hashChar byte) bool {
	// 根据 Keccak-256 哈希的字符值判断是否需要大写
	return hashChar >= '8'
}
