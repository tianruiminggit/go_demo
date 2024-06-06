package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MyJson struct {
	Version int32  `json:"version"`
	Openid  string `json:"openid"`
	Scene   int32  `json:"scene"`
	Content string `json:"content"`
}

type MyJson1 struct {
	Uid  int64  `json:"roleid"`
	Str1 string `json:"weekgifts"`
	Str2 string `json:"monthgifts"`
}

func main() {
	//str :="/sgame_gm/gm_wx_sign_in_port/find_user&"
	data := MyJson1{2814792716779531, "prop,2010208,1", "prop,2010208,1"}
	//data :="[{roleid:281479271688953},{weekgifts:prop,200015,1}]"
	//jsonData, err :=json.Marshal(`[{version:2},{openid:"oc72y65HM2E6_tXpIlgzxttlYazg"},{scene:1},{content:"习金瓶"}]`)
	jsonData, err := json.Marshal(data)
	fmt.Println(jsonData)
	//req, err := http.NewRequest(http.MethodPost, "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=79_nrWD4oycQEgNlE1_U6cOA0HBxBUAjxahTD6kfdvxYCmEO87_vO-jxTAwCmmaFsrBVLqOjIVH4u2JpgsxiVm4lemgBAyHDEldZvAwkXuweVquHSwMA2vOU73t4TcYNMjABAHWV",
	//	bytes.NewBuffer(jsonData))
	//finByte := append(jsonData,[]byte(str)...)
	// 定义密钥
	key := []byte("dmgzh.qixia.com")
	Val := []byte("roleid=281479271677953&weekgifts=rmb,50&monthgifts=")
	// 使用SHA-256创建一个新的HMAC实例
	h := hmac.New(sha256.New, key)

	// 写入数据到HMAC实例中
	h.Write(Val)

	// 计算HMAC并获取其字节切片
	mac := h.Sum(nil)

	// 将字节切片转换为十六进制字符串
	macHex := hex.EncodeToString(mac)

	fmt.Printf("HMAC-SHA-256: %s\n", macHex)
	req, err := http.NewRequest(http.MethodPost, "http://192.168.4.65:882/sgame_gm/gm_wx_sign_in_port/send_mail_award?sign="+macHex, bytes.NewBuffer(Val))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(respBody))
}
