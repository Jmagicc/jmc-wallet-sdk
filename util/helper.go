package util

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func Int64ToBytes(n int64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	size := binary.PutVarint(buf, n)
	return buf[:size]
}

func IsMainCoin(currency string) bool {
	switch currency {
	case "ETH", "BNB", "MATIC":
		return true
	default:
		return false
	}
}

// GweiToEther GweiToEther 将 gwei 转换为 ether
func GweiToEther(gwei string) (string, error) {
	gweiDec, err := decimal.NewFromString(gwei)
	if err != nil {
		return "", err
	}

	etherDec := gweiDec.Div(decimal.NewFromFloat(1000000000))

	return etherDec.String(), nil
}

// WeiToEther 将 Wei 转换为 Ether
func WeiToEther(wei string) (string, error) {
	weiDec, err := decimal.NewFromString(wei)
	if err != nil {
		return "", err
	}

	etherDec := weiDec.Div(decimal.NewFromFloat(1000000000000000000.0))

	return etherDec.String(), nil
}

// ShiftBalance 按精度转换
func ShiftBalance(balance, precision string) (string, error) {
	if balance == "" || precision == "" {
		return "", nil
	}

	var precisionByint64 int64
	switch precision {
	case "6":
		precisionByint64 = int64(1000000)
	case "18":
		precisionByint64 = int64(1000000000000000000)
	case "1":
		precisionByint64 = int64(1)
	default:
		precisionByint64 = int64(1000000000000000000)
	}

	value, success := new(big.Float).SetString(balance)
	if !success {
		return "", fmt.Errorf("invalid input balance: %s", balance)
	}

	precisionBig := new(big.Float).SetInt(big.NewInt(precisionByint64))

	// 检查除数是否为0，避免除以0的错误
	if precisionBig.Cmp(big.NewFloat(0)) == 0 {
		log.Panicln(precision, "为什么这精度是0 啊", precisionBig)
		return "", errors.New("division by zero")
	}

	result := new(big.Float).Quo(value, precisionBig)

	// 将结果转换为固定小数位数的十进制表示形式
	resultStr := result.Text('f', -1)

	return resultStr, nil
}

func AddFloatStrings(a, b string) (string, error) {
	decimalA, err := decimal.NewFromString(a)
	if err != nil {
		return "", fmt.Errorf("invalid input: %v", err)
	}

	decimalB, err := decimal.NewFromString(b)
	if err != nil {
		return "", fmt.Errorf("invalid input: %v", err)
	}

	result := decimalA.Add(decimalB)

	return result.String(), nil
}

// MultiplyFloatStrings 两个字符串类型的浮点数相乘
func MultiplyFloatStrings(a, b string) (string, error) {
	decimalA, err := decimal.NewFromString(a)
	if err != nil {
		return "", fmt.Errorf("invalid input: %v", err)
	}

	decimalB, err := decimal.NewFromString(b)
	if err != nil {
		return "", fmt.Errorf("invalid input: %v", err)
	}

	result := decimalA.Mul(decimalB)

	return result.String(), nil
}

func Md5GenerateSaltedHash(password string, salt int64) string {
	// 将密码和盐进行拼接
	str := strconv.FormatInt(salt, 10)
	saltedPassword := password + str

	// 创建一个MD5哈希对象
	hash := md5.New()

	// 将拼接后的字符串转换为字节数组并计算哈希值
	io.WriteString(hash, saltedPassword)
	hashedPassword := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串表示
	hashedPasswordHex := hex.EncodeToString(hashedPassword)

	return hashedPasswordHex
}

// ConvertTimestampToFormat 时间戳转换成日期格式
func ConvertTimestampToFormat(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02")
}

// CalculateGasFee 计算最终的GasFee
func CalculateGasFee(gasLimitStr, gasPriceStr string) (gasFeeStr string, err error) {
	gasLimit, err := strconv.ParseUint(gasLimitStr, 10, 64)
	if err != nil {
		return "", err
	}

	gasPrice, ok := new(big.Int).SetString(gasPriceStr, 10)
	if !ok {
		return "", fmt.Errorf("invalid gas price")
	}

	gasFee := new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(gasLimit))
	gasFeeStr = gasFee.String()

	return gasFeeStr, nil
}

func If(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

// 将一个string类型的数组转换成string类型用逗号分割的字符串
func StringArrayToString(arr []string) string {
	return strings.Join(arr, ",")
}

// RFC3339ToNormalTime RFC3339 日期格式标准化
func RFC3339ToNormalTime(rfc3339 string) string {
	if len(rfc3339) < 19 || rfc3339 == "" || !strings.Contains(rfc3339, "T") {
		return rfc3339
	}
	return strings.Split(rfc3339, "T")[0] + " " + strings.Split(rfc3339, "T")[1][:8]
}

/*func GenerateToken(id uint, identity, name string, second int) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(second))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}
*/
// httpRequest .
func httpRequest(url, method string, data, header []byte) ([]byte, error) {
	var err error
	reader := bytes.NewBuffer(data)
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// 处理 header
	if len(header) > 0 {
		headerMap := new(map[string]interface{})
		err = json.Unmarshal(header, headerMap)
		if err != nil {
			return nil, err
		}
		for k, v := range *headerMap {
			if k == "" || v == "" {
				continue
			}
			request.Header.Set(k, v.(string))
		}
	}
	//request.SetBasicAuth(define.EmqxKey, define.EmqxSecret)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}

func HttpDelete(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "DELETE", data, header)
}

func HttpPut(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "PUT", data, header)
}

func HttpPost(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "POST", data, header)
}

func HttpGet(url string, header ...byte) ([]byte, error) {
	return httpRequest(url, "GET", []byte{}, header)
}
