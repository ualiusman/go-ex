package main

import "fmt"

type Employee struct {
	fname string
	lname string
}

func (emp Employee) fullName() {
	fmt.Println(emp.fname + " " + emp.lname)
}

func fun() int {
	return 32323
}

func main() {

	e1 := Employee{"Badar", "Munir"}
	e1.fullName()

	st := fun()

	fmt.Println(st)
}
