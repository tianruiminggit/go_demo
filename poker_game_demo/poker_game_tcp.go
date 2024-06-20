package main

import (
	"encoding/json"
	"log"
	"net"
	"net/url"
)

type UserSocket struct {
	Name              string    `json:"name"`
	NetConn           *net.Conn `json:"net_conn"`
	chanHandleRequest *chan UserMsg
}

type UserMsg struct {
	Name string `json:"name"`
	CMD  string `json:"cmd"`
	Msg  string `json:"msg"`
}

type BroadcastMsg struct {
	CMD string `json:"cmd"`
	Msg string `json:"msg"`
}

var UserSocketMap map[string]UserSocket

func main() {
	listener, err := net.Listen("tcp", "192.168.4.65:8081")
	if err != nil {
		log.Println("Error listening:", err.Error())
		return
	}
	//defer listener.Close()
	initData()

	log.Println("Server started on port 8080")

	//初始数据

	for {
		// 接受新的连接请求
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			continue
		}

		// 在新的goroutine中处理连接
		go goReceiveClientMsg(conn)
	}
}

// Goroutine ：：玩家的socket监听  接受原始请求
func goReceiveClientMsg(conn net.Conn) {
	buffer := make([]byte, 512)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading:", err.Error())
			return
		}
		head := buffer[:4]
		headNum := int32(head[0]<<24) + int32(head[1]<<16) + int32(head[2]<<8) + int32(head[3])
		message := buffer[4 : headNum+4]
		log.Println("Received message:", headNum, message, string(message))
		var userMsg UserMsg
		msgErr := json.Unmarshal(message, &userMsg)
		if msgErr != nil {
			log.Println("Error decoding from JSON:", msgErr)
			return
		}
		argsMap, _ := url.ParseQuery(userMsg.Msg)
		userSocket, exist := UserSocketMap[userMsg.Name]
		log.Println("Msg is ===", userMsg, argsMap, exist, UserSocketMap)
		if exist {
			*(userSocket.chanHandleRequest) <- userMsg
		} else {
			//添加
			//goHandleClientQuest 通道 500个请求缓冲
			chanGoHandleClientQuest := make(chan UserMsg, 500)
			UserSocketMap[userMsg.Name] = UserSocket{userMsg.Name, &conn, &chanGoHandleClientQuest}
			go goHandleClientQuest(&conn, &chanGoHandleClientQuest)
			chanGoHandleClientQuest <- userMsg
		}
		//go goHandleClientQuest(&conn,&chanGoHandleClientQuest)

	}
}

func goSendMsgToClient(conn *net.Conn, chanMsg *chan BroadcastMsg) {
	for {
		msg := <-*chanMsg
		_, err := (*conn).Write([]byte(msg.Msg))
		if err != nil {
			log.Println(" goSendMsgTOClient Error decoding from JSON:", err)
			return
		}
	}
}

func goHandleClientQuest(conn *net.Conn, chanMsg *chan UserMsg) {
	for msg := range *chanMsg {
		response := cmd2Function(msg)
		log.Println("goHandleClientQuest call return", response)
		// 发送响应
		_, err := (*conn).Write([]byte(response.Msg))
		if err != nil {
			log.Println("Error decoding from JSON:", err)
			return
		}
	}
}

func initData() {
	//玩家socketMap
	UserSocketMap = make(map[string]UserSocket)
	initRoom()
}

func sendClientMsg(names []string, msg map[string]string) {
	log.Println("sendClientMsg", names)
	for _, name := range names {
		socket, exist := UserSocketMap[name]
		if exist {
			log.Println("sendClientMsg name = ", name)
			msg["receive"] = name
			// 将map转换为JSON字节切片
			jsonBytes, err := json.Marshal(msg)
			if err != nil {
				log.Println("Error marshaling map:", err)
				return
			}

			_, err = (*(socket.NetConn)).Write(jsonBytes)
			if err != nil {
				log.Println("Error decoding from JSON:", err)
				return
			}
		}
	}
}
