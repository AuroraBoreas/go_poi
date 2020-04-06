package main

import (
	"fmt"
)

func main() {
	fmt.Println("pointer, value")
	for i := 0; i < 5; i++ {
		fmt.Println(&i, i)
	}

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, "FuzzBarz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fuzz")
		} else if i%5 == 0 {
			fmt.Println(i, "Barz")
		} else {
			continue
		}
	}

}
