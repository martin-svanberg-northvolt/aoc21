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
		width = lib.Max(lib.Max(width, curLine.start.x), curLine.end.x)
		height = lib.Max(lib.Max(height, curLine.start.y), curLine.end.y)
	}
	grid := make([]int, (width+1)*(height+1))
	for _, line := range lines {
		startY := lib.Min(line.start.y, line.end.y)
		endY := lib.Max(line.start.y, line.end.y)
		startX := lib.Min(line.start.x, line.end.x)
		endX := lib.Max(line.start.x, line.end.x)
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

func interpolate(s int, e int, t int) int {
	if s == e {
		return s
	} else if s < e {
		return s + t
	} else {
		return s - t
	}
}
