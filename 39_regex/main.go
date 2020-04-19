package main

import (
	// "bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)", "peach")
	fmt.Println(match)
}
