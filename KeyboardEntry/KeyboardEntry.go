package main

import "fmt"

func main() {
	//获取一行的数据
	var name string
	var age byte
	var sal float32
	var isPass bool

	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水：")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试：")
	fmt.Scanln(&isPass)

	fmt.Printf("名字：%v \n年龄：%v \n薪水：%v \n是否通过考试：%v \n", name, age, sal, isPass)
}
