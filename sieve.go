package main

import "fmt"

func generate(ch chan int) {
	for i := 2; i < 999; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
func main() {
	ch := make(chan int)
	go generate(ch)
	for i := 0; i < 3; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
