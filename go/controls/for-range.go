package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, value := range nums {
		sum += value
	}

	fmt.Println("sum ", sum)

	for i, value := range nums {
		if value == 3 {
			fmt.Println("3 index ", i)
		}
	}

	kvs := map[string]string{"1": "mango", "2": "Raja", "3": "family"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//Single line comment
	/*multiline comments */

	for k := range kvs {
		fmt.Println(k)
	}

	for i, c := range "hi" {
		fmt.Println(i, c)
	}
}
