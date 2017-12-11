package main

import (
	"fmt"
)

var locations = make(map[string]int)

type location struct {
	x int
	y int
}

func (l location) String() string {
	return fmt.Sprintf("%d, %d", l.x, l.y)
}

func day3Part2() {
	n := 361527

	l := location{0, 0}
	locations[l.String()] = 1

	lastRes := 1
	rot := 0 //0 right 1 up 2 left 3 down
	moveAmount := 1
	rotationStep := 0
	movementStep := 0

	for lastRes <= n {
		switch rot {
		case 0:
			l.x++
		case 1:
			l.y--
		case 2:
			l.x--
		case 3:
			l.y++
		}

		lastRes = getSums(l, locations)

		locations[l.String()] = lastRes

		movementStep++
		if movementStep == moveAmount {
			movementStep = 0
			rot = (rot + 1) % 4
			if rotationStep == 1 {
				moveAmount++
			}
			rotationStep = (rotationStep + 1) % 2
		}
	}

	fmt.Println(lastRes)
}

func getValue(m map[string]int, l location) int {
	if val, ok := m[l.String()]; ok {
		return val
	}
	return 0
}

func getSums(l location, m map[string]int) (total int) {
	total = 0
	total += getValue(m, location{l.x + 1, l.y})
	total += getValue(m, location{l.x + 1, l.y + 1})
	total += getValue(m, location{l.x + 1, l.y - 1})
	total += getValue(m, location{l.x, l.y + 1})
	total += getValue(m, location{l.x, l.y - 1})
	total += getValue(m, location{l.x - 1, l.y})
	total += getValue(m, location{l.x - 1, l.y + 1})
	total += getValue(m, location{l.x - 1, l.y - 1})
	return
}
