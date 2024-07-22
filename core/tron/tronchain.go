package tron

import (
	"fmt"
	"github.com/Jmagicc/jmc-wallet-sdk/core/base"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"google.golang.org/grpc"
	"math/big"
	"time"
)

type TronChain struct {
	RemoteRpcClient *client.GrpcClient
	RpcClient       *client.GrpcClient
	timeout         time.Duration
	chainId         *big.Int
	rpcUrl          string
}

func NewTronChain() *TronChain {
	timeout := 60 * time.Second
	return &TronChain{
		timeout: timeout,
	}
}

func (e *TronChain) CreateRemote(rpcUrl string) (chain *TronChain, err error) {
	return e.CreateRemoteWithTimeout(rpcUrl, 0)
}

// @param timeout time unit millsecond. 0 means use chain's default: 60000ms.
func (e *TronChain) CreateRemoteWithTimeout(rpcUrl string, timeout int64) (chain *TronChain, err error) {
	defer base.CatchPanicAndMapToBasicError(&err)

	remoteRpcClient := client.NewGrpcClient(rpcUrl)
	err = e.RemoteRpcClient.Start(grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("grpc client start error: %v", err)
	}

	e.chainId = big.NewInt(1)
	e.RpcClient = remoteRpcClient
	e.RemoteRpcClient = remoteRpcClient
	e.rpcUrl = rpcUrl
	return e, nil
}

func (e *TronChain) ConnectRemote(rpcUrl string) error {
	_, err := e.CreateRemote(rpcUrl)
	return err
}

func (e *TronChain) Close() {
	if e.RemoteRpcClient != nil {
		e.RemoteRpcClient.Stop()
	}
	if e.RpcClient != nil {
		e.RpcClient.Stop()
	}
}
