package main

import (
	"fmt"
)

func main() {
	fmt.Println("halo zl")

	f("direct")
	go f("goroutine")

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")

	// fmt.Scanln()
	// fmt.Println("done")
}

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}
