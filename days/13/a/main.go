package main

import (
	"fmt"
	"regexp"
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
	re := regexp.MustCompile(`([x|y])=(\d+)`)
	parsingPoints := true
	for _, row := range input {
		if parsingPoints {
			if strings.TrimSpace(row) == "" {
				parsingPoints = false
				continue
			}
			splits := lib.StringsToInts(strings.Split(row, ","))
			p := point{x: splits[0], y: splits[1]}
			dots[p] = true
		} else {
			matches := re.FindStringSubmatch(row)
			folds = append(folds, fold{
				horizontal: matches[1] == "x",
				length:     lib.MustAtoi(matches[2]),
			})
		}
	}
	visible := make(map[point]bool)
	for p := range dots {
		visible[reflect(p, folds[0])] = true
	}
	fmt.Println(len(visible))
}

func reflect(p point, f fold) (out point) {
	out = p
	if f.horizontal {
		if p.x >= f.length {
			out.x = 2*f.length - out.x
		}
	} else {
		if p.y >= f.length {
			out.y = 2*f.length - out.y
		}
	}
	return
}
