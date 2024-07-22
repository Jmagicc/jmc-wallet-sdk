package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "Server desertion, to give it a try again later"
	message[REUQEST_PARAM_ERROR] = "Parameter error"
	message[TOKEN_EXPIRE_ERROR] = "Token fails, please login again"
	message[TOKEN_GENERATE_ERROR] = "Failure to generate the token"
	message[DB_ERROR] = "The database is busy, please try again later"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "Update the data affect the number of rows of 0"

	// 链模块
	message[CHAIN_RPC_ERROR] = "Chain rpc call failed"
	message[CHAIN_RPC_TIMEOUT_ERROR] = "Chain call timeout"
	message[CHAIN_DB_COIN_TABLEERROR] = "Currency table information query database failure"

	//转账模块
	message[CALCULATE_GAS_FEE_ERROR] = "Calculate gasFee failed"
	message[SEND_TRANSACTION_ERROR] = "Failed to initiate transaction"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "Server desertion, to give it a try again later"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
