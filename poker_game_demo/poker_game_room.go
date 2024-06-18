package main

import (
	"fmt"
	"strconv"
)

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
	Name string `json:"name"`
	Pos  int    `json:"pos"` //位置

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
	fmt.Println(playerMap)

	select {
	case operateMsg := <-*chanPlayerOperate:
		doPlayerOperate(operateMsg, &playerMap)
		break
	case <-*chanRoomProgress:
		break
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
			break
		}
		//获取房间内所有玩家,推送新玩家加入
		var keys []string
		for key := range *playerMap {
			keys = append(keys, key)
		}
		sendClientMsg(keys, map[string]string{"name": "trm"})
		//加入新玩家
		(*playerMap)[name] = roomPlayer{name, 1}
		break
	case "2":
		fmt.Println("doPlayerOperate =2=",*playerMap)
	}
}
