package main

import (
	"fmt"
)

func main() {
	fmt.Println("halo zl")

	message := make(chan string, 2)

	message <- "hello"
	message <- "zl"

	fmt.Println(<-message)
	fmt.Println(<-message)

}
