package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	input := lib.GetInput()
	heights := make([][]int, 0)
	for _, row := range input {
		heights = append(heights, lib.StringsToInts(strings.Split(row, "")))
	}
	lowPoints := make([][]int, 0)
	for y, row := range heights {
		for x := range row {
			if isLowPoint(heights, x, y) {
				lowPoints = append(lowPoints, []int{x, y})
			}
		}
	}
	basins := make([]int, 0)
	for _, lp := range lowPoints {
		basins = append(basins, -countBasinSize(heights, lp[0], lp[1]))
	}
	sort.Ints(basins)
	fmt.Println(-basins[0] * basins[1] * basins[2])
}

func isLowPoint(heights [][]int, x, y int) bool {
	lowest := math.MaxInt
	this := heights[y][x]
	if x > 0 {
		lowest = min(lowest, heights[y][x-1])
	}
	if x < len(heights[0])-1 {
		lowest = min(lowest, heights[y][x+1])
	}
	if y > 0 {
		lowest = min(lowest, heights[y-1][x])
	}
	if y < len(heights)-1 {
		lowest = min(lowest, heights[y+1][x])
	}
	return this < lowest
}

func countBasinSize(input [][]int, x, y int) (size int) {
	nodes := make([][]int, 1)
	nodes[0] = []int{x, y, -1}
	visited := make([][]bool, len(input))
	for i, row := range input {
		visited[i] = make([]bool, len(row))
	}
	width, height := len(input[0]), len(input)
	for len(nodes) > 0 {
		q := nodes[0]
		x, y := q[0], q[1]
		nodes = nodes[1:]
		if x >= 0 && x < width && y >= 0 && y < height && !visited[y][x] {
			value := input[y][x]
			if q[2] < value && value != 9 {
				visited[y][x] = true
				nodes = append(nodes, [][]int{
					{x - 1, y, value},
					{x + 1, y, value},
					{x, y - 1, value},
					{x, y + 1, value},
				}...)
			}
		}
	}
	for _, row := range visited {
		for _, b := range row {
			if b {
				size += 1
			}
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
