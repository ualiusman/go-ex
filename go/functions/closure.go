package main

func main() {

	num := 10

	squareNum := func() int {
		num *= num
		return num
	}

	println(squareNum())
	println(squareNum())
}
