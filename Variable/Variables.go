package main

import "fmt"

func main() {
	//变量
	//var 变量名字 类型 = 表达式；变量名:=表达式

	//var x int32
	//x = 10
	x := 10
	println("x=", x)
	//fmt.Println("x=",x)

	var s string
	s = "string01"
	println("s=", s)

	s1 := "string02"
	println("s1=", s1)

	var a = "hello"
	//fmt.Println("a="+a)
	fmt.Println("a=", a)

	var b int
	fmt.Println("b=", b)

	var c bool
	fmt.Println("c=", c)
}
