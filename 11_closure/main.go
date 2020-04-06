package main

import (
	"fmt"
)

func main() {
	// streams in LISP, generator in Python
	fmt.Println("halo zl")
	nextInt := intseq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newints := intseq()
	fmt.Println(newints())
}

func intseq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
