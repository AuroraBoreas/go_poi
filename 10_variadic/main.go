package main

import (
	"fmt"
)

func main() {
	fmt.Println("halo zl")
	arrnums := []int{2, 3, 4}
	fmt.Println("sum([2,3,4]) :", sum(arrnums...))
}

func sum(nums ...int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}
