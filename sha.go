package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type MyData1 struct {
	Kv_list []MyData2 `json:"kv_list"`
}
type MyData2 struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	jsonStr := []byte(`{"kv_list":[{"key":"score","value":"100"},{"key":"gold","value":"3000"}]}`)
	fmt.Println(jsonStr)
	// 解析JSON字符串为Go结构体
	//var jsonData1 MyData1
	//var jsonData2 MyData2
	//err := json.Unmarshal([]byte(jsonStr), &jsonData1)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// 将结构体转换为JSON字节
	//jsonBytes, err := json.Marshal(jsonData1)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// 定义密钥
	key := []byte("9hAb/NEYUlkaMBEsmFgzig==")
	fmt.Println(key)
	// 创建HMAC-SHA256哈希计算器
	h := hmac.New(sha256.New, key)

	// 将JSON字节写入哈希计算器
	h.Write(jsonStr)

	// 获取HMAC-SHA256哈希结果
	hash := h.Sum(nil)

	// 将哈希结果转换为十六进制字符串便于查看
	hashHex := hex.EncodeToString(hash)
	fmt.Println(hashHex)
}
