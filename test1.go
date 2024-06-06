package main

import "fmt"

func main() {
    
    var str string ="tianruiming"
    fmt.Println(str[1])

    char := str[1]
    var i int 
    i = int(char+11)
    fmt.Println(char+1,i)

    fmt.Println("abc\nabc")

    fmt.Println(`abc\nabc`)

    fmt.Println(str+"\nnb")

    fmt.Printf("type is %T",char)
    
}