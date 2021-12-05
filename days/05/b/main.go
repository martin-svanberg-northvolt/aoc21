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
	grid := make([]int, (width+1)*(height+1))
	for _, line := range lines {
		startY := min(line.start.y, line.end.y)
		endY := max(line.start.y, line.end.y)
		startX := min(line.start.x, line.end.x)
		endX := max(line.start.x, line.end.x)
		length := endX - startX
		if line.start.x == line.end.x {
			length = endY - startY
		}
		for i := 0; i <= length; i++ {
			x := interpolate(line.start.x, line.end.x, i)
			y := interpolate(line.start.y, line.end.y, i)
			grid[y*height+x] += 1
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

func interpolate(s int, e int, t int) int {
	if s == e {
		return s
	} else if s < e {
		return s + t
	} else {
		return s - t
	}
}
