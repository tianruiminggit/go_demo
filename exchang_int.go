package main

import "fmt"

func main() {
	var (
		x = 10
		y = 20
	)
	x, y = y, x
	fmt.Println("X=", x, "y=", y)
}
