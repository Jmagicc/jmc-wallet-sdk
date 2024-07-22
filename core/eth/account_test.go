package eth

import (
	"fmt"
	"github.com/Jmagicc/jmc-wallet-sdk/core/wallet"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"log"
	"testing"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type SecretFileInfo struct {
	Id    int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Uid   string `gorm:"column:uid" json:"uid"`
	Path1 string `gorm:"column:path1" json:"path1"`
	Path2 string `gorm:"column:path2" json:"path2"`
	Path3 string `gorm:"column:path3" json:"path3"`
	BaseModel
}

func (s SecretFileInfo) TableName() string {
	return "secret_fileinfo"
}

// Test_CreateNewAccount 测试创建新账户
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
}

func TestSignAndVerify(t *testing.T) {
	tempMnemonic := "candy maple cake sugar pudding cream honey rich smooth crumble sweet treat"
	account, _ := NewAccountWithMnemonic(tempMnemonic)

	message := "fjppdipsidjsosososofdafjiowewsosap"
	msgbytes := []byte(message)
	messageHex := types.HexEncodeToString(msgbytes)

	signedString, err := account.SignHex(messageHex, "")
	if err != nil {
		t.Fatal(err)
	}

	// ============================================
	valid := VerifySignature(account.PublicKeyHex(), messageHex, signedString.Value)
	if valid {
		t.Log("Sign & Verify succeed!")
	} else {
		t.Fatal("Sign & Verify failured.")
	}
}
