package main

import "github.com/gotoronto/workshop/slices/display"

func Pic(dx, dy int) [][]int {
	// implement
	var output [][]int
	output = make([][]int, dx)

	for x:= 0; x < dx; x++{
		output[x] = make([]int, dy)
		for y := 0; y < dy; y++{
			output[x][y] = x^y
		}
	}
	return output
}

func main() {
	display.Show(Pic(80, 25))
}
