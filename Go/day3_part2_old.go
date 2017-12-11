package main

/*

!!!!!!!!!WARNING!!!!!!!!
THIS CODE DOES NOT WORK
UPLOADED JUST FOR WHY NOT

*/

import (
	"fmt"

	matrix "gonum.org/v1/gonum/mat"
)

func day3Part2_old() {
	n := 361527

	mat := matrix.NewDense(2, 2, []float64{4, 2, 1, 1})
	fmt.Print(mat.At(0, 0), " ")
	fmt.Println(mat.At(1, 0))
	fmt.Print(mat.At(0, 1), " ")
	fmt.Println(mat.At(1, 1))
	fmt.Println()

	rot := 2 //0 right 1 up 2 left 3 down

	x, y := 0, 0

	lastResult := 4.0

	for int(lastResult) < n {
		xSize, ySize := mat.Dims()
		fmt.Printf("XSize %d YSize %d\n", xSize, ySize)
		fmt.Println("Last", lastResult)
		fmt.Println(*mat)
		fmt.Printf("Rot: %d\n", rot)
		switch rot {
		case 0:
			if x == xSize-1 {
				fmt.Println("growing")
				fmt.Println(mat.Dims())
				mat = mat.Grow(0, 1).(*matrix.Dense)
				fmt.Println(mat.Dims())
				fmt.Println("done")
				rot = (rot + 1) % 4
			}
			x++
			mat.Set(x, y, addNums(x, y, mat))
		case 1:
			if y == 0 {
				mat = moveRows(ySize, mat)
				xS, xY := mat.Dims()
				fmt.Printf("did move? x %d y %d", xS, xY)
				rot = (rot + 1) % 4
				y++
			}
			y--
			mat.Set(x, y, addNums(x, y, mat))
		case 2:
			if x == 0 {
				mat = moveColumns(ySize, mat)
				xS, xY := mat.Dims()
				fmt.Printf("did move? x %d y %d", xS, xY)
				rot = (rot + 1) % 4
				x++
			}
			x--
			mat.Set(x, y, addNums(x, y, mat))
		case 3:
			if y == ySize-1 {
				fmt.Println("growing")
				fmt.Println(mat.Dims())
				mat = mat.Grow(1, 0).(*matrix.Dense)
				fmt.Println(mat.Dims())
				fmt.Println("done")
				rot = (rot + 1) % 4
			}
			y++
			mat.Set(x, y, addNums(x, y, mat))
		}
		lastResult = mat.At(x, y)
		fmt.Println()
	}

	fmt.Println()
}

func moveColumns(xSize int, m *matrix.Dense) *matrix.Dense {
	xSize, ySize := m.Dims()
	fmt.Printf("Moving and growing columns. Prev size %d %d\n", xSize, ySize)
	mat := m.Grow(0, 1).(*matrix.Dense)
	xSize, ySize = mat.Dims()
	fmt.Printf("After size %d %d\n", xSize, ySize)
	for i := xSize - 2; i > -1; i-- {
		slice := vecToSlice(mat.ColView(i))
		for j := 0; j < len(slice); j++ {
			mat.Set(j, i+1, slice[j])
		}
	}
	return mat
}

func moveRows(ySize int, m *matrix.Dense) *matrix.Dense {
	xSize, ySize := m.Dims()
	fmt.Printf("Moving and growing rows. Prev size %d %d\n", xSize, ySize)
	mat := m.Grow(1, 0).(*matrix.Dense)
	xSize, ySize = mat.Dims()
	fmt.Printf("After size %d %d\n", xSize, ySize)
	for i := ySize - 2; i > -1; i-- {
		slice := vecToSlice(mat.RowView(i))
		for j := 0; j < len(slice); j++ {
			mat.Set(i+1, j, slice[j])
		}
	}
	return mat
}

func vecToSlice(v matrix.Vector) []float64 {
	out := []float64{}
	xDim, yDim := v.Dims()
	fmt.Printf("Vec dims x %d y %d\n", xDim, yDim)
	for i := 0; i < v.Len(); i++ {
		out = append(out, v.At(i, 0))
	}
	return out
}

func addNums(x, y int, m *matrix.Dense) float64 {
	total := 0.0
	xSize, ySize := m.Dims()
	fmt.Printf("\nX %d Y %d Xsize %d Ysize %d\n", x, y, xSize, ySize)
	switch {
	case x == 0:
		switch {
		case y == 0:
			total += m.At(x+1, y)
			total += m.At(x, y+1)
			total += m.At(x+1, y+1)
		case y == ySize-1:
			total += m.At(x+1, y)
			total += m.At(x, y-1)
			total += m.At(x+1, y-1)
		default:
			total += m.At(x+1, y)
			total += m.At(x, y-1)
			total += m.At(x, y+1)
			total += m.At(x+1, y+1)
			total += m.At(x+1, y-1)
		}
	case x == xSize-1:
		switch {
		case y == 0:
			total += m.At(x-1, y)
			total += m.At(x, y+1)
			total += m.At(x-1, y+1)
		case y == ySize-1:
			total += m.At(x-1, y)
			total += m.At(x, y-1)
			total += m.At(x-1, y-1)
		default:
			total += m.At(x-1, y)
			total += m.At(x, y-1)
			total += m.At(x, y+1)
			total += m.At(x-1, y+1)
			total += m.At(x-1, y-1)
		}
	default:
		switch {
		case y == 0:
			total += m.At(x-1, y)
			total += m.At(x+1, y)
			total += m.At(x, y+1)
			total += m.At(x-1, y+1)
			total += m.At(x+1, y+1)
		case y == ySize-1:
			total += m.At(x-1, y)
			total += m.At(x+1, y)
			total += m.At(x, y-1)
			total += m.At(x-1, y-1)
			total += m.At(x+1, y-1)
		default:
			total += m.At(x-1, y)
			total += m.At(x+1, y)
			total += m.At(x, y-1)
			total += m.At(x, y+1)
			total += m.At(x-1, y+1)
			total += m.At(x-1, y-1)
			total += m.At(x+1, y+1)
			total += m.At(x+1, y-1)
		}
	}
	return total
}
