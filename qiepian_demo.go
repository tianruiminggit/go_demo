package main

import "fmt"

func main()  {

	myAyyay := []int{1,2,3,4,5,6,7,8,9}
	mySlice := myAyyay[2:4:5]
	mySlice=append(mySlice,20)
	mySlice=append(mySlice,21)
	fmt.Println(mySlice,cap(mySlice),mySlice[1])
	fmt.Println(myAyyay)
	s:=make([]byte,5)
	fmt.Println(cap(s),len(s))
}
