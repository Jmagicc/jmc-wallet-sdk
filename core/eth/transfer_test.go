package eth

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

// 主币转账
func Test_TransferCoin(t *testing.T) {

	//alcohol enlist margin alter general vote jelly tiny calm treat quote very
	//2023/08/01 16:58:00 助记词: alcohol enlist margin alter general vote jelly tiny calm treat quote very
	//2023/08/01 16:58:00 公钥: 0xACCcdf07a7c6Ed7824e5643950Ab1a085F7A8743
	//2023/08/01 16:58:00 私钥: d78c62549254e9464f7b48ba0331bf6c06effbe70b0d33f4be19e960a1a8f4c2

	//ecology account lawn hope robot list social couple good stool burger tribe
	//2023/08/01 16:57:13 助记词: ecology account lawn hope robot list social couple good stool burger tribe
	//2023/08/01 16:57:13 公钥: 0x6BaC344C3e91Da32DfF048ce01d385e93fC7FF82
	//2023/08/01 16:57:13 私钥: 96288ce0d39ce5eb1ea4052b9697b902714aa791698a1b7fd682a5c61062d9ee

	var (
		address    = "0xACCcdf07a7c6Ed7824e5643950Ab1a085F7A8743"
		toAddress  = "0x6BaC344C3e91Da32DfF048ce01d385e93fC7FF82"
		PrivateKey = "d78c62549254e9464f7b48ba0331bf6c06effbe70b0d33f4be19e960a1a8f4c2"
		Amount     = "1000000000000000000" //转账金额    0.01以太坊
	)
	// 通过链ID获取链的配置信息
	config := GetRpcInfoByChainId(11155111)
	ethChain := NewEthChain()
	_, err := ethChain.CreateRemote(config.Url)
	if err != nil {
		log.Fatalln("创建链客户端失败", err)
	}

	nonce, _ := ethChain.Nonce(address)
	nonceInt, _ := strconv.ParseInt(nonce, 10, 64)
	// 构造多笔交易则nonce + 1
	callMethodOpts := &CallMethodOpts{
		Nonce: nonceInt,
		Value: Amount,
	}
	TxResult, errTxResult := ethChain.BuildTransferTx(PrivateKey, toAddress, callMethodOpts)
	if errTxResult != nil {
		log.Fatalln("构造交易失败")
	}

	sendTxHash, err := ethChain.SendRawTransaction(TxResult.TxHex)
	if err != nil {
		log.Fatalln("发送交易失败")
	}
	log.Println("交易hash:", sendTxHash)

}

// 代币转账
func Test_TransferToken(t *testing.T) {

	var (
		address       = "0xACCcdf07a7c6Ed7824e5643950Ab1a085F7A8743"
		toAddress     = "0x6BaC344C3e91Da32DfF048ce01d385e93fC7FF82"
		PrivateKey    = "d78c62549254e9464f7b48ba0331bf6c06effbe70b0d33f4be19e960a1a8f4c2"
		ConstractAddr = "0xa9663dF61f3e14611c3Bf2cbfDEc91EF3a549Db6"
		Amount        = "1000000" //转账金额  1U

	)
	// 通过链ID获取链的配置信息
	config := GetRpcInfoByChainId(11155111)
	ethChain := NewEthChain()
	_, err := ethChain.CreateRemote(config.Url)
	if err != nil {
		log.Fatalln("创建链客户端失败", err)
	}

	// 代币转账
	nonce, _ := ethChain.Nonce(address)
	nonceInt, _ := strconv.ParseInt(nonce, 10, 64)
	// 构造多笔交易则nonce + 1
	callMethodOpts := &CallMethodOpts{
		Nonce:    nonceInt,
		GasLimit: "81000",
	}
	abi, ok := FindMainnetABI(ConstractAddr)
	if !ok {
		log.Println("没有找到对应的ABI...")
	}

	erc20JsonParams := `{"toAddress":"` + toAddress + `",` + `"amount":"` + Amount + `", "method":"transfer"}`
	// erc20 代币转账
	buildTxResult, buildCallMethodTxerr := ethChain.BuildCallMethodTx(
		PrivateKey,
		ConstractAddr,
		abi.Abi,
		// 调用的合约方法名
		"transfer",
		callMethodOpts,
		// 转账目标地址
		erc20JsonParams)

	if buildCallMethodTxerr != nil {
		fmt.Printf("build erc20 call method tx error: %v\n", buildCallMethodTxerr)
		return
	}

	// 发送交易
	sendTxHash, err := ethChain.SendRawTransaction(buildTxResult.TxHex)
	if err != nil {
		log.Printf("send raw transaction error: %v\n\n", err)
	}
	fmt.Println("交易hash:", sendTxHash)

}

// 查看交易信息
func Test_TranscationInfo(t *testing.T) {
	var (
		txHash = "0x4fb4c8740aba74210ec181b9dedcb001ba4ff776602cd5f70c251667182e9e04"
	)
	config := GetRpcInfoByChainId(11155111)
	ethChain := NewEthChain()
	chainClient, err := ethChain.CreateRemote(config.Url)
	if err != nil {
		log.Fatalln("创建链客户端失败", err)
	}

	detail, err := chainClient.TransactionReceiptByHash(txHash)
	log.Println("交易详情:", detail.Status)

}

// 查看链是否能用
func Test_Chain(t *testing.T) {
	config := GetRpcInfoByChainId(56)
	ethChain := NewEthChain()
	chainClient, err := ethChain.CreateRemote(config.Url)
	if err != nil {
		log.Fatalln("创建链客户端失败", err)
	}

	detail, err := chainClient.LatestBlockNumber()
	log.Println(detail, err)

}
