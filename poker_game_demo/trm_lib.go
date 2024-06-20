package main

import (
	"fmt"
	"net/url"
	"reflect"
)

type ClientResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func cmd2Function(msg UserMsg) ClientResponse {
	//取出cmd  根据CMD找到对应的函数名
	//这里测试 cmd就是方法名
	funcName := msg.CMD
	// 获取函数的reflect.Value
	funcValue := reflect.ValueOf(&msg).MethodByName(funcName)
	// 检查函数是否存在
	if funcValue.IsValid() && funcValue.Kind() == reflect.Func {
		//准备函数调用的参数
		argsMap, _ := url.ParseQuery(msg.Msg)
		args := []reflect.Value{reflect.ValueOf(argsMap)}
		// 调用函数
		res := funcValue.Call(args)[0]
		//fmt.Println(res.Kind())
		//elems := res.Elem()
		return ClientResponse{res.Field(0).Interface().(int), res.Field(1).Interface().(string)}

	} else {
		fmt.Println("Function not found", funcName, funcValue, funcValue.IsValid(), funcValue.Kind())
		return ClientResponse{-1, "Function not found"}
	}
}

func (msg *UserMsg) Raise(args map[string][]string) ClientResponse {
	fmt.Println(msg)
	return ClientResponse{1, "ok"}
}

func (msg *UserMsg) EnterRoom(args map[string][]string) ClientResponse {
	fmt.Println(msg)
	return ClientResponse{1, "ok"}
}

func (msg *UserMsg) GetUserSocket(args map[string][]string) ClientResponse {
	fmt.Println(UserSocketMap)
	return ClientResponse{1, "ok"}
}

func (msg *UserMsg) PlayerRoomOperate(args map[string][]string) ClientResponse {
	roomId := args["roomId"][0]
	operate := args["operate"][0]
	room, exist := RoomMap[roomId]
	fmt.Println("PlayerRoomOperate Args =", args, RoomMap, exist)
	if exist {
		*(room.chanPlayerOperate) <- RoomMsg{msg.Name, roomId, operate}
		return ClientResponse{1, "PlayerRoomOperate"}
	} else {
		return ClientResponse{-1, "do not find room"}
	}
}
