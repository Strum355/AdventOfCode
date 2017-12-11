package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func day8Part1() {
	b, err := ioutil.ReadFile("../day8_input.txt")
	if err != nil {
		panic(err)
	}

	rows := func() [][]string {
		r := strings.Split(string(b), "\n")
		out := make([][]string, len(r))
		for i, j := range r {
			out[i] = strings.Split(j, " ")
		}
		return out
	}()

	registers := make(map[string]int, len(rows)) // some over compensation but thats not the end of the world

	//init the map
	for _, r := range rows {
		registers[r[0]] = 0
	}

	for _, r := range rows {
		r1 := r[0]
		op := r[1]
		amount1, _ := strconv.Atoi(r[2])
		r2 := r[4]
		check := r[5]
		amount2, _ := strconv.Atoi(r[6])

		switch check {
		case "==":
			if registers[r2] != amount2 {
				continue
			}
		case "!=":
			if registers[r2] == amount2 {
				continue
			}
		case "<":
			if registers[r2] >= amount2 {
				continue
			}
		case ">":
			if registers[r2] <= amount2 {
				continue
			}
		case ">=":
			if registers[r2] < amount2 {
				continue
			}
		case "<=":
			if registers[r2] > amount2 {
				continue
			}
		}

		switch op {
		case "inc":
			registers[r1] += amount1
		case "dec":
			registers[r1] -= amount1
		}
	}

	max := 0
	for _, val := range registers {
		if val > max {
			max = val
		}
	}
	fmt.Println(max)
}
