package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	input := lib.GetInput()
	heights := make([][]int, 0)
	for _, row := range input {
		heights = append(heights, lib.StringsToInts(strings.Split(row, "")))
	}
	total := 0
	for y, row := range heights {
		for x, value := range row {
			if isLowPoint(heights, x, y) {
				total += 1 + value
			}
		}
	}
	fmt.Println(total)
}

func isLowPoint(input [][]int, x, y int) bool {
	lowest := math.MaxInt
	this := input[y][x]
	if x > 0 {
		lowest = min(lowest, input[y][x-1])
	}
	if x < len(input[0])-1 {
		lowest = min(lowest, input[y][x+1])
	}
	if y > 0 {
		lowest = min(lowest, input[y-1][x])
	}
	if y < len(input)-1 {
		lowest = min(lowest, input[y+1][x])
	}
	return this < lowest
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
