package main

import "fmt"

func generate1() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter1(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate1()
		for {
			prime := <-ch
			ch = filter1(ch, prime)
			out <- prime
		}
	}()
	return out
}
func main() {
	primeCh := sieve()
	for {
		fmt.Println(<-primeCh)
	}
}
