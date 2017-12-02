package main

import (
	"io/ioutil"
    "fmt"
    "strconv"
)

func day1Part2() {
    in, _ := ioutil.ReadFile("../day1_input.txt")
    s := string(in)
    total := 0
    for i, curr := range s {
        if curr == rune(s[(i+len(s)/2)%len(s)]) {
            num, _ := strconv.Atoi(string(curr))
            total += num
        }
    }
    fmt.Println(total)
}