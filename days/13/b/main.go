package main

import (
	"fmt"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

type point struct {
	x, y int
}

type fold struct {
	horizontal bool
	length     int
}

func main() {
	input := lib.GetInput()
	dots := make(map[point]bool)
	folds := make([]fold, 0)
	for _, row := range input {
		if strings.Contains(row, ",") {
			splits := lib.StringsToInts(strings.Split(row, ","))
			dots[point{x: splits[0], y: splits[1]}] = true
		} else if strings.Contains(row, "=") {
			folds = append(folds, fold{
				horizontal: strings.Contains(row, "x="),
				length:     lib.MustAtoi(strings.Split(row, "=")[1]),
			})
		}
	}
	visible := make(map[point]bool)
	for _, fold := range folds {
		for p := range dots {
			visible[reflect(p, fold)] = true
		}
		dots = visible
	}
	print(visible)
}

func reflect(p point, f fold) (out point) {
	out = p
	if f.horizontal && p.x >= f.length {
		out.x = 2*f.length - out.x
	} else if !f.horizontal && p.y >= f.length {
		out.y = 2*f.length - out.y
	}
	return
}

func print(m map[point]bool) {
	for y := 0; y < 6; y++ {
		for x := 0; x < 39; x++ {
			if m[point{x, y}] {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
