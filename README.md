[//]: # (eth , tron , aptos)
[//]: # (btc , doge , cosmos , solana , sui , polka)

+ 创建服务命令

```shell
#首次生成代码(研发环境)1
goctl api new 服务名称
goctl rpc new 服务名称

#生成api服务代码(研发环境)
goctl api go -api wallet.api -dir .  -style go_zero
go run wallet.go -f etc/wallet.yaml

#4.生成RPC服务代码(研发环境)
goctl rpc protoc wallet.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. --style go_zero
go run wallet.go -f etc/wallet.yaml

#生成数据库curd代码  cd ~/workspace/model(研发环境)
goctl model mysql ddl --src bc_wallet_order.sql --dir .  -cache=true
goctl model mysql ddl --src bc_chain.sql  --dir .  -cache=true


#测试环境启动命令  Linux下   wallet/cmd    两个服务,一个api,一个rpc   
go run wallet.go -f etc/wallet-test.yaml
```


