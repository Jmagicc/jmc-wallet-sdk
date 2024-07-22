package ecdh

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

// 用AES对数据进行加密(AES的秘钥是由ECDH生成的共享密钥), 用ECDH生成共享密钥进行数据传输
// 1.登录成功后，客户端生成ECDH密钥对，将公钥发送给服务端（存哪需要讨论，要过期时间）
// 2.登录成功后，服务端生成ECDH密钥对，将公钥发送给客户端（客户端可以存SessionStorage）
// 3.客户端使用自己的私钥和服务端的公钥计算出共享密钥，服务端使用自己的私钥和客户端的公钥计算出共享密钥
// 4.后续一些数据传输，客户端使用共享密钥对数据进行加密，发送给服务端。服务端用共享密钥对数据进行AES解密,得到明文数据

func main() {
	// 选择曲线，例如 P-256
	curve := elliptic.P256()

	// 生成私钥 （服务端生成，讨论存储在哪，要过期时间）
	privateKeyA, err := EcdhGenerateKey(rand.Reader, curve)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// 模拟客户端代码
	privateKeyB, err := EcdhGenerateKey(rand.Reader, curve)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// 获取公钥
	publicKeyA := &privateKeyA.PublicKey
	publicKeyB := &privateKeyB.PublicKey

	// 计算共享密钥
	sharedKeyA, err := EcdhSharedKey(privateKeyA, publicKeyB)
	if err != nil {
		fmt.Println("Error computing shared key:", err)
		return
	}

	sharedKeyB, err := EcdhSharedKey(privateKeyB, publicKeyA)
	if err != nil {
		fmt.Println("Error computing shared key:", err)
		return
	}

	// Output:
	if string(sharedKeyA) == string(sharedKeyB) {
		fmt.Println("ECDH shared key generation successful!")
	} else {
		fmt.Println("ECDH shared key generation failed!")
	}
}

// EcdhGenerateKey 生成ECDH密钥对
func EcdhGenerateKey(rand io.Reader, curve elliptic.Curve) (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(curve, rand)
}

// EcdhSharedKey 计算ECDH共享密钥
func EcdhSharedKey(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, error) {
	x, _ := publicKey.Curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())
	sharedKey := x.Bytes()
	return sharedKey[:32], nil // 选择共享密钥的前32个字节作为实际密钥
}

func LoadPrivateKeyFromFile(filepath string) (*ecdsa.PrivateKey, error) {
	// 从文件中读取 PEM 格式的私钥数据
	pemData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	// 解码 PEM 数据
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("invalid private key data")
	}

	// 解析 DER 格式的私钥
	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return key, nil
}

func LoadPublicKeyFromFile(filepath string) (*ecdsa.PublicKey, error) {
	// 从文件中读取 PEM 格式的公钥数据
	pemData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key file: %w", err)
	}

	// 解码 PEM 数据
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("invalid public key data")
	}

	// 解析 DER 格式的公钥
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	// 将解析的公钥转换为 *ecdsa.PublicKey 类型
	ecdsaKey, ok := key.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to convert public key to *ecdsa.PublicKey")
	}

	return ecdsaKey, nil
}

func LoadString2PublicKey(pemData string) (*ecdsa.PublicKey, error) {
	// 解码 PEM 数据
	block, _ := pem.Decode([]byte(pemData))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("invalid public key data")
	}

	// 解析 DER 格式的公钥
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	// 将解析的公钥转换为 *ecdsa.PublicKey 类型
	ecdsaKey, ok := key.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to convert public key to *ecdsa.PublicKey")
	}

	return ecdsaKey, nil
}

func LoadString2PrivateKey(pemData string) (*ecdsa.PrivateKey, error) {
	// 解码 PEM 数据
	block, _ := pem.Decode([]byte(pemData))
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("invalid private key data")
	}

	// 解析 DER 格式的私钥
	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return key, nil
}

// 计算共享密钥
func EcdhSharedKeyByString(client, server string) ([]byte, error) {
	privateKeyServer, err := LoadString2PrivateKey(server)
	if err != nil {
		return nil, errors.New("Error loading private key: " + err.Error())
	}
	publicKeyClient, err := LoadString2PublicKey(client)
	if err != nil {
		return nil, errors.New("Error loading public key: " + err.Error())
	}
	sharedKeyServer, err := EcdhSharedKey(privateKeyServer, publicKeyClient)
	if err != nil {
		return nil, errors.New("Error computing shared key: " + err.Error())
	}
	return sharedKeyServer, nil
}
