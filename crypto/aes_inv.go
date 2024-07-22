package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func EncryptByinv(key, iv []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	paddedPlaintext := padPKCS7([]byte(plaintext), block.BlockSize())
	ciphertext := make([]byte, len(paddedPlaintext))

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(ciphertext, paddedPlaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptByinv(key, iv []byte, ciphertext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	cbc := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(cipherBytes))
	cbc.CryptBlocks(decrypted, cipherBytes)

	unpaddedPlaintext, err := unpadPKCS7(decrypted, block.BlockSize())
	if err != nil {
		return "", err
	}

	return string(unpaddedPlaintext), nil
}

func padPKCS7(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	paddedData := append(data, bytesRepeatingByte(padding, byte(padding))...)
	return paddedData
}

func unpadPKCS7(data []byte, blockSize int) ([]byte, error) {
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, errors.New("Invalid padding")
	}

	padding := int(data[len(data)-1])
	if padding > blockSize || padding == 0 {
		return nil, errors.New("Invalid padding")
	}

	for i := len(data) - padding; i < len(data); i++ {
		if data[i] != byte(padding) {
			return nil, errors.New("Invalid padding")
		}
	}

	return data[:len(data)-padding], nil
}

func bytesRepeatingByte(length int, b byte) []byte {
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = b
	}
	return bytes
}
