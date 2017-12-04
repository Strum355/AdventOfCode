package main

import (
	"fmt"
	"io/ioutil"
)

func day1Part1() {
	in, _ := ioutil.ReadFile("../day1_input.txt")
	s := string(in)
	total := 0
	for i, curr := range s {
		if curr == rune(s[(i+1)%len(s)]) {
			total += int(curr - '0')
		}
	}
	fmt.Println(total)
}
