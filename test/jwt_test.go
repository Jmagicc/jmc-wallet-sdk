package test

import (
	"fmt"
	"github.com/Jmagicc/jmc-wallet-sdk/util"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"runtime"
	"testing"
	"time"
)

// JWT 生成, 用于测试
func Test_JWT(t *testing.T) {
	secretKey := "041bedfe-e13f-464e-9a67-c43dedf7113c"
	uid := 27
	// 创建一个新的 Token
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置 Token 的声明（Claim）
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = uid
	// 设置令牌过期时间1年
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	//claims["exp"] = time.Now().Add(time.Hour).Unix() // 设置令牌过期时间

	// 使用密钥进行签名
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("生成JWT token失败:", err)
		return
	}
	fmt.Println("生成的JWT token:", tokenString)
}

func Test_ID(t *testing.T) {
	snowflake := util.NewSnowflakeIDGenerator(1609459200000, 0, 12, 10)
	walletId := snowflake.Generate()
	fmt.Println(walletId)
}

func Test_DIR(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
	fmt.Println(runtime.GOOS)

}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE5NjA4NjIsInVpZCI6MjB9.HymVjg1CNaH_fd-uPcrgyAq69jYjoNfOWdqNS7CQWec  -20
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE5NjY0MTUsInVpZCI6MjF9.GgC-UADqj9m_VHceD4pDMPq-OT-LCDRVec7CNyyb_lg  -21
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE5NjY2OTEsInVpZCI6MjJ9.QL4dfeepqd7VS3WPR6PzWR36OmLGnTO4r94EJ6IViZg  -22

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE5ODEzMjYsInVpZCI6MjR9.NiNb_J3WTyasJuUmmEsjAQW5srZzQojKyJ6D_T8SWmE -24
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjIxMzQxMDMsInVpZCI6MjV9.sYJYHsLz1Egi4vjuZhqRDxNRD5Loi3woz2pJfrPdolU -25
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjI0NzYyMzgsInVpZCI6MjZ9.0P8cL3p39_5pvnoezIHE_7_tnyEhQNZt52uPoi9yf1I -26

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjI1MDM2MDcsInVpZCI6Mjd9.OeuIj2U_Xg-o2asY8LOK1ufMahcMMf9nCLHE0SjUyQg -27
