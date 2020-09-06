package main

import "fmt"

func main() {
	//数组
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Numbers数组为：")
	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
	}
	fmt.Println()
	fmt.Println("Strings数组为：")
	strings := []string{"Apple", "Google", "Microsoft"}
	for i := 0; i < len(strings); i++ {
		fmt.Print(strings[i], " ")
	}
}
