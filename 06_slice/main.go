package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	fmt.Println("hello zl")

	s := make([]string, 3)
	fmt.Println("emp: ", s)

	rand.Seed(142)
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	for i := 0; i < len(s); i++ {
		s[i] = answers[rand.Intn(len(answers))]
	}
	fmt.Println("populated slice: ", s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("s: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[2:]
	fmt.Println("sl3: ", l)

	t := []string{"c", "b", "a"}
	sort.Strings(t)
	fmt.Println("sorted t: ", t)
	joined := strings.Join(t, "~")
	fmt.Println(joined)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD: ", twoD)
}
