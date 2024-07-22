package wallet

import (
	"fmt"
	"github.com/Jmagicc/jmc-wallet-sdk/core/eth"
	"io/ioutil"
	"log"
	"reflect"
	"testing"
)

func TestGenMnemonic(t *testing.T) {
	mnemonic, err := GenMnemonic()
	if err != nil {
		t.Errorf("Error generating mnemonic: %v", err)
	}

	if !IsValidMnemonic(mnemonic) {
		t.Errorf("Generated mnemonic is not valid")
	}
}

func TestSplitAndCombineShares(t *testing.T) {
	mnemonic, err := GenMnemonic()
	if err != nil {
		log.Fatalln(err, "GenMnemonic")
		return
	}
	fmt.Println(mnemonic)
	account, err := eth.NewAccountWithMnemonic(mnemonic)
	if err != nil {
		log.Fatalln(err, "NewAccountWithMnemonic")
		return
	}
	privateKey, err := account.PrivateKeyHex()
	if err != nil {
		log.Fatalln(err, "account.PrivateKeyHex()")
		return
	}

	// 开始分割
	minimumShares := 3
	totalShares := 3

	shares, err := SplitToShares(privateKey, minimumShares, totalShares)

	if err != nil {
		t.Errorf("Error splitting private key: %v", err)
	}

	combinedPrivateKey, err := CombineShares(shares[:minimumShares])
	if err != nil {
		t.Errorf("Error combining shares: %v", err)
	}

	if !reflect.DeepEqual(privateKey, string(combinedPrivateKey)) {
		fmt.Println("privateKey:", privateKey)
		fmt.Println("combinedPrivateKey:", string(combinedPrivateKey))
		t.Errorf("Combined private key does not match original private key ByEncodeToString")
	}

}

// 从文件中读取分割的助记词
func TestCombineShares(t *testing.T) {
	dir := "E:/haijiajia/github.com/Jmagicc/jmc-wallet-sdk/sharesDir/339297526983163904/0x916bAC8FB857ae7e5cddf54DB27dd005621242fC/Z"
	byteAll := [][]byte{}
	for i := 1; i < 4; i++ {
		filename := fmt.Sprintf("%s/%d.txt", dir, i)
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("Error reading file: %v", err)
			return
		}
		byteAll = append(byteAll, content)
	}
	shares, err := CombineShares(byteAll)
	if err != nil {
		t.Errorf("Error combining shares: %v", err)
		return
	}
	fmt.Println("shares:", string(shares))
}
