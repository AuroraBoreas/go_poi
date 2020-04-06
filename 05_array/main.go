package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello array!")
	// Declare an Array Variable with a Specific Size
	var distros [5]string
	var ids [5]int64

	// Assign a Value to a Specific Element in an Array
	distros[0] = "hello"
	ids[1] = 1
	fmt.Println(distros)
	fmt.Println(ids)
	// Access a particular Element in an Array
	ms := distros[0]
	mi := ids[1]
	fmt.Println(ms)
	fmt.Println(mi)
	// Display All or Specific Elements from an Array
	for i := 0; i < len(distros); i++ {
		fmt.Println(i, distros[i])
	}
	// Initialize and Assign values to Array at the Same time
	arrList := [3]int64{1, 2, 3}
	fmt.Println(arrList)
	// Initialize an Array using Multi-Line Syntax
	temp := [5]float64{
		98.4,
		65.5,
		83.2,
	}
	fmt.Println(temp)
	names := [3]string{
		"ZL",
		"Jason",
		"Lily",
	}
	fmt.Println(names)
	// Identify Length of an Array – How many Elements are there?
	fmt.Println("length of an array names is:", len(names))
	// Default value of an Array Element
	fmt.Println(distros)
	// Loop through Array Elements using For and Range
	// Loop through an Array and get only the Values (Ignore Array Index)
	for _, v := range distros {
		fmt.Println(v)
	}
	// Initialize an Int Array elements with Number Sequence
	var seq [10]int
	cnt := 10
	for index, _ := range seq {
		seq[index] = cnt
		cnt += 5
	}
	fmt.Println(seq)
	// Define and Intialize Multi-Dimentional Arrays
	cnt = 1
	var multi [4][2]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			multi[i][j] = cnt
			cnt++
		}
	}
	fmt.Println(multi)
	// Full Working GoLang Code with All Array Examples
}
