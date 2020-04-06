package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello zl")
	fmt.Println("factorial 5: ", factorial(5))
	fmt.Println("add(a, b) : ", add(5, 10))
	fmt.Println("sub(a, b) : ", sub(5, 3))
	fmt.Println("mul(a, b) : ", mul(5, 4))

}

func add(a, b int64) int64 {
	return a + b
}

func mul(a, b int64) int64 {
	return a * b
}

func sub(a, b int64) int64 {
	return a - b
}

func factorial(n int64) int64 {
	return f(n, 1)
}

func f(n, product int64) int64 {
	if n == 1 {
		return product
	}
	return f(n-1, n*product)
}
