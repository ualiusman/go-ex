package main

import "fmt"

func main() {
	fmt.Println("enter Number:")

	var input int
	fmt.Scanln(&input)

	switch input {
	case 10:
		fmt.Println("this is 10")
	case 20:
		fmt.Println("this is 20")
	case 30:
		fmt.Println("this is 30")
	case 50:
		fmt.Println("this is 50")
	default:
		fmt.Println("no in 10,20,30, 50")
	}
}
