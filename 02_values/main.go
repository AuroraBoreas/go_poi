package main

import (
	"fmt"
	// "math"
)

type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	p := Person{name: "ZL", age: 35, sex: "m"}
	fmt.Println(p)
	pp := &p
	fmt.Println(pp.name, pp.age, pp.sex)
	fmt.Println("c'est la vie!")

	var a bool = false
	var b int64 = 9
	if a && b > 7 {
		fmt.Println(a)
	} else if a || b > 7 {
		fmt.Println(b)
	} else {
		fmt.Println("hello")
	}
}
