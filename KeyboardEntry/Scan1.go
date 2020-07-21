package main

import "fmt"

func main() {
	//获取一行的数据
	var name string
	var age byte
	var sal float32
	var isPass bool

	fmt.Println("请输入姓名， 年龄， 薪水 是否通过考试， 使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)

	fmt.Printf("名字：%v \n年龄：%v \n薪水：%v \n是否通过考试：%v \n", name, age, sal, isPass)
}
