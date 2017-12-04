package main 

import (
	"strings"
	"io/ioutil"
	"fmt"
	"sort"
)

func day4Part2() {
	bytes, _ := ioutil.ReadFile("../day4_input.txt")

	s := strings.Split(string(bytes), "\n")
	data := func() [][]string {
		out := make([][]string, 0)
		for _, row := range s {
			out = append(out, strings.Split(row, " "))
		}
		return out
	}()

	total := 0
	for _, row := range data {
		unique := make(map[string]bool)
		for _, value := range row {
			s := strings.Split(value, "")
			sort.Strings(s)
			unique[strings.Join(s, "")] = true
		}
		if len(unique) == len(row) {
			total++
		}
	}
	fmt.Println(total)
}