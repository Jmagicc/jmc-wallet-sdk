package xerr

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const DB_ERROR uint32 = 100005
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006

// 钱包创建模块
const WALLET_CREATE_ERROR uint32 = 200001                // 钱包创建失败
const WALLET_CREATE_ADDRESS_ERROR uint32 = 200002        // 钱包地址生成失败
const WALLET_CREATE_ADDRESS_EXIST_ERROR uint32 = 200003  // 钱包地址未存在
const WALLET_CREATE_ADDRESS_INSERT_ERROR uint32 = 200004 // 钱包地址验证失败
const WALLET_CREATE_TRADEPWD_ERROR uint32 = 200005       // 交易密码验证失败
const WALLET_CREATE_TRADEPWD_SET_ERROR uint32 = 200006   // 设置交易密码失败
const WALLET_RESET_TRADEPWD_ERROR uint32 = 200007        //重置密码失败
const WALLET_VERIFY_MNEMONIC_ERROR uint32 = 200008       // 验证助记词失败

// 链模块
const CHAIN_RPC_ERROR uint32 = 300001          // 链rpc调用失败
const CHAIN_RPC_TIMEOUT_ERROR uint32 = 300002  //链调用超时
const CHAIN_DB_COIN_TABLEERROR uint32 = 300003 //查询数据库币表信息失败
const CHAIN_GET_ABI_ERROR uint32 = 300004      //获取ABI失败

// 转账模块
const CALCULATE_GAS_FEE_ERROR uint32 = 400001      // 计算gasFee失败
const SEND_TRANSACTION_ERROR uint32 = 400002       //发起交易失败
const BUILD_TRANSACTION_ERROR uint32 = 400003      //构造交易失败
const CHAIN_SEND_TRANSACTION_ERROR uint32 = 400004 //发送交易失败
