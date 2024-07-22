package middleware

import (
	"encoding/json"
	"github.com/Jmagicc/jmc-wallet-sdk/crypto"
	"log"
	"net/http"
	"strconv"
)

type AESMiddleware struct {
	key string
	iv  string
}

func NewUserAgentMiddleware(key, iv string) *AESMiddleware {
	return &AESMiddleware{
		key: key,
		iv:  iv,
	}
}

// Handle 弃用
func (m *AESMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 调用 next.ServeHTTP 将请求传递给下一个处理程序
		rec := &responseCapture{ResponseWriter: w}
		next.ServeHTTP(rec, r)
		rec.encryptData([]byte(m.key), []byte(m.iv))
	}
}

type responseCapture struct {
	http.ResponseWriter
	body []byte
}

func (r *responseCapture) Write(b []byte) (int, error) {
	r.body = append(r.body, b...)
	return r.ResponseWriter.Write(b)
}

func (r *responseCapture) encryptData(key, iv []byte) {
	// 解析 JSON 响应
	var response map[string]interface{}
	err := json.Unmarshal(r.body, &response)
	if err != nil {
		log.Println("Error decoding JSON response:", err)
		return
	}

	data, exists := response["data"]
	if !exists {
		log.Println("No data field found in the response")
		return
	}

	var encryptedData string

	switch t := data.(type) {
	case string:
		encryptedData, err = crypto.EncryptByinv(key, iv, t)
		if err != nil {
			log.Println("Error encrypting data:", err)
			return
		}
	case []interface{}:
		dataSlice, ok := data.([]interface{})
		if !ok {
			log.Println("Invalid data type found in data field")
			return
		}

		var strData string
		for _, item := range dataSlice {
			str, ok := item.(string)
			if !ok {
				log.Println("Invalid element type found in data field")
				return
			}
			strData += str
		}

		encryptedData, err = crypto.EncryptByinv(key, iv, strData)
		if err != nil {
			log.Println("Error encrypting data:", err)
			return
		}
	default:
		log.Println("我找不到啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊")

	}

	response["data"] = encryptedData

	encryptedResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("Error encoding JSON response:", err)
		return
	}

	r.ResponseWriter.Header().Set("Content-Length", strconv.Itoa(len(encryptedResponse)))
	r.ResponseWriter.Header().Set("Content-Type", "application/json")

	_, err = r.ResponseWriter.Write(encryptedResponse)
	if err != nil {
		log.Println("Error writing encrypted response:", err)
		return
	}
}
