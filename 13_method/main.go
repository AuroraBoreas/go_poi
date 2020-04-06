package main

import (
	"fmt"
)

func main() {
	fmt.Println("halo zl")
	var r rect
	r = rect{length: 10, width: 5}
	fmt.Println("rect area: ", r.area())
	fmt.Println("rect perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}

type rect struct {
	length, width int
}

func (r *rect) area() int {
	return r.length * r.width
}

func (r rect) perim() int {
	return 2 * (r.length + r.width)
}
