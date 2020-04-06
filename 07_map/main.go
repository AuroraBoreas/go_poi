package main

import (
	"fmt"
)

func main() {
	fmt.Println("hellow world!")

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("m: ", m)
	fmt.Println("len of m: ", len(m))

	delete(m, "k1")
	fmt.Println("m: ", m)

	for _, v := range m {
		fmt.Println(v)
	}

	n := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println("n: ", n)
}
