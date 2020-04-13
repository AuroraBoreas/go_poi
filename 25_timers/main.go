package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("halo zl")

	timer1 := time.NewTimer(time.Second * 2)
	// capital C means blocking on timer1's channel til it sends a value out
	<-timer1.C
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 stopped")
	}
}
