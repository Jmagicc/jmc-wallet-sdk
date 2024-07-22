package ecdh_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/Jmagicc/jmc-wallet-sdk/crypto"
	"github.com/Jmagicc/jmc-wallet-sdk/crypto/ecdh"
	"io/ioutil"
	"log"
	"math/big"
	"testing"
)

func Test_Secret2File(t *testing.T) {
	// 选择曲线，例如 P-256
	curve := elliptic.P256()

	// 生成私钥
	privateKey, err := ecdh.EcdhGenerateKey(rand.Reader, curve)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// 将公钥私钥分别写入到当前的文件夹下,分别命名为privateKey.key和PublicKey.key
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		fmt.Println("Error marshaling private key:", err)
		return
	}
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	err = ioutil.WriteFile("privateKey.key", privateKeyPEM, 0600)
	if err != nil {
		fmt.Println("Error writing private key to file:", err)
		return
	}

	// 生成公钥
	publicKey := privateKey.PublicKey

	// 将公钥写入文件
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		fmt.Println("Error marshaling public key:", err)
		return
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	err = ioutil.WriteFile("PublicKey.key", publicKeyPEM, 0644)
	if err != nil {
		fmt.Println("Error writing public key to file:", err)
		return
	}

	// 文件写入完成
	fmt.Println("Private and public keys are written to files: privateKey.key, PublicKey.key")
}

func Test_File2Secret(t *testing.T) {
	pub, err := ecdh.LoadPublicKeyFromFile("PublicKey.key")
	if err != nil {
		fmt.Println("Error loading public key:", err)
		return
	}
	//pri, err := LoadPrivateKeyFromFile("privateKey.key")
	//if err != nil {
	//	fmt.Println("Error loading private key:", err)
	//	return
	//}
	//key := pri.Public()

	fmt.Println("Public key is:", pub)
	st := "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzmPGt35Sw8YsW+o/9jx8DfwkCpVZ\nC3kKCAkGIge7Rj/W94sFyeHDZPN40B01by4QCEUYTEGL7H4fhoqAdGbXrA==\n-----END PUBLIC KEY-----"
	key2String, err := ecdh.LoadString2PublicKey(st)
	if pub.Equal(key2String) {
		fmt.Println("Keys match")
	} else {
		fmt.Println("match failed")
	}

	// 判断 key 和 pub 是否相同
	//if pub.Equal(key) {
	//	fmt.Println("Keys match")
	//} else {
	//	fmt.Println("match failed")
	//}

}

func Test_Ecdh(t *testing.T) {

	// 选择曲线，例如 P-256
	curve := elliptic.P256()

	// 生成私钥 （服务端生成，讨论存储在哪，要过期时间）
	privateKeyA, err := ecdh.EcdhGenerateKey(rand.Reader, curve)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}
	// 将私钥以十六进制打印出来
	privateKeyHex := hex.EncodeToString(privateKeyA.D.Bytes())
	fmt.Printf("服务端私钥: %s\n", privateKeyHex)

	// 生成公钥
	publicKeyA := &privateKeyA.PublicKey

	// 将公钥以十六进制打印出来
	publicKeyBytes := elliptic.Marshal(curve, publicKeyA.X, publicKeyA.Y)
	publicKeyHex := hex.EncodeToString(publicKeyBytes)
	fmt.Printf("服务端公钥: %s\n", publicKeyHex)

	publickeyB := Hex2Publickey("043e5971cd28f698ac66189eda1567d4d3952e19fe1513b6f404e776f25bc07a919e68aaade64ba7cc9079730ef8dbf195241351cdefd7ab641bab1ca8286f437f")

	fmt.Println("publickeyB:", publickeyB)
	//sharekey, err := ecdh.EcdhSharedKey(privateKeyA, publickeyB)
	//fmt.Println("sharekey:", sharekey)

}

func Hex2Publickey(publicKeyHex string) *ecdsa.PublicKey {
	// 解码为字节表示
	publicKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		log.Fatal("Error decoding public key:", err)
	}

	// 提取X和Y坐标
	curve := elliptic.P256()
	x, y := elliptic.Unmarshal(curve, publicKeyBytes)

	// 创建*ecdsa.PublicKey对象
	publicKey := &ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}

	// 打印公钥对象
	fmt.Printf("公钥: %+v\n", publicKey)
	return publicKey
}

func Test_Two(t *testing.T) {

	serverPub := "04fa9749f272fcb8f6d19649efffd0225d6d75c5846f3ff1648528d8907bbfb7316300b54a07030ede0c08db9d7c7e3470be27e379ad8afcd138955a11ec8ca2a0"
	//serverPri := "d9f9194105a63c4a4c0ec4c2bc4655387154cb34d75cff121d855e60e80312ec"

	//kehuPub := "04e9ba726f4ccc644c376643ee058e7732dfff11981456257760be561c4d54ed4fe09f37db73de3de9f121ec3318171394721505ba92c2388147bafa84d2cc078d"
	kehuPri := "9be2fd76786c5d3ad2299bf7048d34804a1361e3cf6e5ef7c9abf59fc43fb2c7"

	Bors(serverPub, kehuPri)
	//Bors(kehuPub, serverPri)

	//fmt.Println("bors1:", bors1)
	//fmt.Println("bors2:", bors2)
	//
	//if bors1 == bors2 {
	//	fmt.Println("相同")
	//} else {
	//	fmt.Println("不相同")
	//}

}

func Bors(publicKeyHex, privateKeyHex string) string {
	// 以十六进制表示的公钥和私钥

	// 解码为字节表示
	publicKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		log.Fatal("Error decoding public key:", err)
	}
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		log.Fatal("Error decoding private key:", err)
	}

	// 创建*ecdsa.PublicKey对象
	curve := elliptic.P256()
	x, y := elliptic.Unmarshal(curve, publicKeyBytes)
	publicKey := &ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}

	// 创建*ecdsa.PrivateKey对象
	privateKey := new(ecdsa.PrivateKey)
	privateKey.Curve = curve
	privateKey.D = new(big.Int).SetBytes(privateKeyBytes)
	privateKey.PublicKey = *publicKey

	// 计算共享密钥
	sharedKey, err := ecdh.EcdhSharedKey(privateKey, publicKey)
	if err != nil {
		log.Fatal("Error calculating shared key:", err)
	}

	// 打印共享密钥
	sharedKeyHex := hex.EncodeToString(sharedKey)
	fmt.Println("共享密钥:", sharedKeyHex)

	return sharedKeyHex
}

func Test_Aes(t *testing.T) {
	ss := "Hello, AES!"
	key := []byte("0123456789abcdef")

	// aes加密
	encrypt, err := crypto.AesEncrypt([]byte(ss), key)
	if err != nil {
		panic(err)
	}
	fmt.Println("加密后", string(encrypt))
	decrypt, err := crypto.AesDecrypt(encrypt, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后", string(decrypt))

}
