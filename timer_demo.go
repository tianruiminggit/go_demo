package main

import (
	"fmt"
	"time"
)

func main(){
	timerExample()
	tickerExample()
	//fmt.Println()
}

func timerExample() {
	// 创建一个延迟5秒的一次性定时器
	timer := time.NewTimer(5 * time.Second)

	// 阻塞等待定时器触发
	a:=<-timer.C

	fmt.Println("Timer triggered after 5 seconds",a)

	// 停止并检查定时器状态
	if stopped := timer.Stop(); !stopped {
		fmt.Println("Timer already stopped or expired")
	} else {
		fmt.Println("Timer successfully stopped")
	}
}

func tickerExample() {
	// 创建一个每隔1秒触发一次的周期性定时器
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for i :=range ticker.C {
			fmt.Println("Ticker ticked",i)
		}
	}()
	// 等待3秒后停止ticker
	time.Sleep(3 * time.Second)
	ticker.Stop()
	fmt.Println(time.Now())
	fmt.Println("Ticker stopped")
}
