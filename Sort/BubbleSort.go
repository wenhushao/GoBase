package main

import "fmt"

func main() {
	numbers := []int{4, 5, 6, 1, 98, 12, 88, 23, 15}
	fmt.Print("原数组：")
	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
	}
	fmt.Println()
	fmt.Print("排序后数组：")
	num := bubbleSort(numbers)
	for i := 0; i < len(num); i++ {
		fmt.Print(num[i], " ")
	}

}
//冒泡排序
func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
