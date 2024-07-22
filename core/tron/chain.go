package tron

import (
	"errors"
	"github.com/Jmagicc/jmc-wallet-sdk/core/base"
)

type IChain interface {
	base.Chain
	SubmitTransactionData(account base.Account, to string, data []byte, value string) (string, error)
	GetEthChain() (*TronChain, error)
	//EstimateGasLimit(msg *CallMsg) (gas *base.OptionalString, err error)
}

type Chain struct {
	RpcUrl string
}

func (c *Chain) MainToken() base.Token {
	//return &Token{chain: c}
	return nil
}

func (c *Chain) BalanceOfAddress(address string) (*base.Balance, error) {
	b := base.EmptyBalance()

	if !IsValidAddress(address) {
		return b, errors.New("Invalid hex address")
	}

	chain := NewTronChain()
	balance, err := chain.Balance(address)
	if err != nil {
		return b, err
	}
	return &base.Balance{
		Total:  balance,
		Usable: balance,
	}, nil
}

// Deprecated: This method is no longer supported. Please use NewMethod() instead
func (c *Chain) BalanceOfPublicKey(publicKey string) (*base.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Chain) BalanceOfAccount(account base.Account) (*base.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Chain) SendRawTransaction(signedTx string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Chain) SendSignedTransaction(signedTxn base.SignedTransaction) (*base.OptionalString, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Chain) FetchTransactionDetail(hash string) (*base.TransactionDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Chain) FetchTransactionStatus(hash string) base.TransactionStatus {
	//TODO implement me
	panic("implement me")
}

func (c *Chain) BatchFetchTransactionStatus(hashListString string) string {
	//TODO implement me
	panic("implement me")
}

func (c *Chain) EstimateTransactionFee(transaction base.Transaction) (fee *base.OptionalString, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *Chain) EstimateTransactionFeeUsePublicKey(transaction base.Transaction, pubkey string) (fee *base.OptionalString, err error) {
	//TODO implement me
	panic("implement me")
}

func NewChainWithRpc(rpcUrl string) *Chain {
	return &Chain{
		RpcUrl: rpcUrl,
	}
}
