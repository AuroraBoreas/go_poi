package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("halo zl")
	fmt.Println(time.Now())

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second)
		c2 <- "two"
	}()

	// set i = 3, this program will meet deadlock. but why?
	// because previously we only make two channels: c1, and c2, neither of them have buffering

	// another interesting behavious is 'received two' may come out first.

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
