package tron

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/Jmagicc/jmc-wallet-sdk/core/base"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/golang/protobuf/proto"
	"math/big"
	"strconv"
	"strings"
	"time"
)

// @title    主网代币余额查询
// @description   返回主网代币余额，decimal为代币精度
// @param     (walletAddress)     (string)  合约名称，钱包地址
// @return    (string,error)       代币余额，错误信息
func (e *TronChain) Balance(address string) (string, error) {
	err := e.keepConnect()
	if err != nil {
		return "", err
	}
	account, err := e.RemoteRpcClient.GetAccount(address)
	if err != nil {
		return "0", base.MapAnyToBasicError(err)
	}
	return strconv.FormatInt(account.Balance, 10), nil
}

// 获取最新区块高度
func (e *TronChain) LatestBlockNumber() (string, error) {
	err := e.keepConnect()
	if err != nil {
		return "", err
	}
	number, err := e.RemoteRpcClient.GetNowBlock()
	if err != nil {
		return "0", base.MapAnyToBasicError(err)
	}

	return number.String(), nil
}

/*
保持连接，如果中途连接失败，就重连
*/
func (e *TronChain) keepConnect() error {
	_, err := e.RemoteRpcClient.GetNodeInfo()
	if err != nil {
		if strings.Contains(err.Error(), "no such host") {
			return e.RemoteRpcClient.Reconnect(e.rpcUrl)
		}
		return fmt.Errorf("node connect error: %v", err)
	}
	return nil
}

func (e *TronChain) Transfer(from, to string, amount int64) (*api.TransactionExtention, error) {
	err := e.keepConnect()
	if err != nil {
		return nil, err
	}
	return e.RemoteRpcClient.Transfer(from, to, amount)
}

func (e *TronChain) GetTrc10Balance(addr, assetId string) (int64, error) {
	err := e.keepConnect()
	if err != nil {
		return 0, err
	}
	acc, err := e.RemoteRpcClient.GetAccount(addr)
	if err != nil || acc == nil {
		return 0, fmt.Errorf("get %s account error: %v", addr, err)
	}
	for key, value := range acc.AssetV2 {
		if key == assetId {
			return value, nil
		}
	}
	return 0, fmt.Errorf("%s do not find this assetID=%s amount", addr, assetId)
}

func (e *TronChain) GetTrxBalance(addr string) (*core.Account, error) {
	err := e.keepConnect()
	if err != nil {
		return nil, err
	}
	return e.RemoteRpcClient.GetAccount(addr)
}
func (e *TronChain) GetTrc20Balance(addr, contractAddress string) (*big.Int, error) {
	err := e.keepConnect()
	if err != nil {
		return nil, err
	}
	return e.RemoteRpcClient.TRC20ContractBalance(addr, contractAddress)
}

func (e *TronChain) TransferTrc10(from, to, assetId string, amount int64) (*api.TransactionExtention, error) {
	err := e.keepConnect()
	if err != nil {
		return nil, err
	}
	fromAddr, err := address.Base58ToAddress(from)
	if err != nil {
		return nil, fmt.Errorf("from address is not equal")
	}
	toAddr, err := address.Base58ToAddress(to)
	if err != nil {
		return nil, fmt.Errorf("to address is not equal")
	}
	return e.RemoteRpcClient.TransferAsset(fromAddr.String(), toAddr.String(), assetId, amount)
}

func (e *TronChain) TransferTrc20(from, to, contract string, amount *big.Int, feeLimit int64) (*api.TransactionExtention, error) {
	err := e.keepConnect()
	if err != nil {
		return nil, err
	}
	return e.RemoteRpcClient.TRC20Send(from, to, contract, amount, feeLimit)
}

func (e *TronChain) BroadcastTransaction(transaction *core.Transaction) error {
	err := e.keepConnect()
	if err != nil {
		return err
	}
	result, err := e.RemoteRpcClient.Broadcast(transaction)
	if err != nil {
		return fmt.Errorf("broadcast transaction error: %v", err)
	}
	if result.Code != 0 {
		return fmt.Errorf("bad transaction: %v", string(result.GetMessage()))
	}
	if result.Result == true {
		return nil
	}
	d, _ := json.Marshal(result)
	return fmt.Errorf("tx send fail: %s", string(d))
}

func (e *TronChain) GetTrxTransaction(txhash string) (*core.TransactionInfo, error) {
	err := e.keepConnect()
	if err != nil {
		return nil, err
	}
	return e.RemoteRpcClient.GetTransactionInfoByID(txhash)
}

func (e *TronChain) FreezeBalance(from, delegateTo string, ownerKey *ecdsa.PrivateKey, resource core.ResourceCode, frozenBalance int64) (*api.TransactionExtention, error) {
	var err error
	err = e.keepConnect()
	if err != nil {
		return nil, err
	}
	//获取冻结合约的合约实例
	contract := &core.FreezeBalanceContract{}
	if contract.OwnerAddress, err = common.DecodeCheck(from); err != nil {
		return nil, err
	}
	//contract.OwnerAddress =[]byte(from)

	contract.FrozenBalance = frozenBalance
	contract.FrozenDuration = 3 // Tron Only allows 3 days freeze

	if len(delegateTo) > 0 {
		if contract.ReceiverAddress, err = common.DecodeCheck(delegateTo); err != nil {
			return nil, err
		}
	}
	//contract.ReceiverAddress=[]byte(delegateTo)
	contract.Resource = resource

	tx, err := e.RemoteRpcClient.FreezeBalanceV2(from, core.ResourceCode_BANDWIDTH, frozenBalance)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("bad transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}

	signTx, err := SignTransaction(tx.GetTransaction(), ownerKey)
	if err != nil {
		fmt.Println(222)
		fmt.Println(err)
	}
	fmt.Println("Signed ready for broadcast txhash:::", common.BytesToHexString(signTx))

	result, err := e.RemoteRpcClient.Broadcast(tx.GetTransaction())

	if err != nil {
		fmt.Printf("freeze balance error: %v\n", err)
	}
	fmt.Println(result.String())

	return tx, nil

}

func (e *TronChain) UnFreezeBalance(from, delegateTo string, ownerKey *ecdsa.PrivateKey, resource core.ResourceCode) (*api.TransactionExtention, error) {
	var err error
	err = e.keepConnect()
	if err != nil {
		return nil, err
	}
	//获取冻结合约的合约实例
	contract := &core.UnfreezeAssetContract{}
	if contract.OwnerAddress, err = common.DecodeCheck(from); err != nil {
		return nil, err
	}

	tx, err := e.RemoteRpcClient.UnfreezeBalance(from, delegateTo, resource)
	if err != nil {
		fmt.Println("maybe Three days of unstaked")
		return nil, err
	}

	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("bad transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}

	signTx, err := SignTransaction(tx.GetTransaction(), ownerKey)
	if err != nil {
		fmt.Println(222)
		fmt.Println(err)
	}
	fmt.Println("Signed ready for broadcast txhash:::", common.BytesToHexString(signTx))

	result, err := e.RemoteRpcClient.Broadcast(tx.GetTransaction())

	if err != nil {
		fmt.Printf("unfreeze balance error: %v\n", err)
	}
	fmt.Println(result.String())
	return tx, nil
}

// SignTransaction 签名交易
func SignTransaction(transaction *core.Transaction, key *ecdsa.PrivateKey) ([]byte, error) {
	transaction.GetRawData().Timestamp = time.Now().UnixNano() / 1000000
	rawData, err := proto.Marshal(transaction.GetRawData())
	if err != nil {
		return nil, err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	contractList := transaction.GetRawData().GetContract()
	for range contractList {
		signature, err := crypto.Sign(hash, key)
		if err != nil {
			return nil, err
		}
		transaction.Signature = append(transaction.Signature, signature)
	}
	return hash, nil
}
