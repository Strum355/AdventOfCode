package main

import (
	"fmt"
)

func day3Part1() {
	n := 361527
	x, y := 0, 0
	rot := 3 //0 right 1 up 2 left 3 down
	toInc := 1
	i := 1
	nextMoveAmount := 1
	for i < n {
		rot = (rot + 1) % 4
		toInc = (toInc + 1) % 2
		switch rot {
		case 0:
			x += nextMoveAmount
		case 1:
			y += nextMoveAmount
		case 2:
			x -= nextMoveAmount
		case 3:
			y -= nextMoveAmount
		}
		i += nextMoveAmount
		if toInc == 1 {
			nextMoveAmount++
		}
	}

	diff := i - n
	switch rot {
	case 0:
		x -= diff
	case 1:
		y -= diff
	case 2:
		x += diff
	case 3:
		y += diff
	}
	fmt.Println(abs(x) + abs(y))
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
