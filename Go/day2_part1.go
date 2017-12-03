package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func day2Part1() {
	reader := csv.NewReader(func() *os.File {
		f, _ := os.Open("../day2_input.txt")
		return f
	}())

	reader.Comma = '	'

	data, _ := reader.ReadAll()
	total := 0
	for _, row := range data {
		intRow := toInt(row)
		min := min(intRow)
		max := max(intRow)
		total += max - min
	}
	fmt.Println(total)
}

func toInt(s []string) []int {
	i := make([]int, len(s))
	for j, num := range s {
		i[j] = func(s string) int {
			k, _ := strconv.Atoi(s)
			return k
		}(num)
	}
	return i
}

func min(s []int) int {
	min := s[0]
	for _, curr := range s[1:] {
		if curr < min {
			min = curr
		}
	}
	return min
}

func max(s []int) int {
	max := s[0]
	for _, curr := range s[1:] {
		if curr > max {
			max = curr
		}
	}
	return max
}
