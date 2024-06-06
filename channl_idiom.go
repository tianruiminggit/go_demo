package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	pump(ch)
	suck(ch)
	time.Sleep(1e9)
}

func pump(ch chan int) {
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
}

func suck(ch chan int) {
	go func() {
		for V := range ch {
			fmt.Println(V)
		}
	}()
}
