package main

import "fmt"

type Age int

type MyInterface interface {
	Say() int
}

func main()  {
	c :=s1()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	age := Age(32)
	var age1 =&age
	//方法接受者 指针和值对象都可以使用
	age.String(2)
	age1.String(3)
	age.String1(4)
	age1.String1(5)

	//指定ageM是用指针接受的String方法变量
	ageM := (*Age).String
	ageM(age1,2)
	//类型断言
	var si MyInterface
	si = age
	s2,ok :=si.(Age)
	if ok {
		fmt.Println(s2)
	}else{
		fmt.Println("这个不是指针类型")
	}



	//ageM1 := Age.String
	//ageM1(age,1)
}

func s1() func() int {
	i :=0
	return func() int {
		i++
		return i
	}
}

func (age *Age) String(sex int) {
	fmt.Println("this is ",sex,"age ==",age)
}

func (age Age) String1(sex int){
	fmt.Println("this is ",sex,"age ==",age)
}

func (age Age) Say() int{
	return 1
}