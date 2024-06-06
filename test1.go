package main

import "fmt"

type Empty interface{}

type semaphore chan Empty

func main() {
	var (
		i=1
		j=2
	)
	var str = "hello"
	fmt.Println(i,j,str)
}

func (S semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		S <- e
	}
}

func (S semaphore) V(n int) {
	<-S
}

func (S semaphore)Lock(){
	S.P(1)
}

func (S semaphore) unlock(){
	S.V(1)
}


