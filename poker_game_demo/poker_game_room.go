package main

import (
	"fmt"
	"strconv"
)

/*
玩家状态常量
*/
const (
	PlayerStateFree    = 0 //空闲
	PlayerStatePrepare = 1 //准备
	PlayerStateWatch   = 2 //观战
)

//自定义数据类型

type SendMsg map[string]string

//自定义struct

type GameRoom struct {
	id                int
	chanPlayerOperate *chan RoomMsg
	chanRoomProgress  *chan int
}

type RoomMsg struct {
	name    string
	roomId  string
	operate string
}

type roomPlayer struct {
	Name  string `json:"name"`
	Pos   int    `json:"pos"` //位置
	State int    `json:"state"`
}

var RoomMap map[string]GameRoom

func initRoom() {
	RoomMap = make(map[string]GameRoom)
	for i := 0; i < 5; i++ {

		chanPlayerOperate := make(chan RoomMsg)
		chanRoomProgress := make(chan int)
		RoomMap[strconv.Itoa(i)] = GameRoom{i, &chanPlayerOperate, &chanRoomProgress}
		go goRoomControl(&chanPlayerOperate, &chanRoomProgress)
	}
}

// 房间控制
func goRoomControl(chanPlayerOperate *chan RoomMsg, chanRoomProgress *chan int) {
	//pos当为key
	playerMap := make(map[string]roomPlayer)
	for {
		select {
		case operateMsg := <-*chanPlayerOperate:
			doPlayerOperate(operateMsg, &playerMap)
		case <-*chanRoomProgress:
		}
	}

}
func doPlayerOperate(msg RoomMsg, playerMap *map[string]roomPlayer) {
	fmt.Println("doPlayerOperate Map =", *playerMap, msg)
	operate := msg.operate
	switch operate {
	case "1": //加入房间
		name := msg.name
		_, exist := (*playerMap)[name]
		if exist {
			fmt.Println("player exist")
			break
		}
		temp := roomPlayer{name, 1, PlayerStateFree}
		sendAllRoomMsg(playerMap, formatPlayerSendMsg(temp))
		//加入新玩家
		(*playerMap)[name] = temp
		break
	case "2":
		fmt.Println("doPlayerOperate =2=", *playerMap)
	case "3": //准备/取消准备
		name := msg.name
		player, exist := (*playerMap)[name]
		if exist {
			if player.State == PlayerStateFree {
				player.State = PlayerStatePrepare
			}

			if player.State == PlayerStatePrepare {
				player.State = PlayerStateFree
			}
			sendRoomMsg(playerMap, formatPlayerSendMsg(player), player.Name)
		}
	}
}

// 给房间所有人发送消息
func sendAllRoomMsg(playerMap *map[string]roomPlayer, msg SendMsg) {
	//获取房间内所有玩家,推送新玩家加入
	var keys []string
	for key := range *playerMap {
		keys = append(keys, key)
	}
	sendClientMsg(keys, msg)
}

// 给房间所有人发送消息
func sendRoomMsg(playerMap *map[string]roomPlayer, msg SendMsg, noSendKey string) {
	//获取房间内所有玩家,推送新玩家加入
	var keys []string
	for key := range *playerMap {
		if key != noSendKey {
			keys = append(keys, key)
		}

	}
	sendClientMsg(keys, msg)
}

func formatPlayerSendMsg(player roomPlayer) SendMsg {
	return SendMsg{"name": player.Name, "pos": strconv.Itoa(player.Pos), "state": strconv.Itoa(player.State)}
}
