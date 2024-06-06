package main

import "fmt"

func main(){
	channel1 := make(chan int)
	//close(channel1)
	go func(){
		channel1<- 2
	}()
	i,ok :=<-channel1
	if ok {
		fmt.Println(i)
	}else {
		fmt.Println(ok)
	}

}