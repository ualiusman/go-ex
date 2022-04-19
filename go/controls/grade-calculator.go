package main

import "fmt"

func main() {
	fmt.Print("Enter Marks:")

	var input int
	fmt.Scanln(&input)

	if input < 0 || input > 100 {
		fmt.Println("Marks are invalid")
	} else if input >= 0 && input < 50 {
		fmt.Println("Fail")
	} else if input >= 50 && input < 60 {
		fmt.Println("D Grade")
	} else if input >= 60 && input < 70 {
		fmt.Println("C grade")
	} else if input >= 70 && input < 80 {
		fmt.Println("B grade")
	} else if input >= 80 && input < 90 {
		fmt.Println(" A Grade")
	} else if input >= 90 && input <= 100 {
		fmt.Println("A+ grade")
	}
	fmt.Printf("Mark %d", input)

}
