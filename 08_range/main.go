package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world!")

	nums := []int{3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	m := map[string]string{
		"a": "apple",
		"b": "orange",
		"c": "banana",
		"d": "applepine",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for i, c := range "Golang" {
		fmt.Println(i, c)
	}

}
