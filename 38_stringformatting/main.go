package main

import (
	"fmt"
	"os"
)

var pf = fmt.Printf

type point struct {
	x, y int
}

func main() {
	p :=point{1, 2}
	pf("%v\n", p)

	pf("%+v\n", p)
	pf("%#v\n", p)
	pf("%T\n", p)
	pf("%t\n", true)
	pf("%c\n", 33)
	pf("%d\n", 10)
	pf("%b\n", 10)
	pf("%x\n", 10)
	pf("%f\n", 10.0)
	pf("%e\n", 100000000000.0)
	pf("%E\n", 100000000000.0)

	pf("%s\n", "\"string\"")
	pf("%q\n", "\"string\"")
	pf("%x\n", "hex this")
	pf("%p\n", &p)

	pf("|%6d|%6d|\n", 12, 3456758)
	pf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	pf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	pf("|%6s|%6s|\n", "foo", "bar")
	pf("|%-6s|%-6s|\n", "foo", "bar")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
