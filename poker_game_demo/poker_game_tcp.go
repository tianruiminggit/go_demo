package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type UserSocket struct {
	Name    string    `json:"name"`
	NetConn *net.Conn `json:"net_conn"`
}

type UserMsg struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

var UserSocketMap map[string]UserSocket
var RoomGoroutineMap map[string]int

func main() {
	listener, err := net.Listen("tcp", "192.168.4.65:8081")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	//defer listener.Close()
	initData()

	fmt.Println("Server started on port 8080")

	//初始数据

	for {
		// 接受新的连接请求
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}

		// 在新的goroutine中处理连接
		go receiveClientMsg(conn)
	}
}

func receiveClientMsg(conn net.Conn) {
	buffer := make([]byte, 512)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		head := buffer[:4]
		headNum := int32(head[0]<<24) + int32(head[1]<<16) + int32(head[2]<<8) + int32(head[3])
		message := buffer[4 : headNum+4]
		fmt.Println("Received message:", headNum, message, string(message))
		var userMsg UserMsg
		msgErr := json.Unmarshal(message, &userMsg)
		if msgErr != nil {
			fmt.Println("Error decoding from JSON:", msgErr)
			return
		}
		fmt.Println("Msg is ===", userMsg)
		//添加
		UserSocketMap[userMsg.Name] = UserSocket{userMsg.Name, &conn}
		// 发送响应
		_, err = conn.Write([]byte("Message received by server."))
	}
}

func initData() {
	//玩家socketMap
	UserSocketMap = make(map[string]UserSocket)
	//房间GoRoutineMap
	RoomGoroutineMap = make(map[string]int)
}
