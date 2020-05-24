package main

import "fmt"

func main() {
	//switch
	var b = 1
	switch b {
	case 1:
		fmt.Print("b=1")
		break
	case 2:
		fmt.Print("b=2")
		break
	case 3:
		fmt.Print("b=3")
		break
	default:
		fmt.Print(b)
		break
	}
}
