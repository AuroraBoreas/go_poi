package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("halo zl")
	fmt.Println(time.Now())

	done := make(chan bool, 1)
	go worker(done)

	<-done
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
