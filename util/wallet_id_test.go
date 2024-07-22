package util_test

import (
	"fmt"
	"github.com/Jmagicc/jmc-wallet-sdk/util"
	"testing"
)

func Test_SnowflakesID(t *testing.T) {

	// 创建一个雪花ID生成器
	snowflake := util.NewSnowflakeIDGenerator(1609459200000, 0, 12, 10)
	// 生成一个雪花ID
	id := snowflake.Generate()
	// 输出生成的雪花ID
	fmt.Println("Generated Snowflake ID:", id)
}
