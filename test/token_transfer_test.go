package test

import (
	"encoding/json"
	"fmt"
	"github.com/Jmagicc/jmc-wallet-sdk/util"
	"sync/atomic"
	"testing"
)

type TransferRequest struct {
	ToAddress string `json:"toAddress"`
	Amount    string `json:"amount"`
	Method    string `json:"method"`
}

func Test_transfer(t *testing.T) {
	transfer := TransferRequest{
		ToAddress: "0x178a8AB44b71858b38Cc68f349A06f397A73bFf5",
		Amount:    "10000000",
		Method:    "transfer",
	}

	transferJSON, err := json.Marshal(transfer)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return
	}

	fmt.Println(string(transferJSON))
}

func TestCount(t *testing.T) {
	count := int64(0)

	// 原子性地将 count 递增 1，返回递增后的值
	newCount := atomic.AddInt64(&count, 1)

	// 打印递增后的值
	fmt.Println(newCount)
	fmt.Println(count)

	atomic.AddInt64(&count, 1)
	atomic.AddInt64(&count, 1)
	atomic.AddInt64(&count, 1)
	fmt.Println(count)

}

func Test_(t *testing.T) {
	balance, err := util.ShiftBalance("10000", "1")

	fmt.Println(balance)
	fmt.Println(err)
}
