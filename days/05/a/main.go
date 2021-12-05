package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func main() {
	rows := lib.GetInput()
	rowRe := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	lines := make([]line, 0, len(rows))
	width, height := 0, 0
	for _, row := range rows {
		matches := rowRe.FindAllSubmatch([]byte(row), -1)
		curLine := line{
			start: point{
				x: mustAtoi(matches[0][1]),
				y: mustAtoi(matches[0][2]),
			},
			end: point{
				x: mustAtoi(matches[0][3]),
				y: mustAtoi(matches[0][4]),
			},
		}
		lines = append(lines, curLine)
		width = max(max(width, curLine.start.x), curLine.end.x)
		height = max(max(height, curLine.start.y), curLine.end.y)
	}
	grid := make([]int, width*height)
	for _, line := range lines {
		if line.start.x == line.end.x {
			start := min(line.start.y, line.end.y)
			end := max(line.start.y, line.end.y)
			for i := start; i <= end; i++ {
				grid[i*height+line.start.x] += 1
			}
		} else if line.start.y == line.end.y {
			start := min(line.start.x, line.end.x)
			end := max(line.start.x, line.end.x)
			for i := start; i <= end; i++ {
				grid[line.start.y*height+i] += 1
			}
		}
	}
	total := 0
	for _, cell := range grid {
		if cell >= 2 {
			total += 1
		}
	}
	fmt.Println(total)
}

func mustAtoi(s []byte) int {
	i, _ := strconv.Atoi(string(s))
	return i
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
