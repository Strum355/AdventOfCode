package main 

import (
	"fmt"
	"encoding/csv"
	"os"
	"errors"
)

func day2Part2() {
	reader := csv.NewReader(func() *os.File {
		f, _ := os.Open("../day2_input.txt")
		return f
	}())

	reader.Comma = '	'

	data, _ := reader.ReadAll()
	total := 0
	for _, row := range data {
		intRow := toInt(row)
		for i, num := range intRow {
			j, err := findDivisible(num, intRow[i+1:])
			if err == nil {
				total += j
			}
		}
	}
	fmt.Println(total)
}

func findDivisible(i int, s []int) (int, error) {
	for _, num := range s {
		if i%num == 0 {
			return i/num, nil
		}
		if num%i == 0 {
			return num/i, nil
		}
	}
	return 0, errors.New("none found")
}