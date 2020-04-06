package main

import (
	"fmt"
)

func main() {
	fmt.Println("halo zl")

	concuMsg := make(chan string)

	go func() {
		concuMsg <- "ping"
	}()

	msg := <-concuMsg
	fmt.Println(msg)

}
