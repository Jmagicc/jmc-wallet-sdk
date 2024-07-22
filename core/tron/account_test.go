package tron

import (
	"fmt"
	"github.com/stretchr/testify/assert"

	"github.com/Jmagicc/jmc-wallet-sdk/core/wallet"
	"log"
	"testing"
)

// Test_CreateNewAccount 测试创建新账户(波场)
func Test_CreateNewAccount(t *testing.T) {
	mnemonic, err := wallet.GenMnemonic()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(mnemonic)
	account, err := NewAccountWithMnemonic(mnemonic)
	if err != nil {
		log.Fatalln(err)
	}
	pri, err := account.PrivateKeyHex()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("助记词:", mnemonic)
	log.Println("公钥:", account.Address())
	log.Println("私钥:", pri)
	// 验证地址是否是合法主网认证的
	if assert.Equal(t, true, account.Util.IsValidAddress(account.Address())) {
		log.Println("地址合法")
	} else {
		log.Println("地址不合法")
	}
}

// Test_CreateNewAccountWithPrivateKey 测试创建新账户(波场)
func Test_CreateNewAccountWithPrivateKey(t *testing.T) {
	//2023/08/12 23:51:08 助记词: hospital weather flash small proud oval stock conduct duck steak embody neither
	//2023/08/12 23:51:08 公钥: TWA5cjj25SD5o4PdLm4j7fbABAEkFZWC  G1
	//2023/08/12 23:51:08 私钥: fa232170486bda9ad49b000e52444b73319b65bf75e26586a96e244c41a84847
	privateKey := "fa232170486bda9ad49b000e52444b73319b65bf75e26586a96e244c41a84847"

	account, err := AccountWithPrivateKey(privateKey)
	if err != nil {
		log.Fatalln(err)
	}
	pri, err := account.PrivateKeyHex()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("公钥:", account.Address())
	log.Println("私钥:", pri)

}
