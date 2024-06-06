package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	chSign := make(chan int)
	go sendData(ch, chSign)
	go getData(ch, 1)
	go getData(ch, 2)
	//getData(ch, 2)
	//select {}
	//
	//time.Sleep(1e9)
	<-chSign
	//go sum

	//chSum := make(chan int)
	//go goSum(4, 5, chSum)
	//x := <-chSum
	//fmt.Println(x)
}

func sendData(ch chan string, chSign chan int) {
	ch <- "t"
	ch <- "i"
	ch <- "a"
	ch <- "n"
	ch <- "r"
	ch <- "u"
	ch <- "i"
	chSign <- 1

}

func getData(ch chan string, id int) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%d%s ", id, input)
	}
}

//func goSum(x, y int, ch chan int) {
//	ch <- x + y
//}
