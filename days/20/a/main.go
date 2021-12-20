package main

import (
	"fmt"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

type point struct {
	x, y int
}

func main() {
	input := lib.GetInput()
	algorithm := input[0]
	image := make(map[point]int)
	for y, row := range input[2:] {
		for x, col := range row {
			if col == '#' {
				image[point{x, y}] = 1
			} else {
				image[point{x, y}] = 0
			}
		}
	}

	for i := 0; i < 2; i++ {
		image = enhance(image, algorithm)
	}

	fmt.Println(countLit(image))
}

func enhance(image map[point]int, algorithm string) map[point]int {
	newImage := make(map[point]int)
	for x := -50; x < 150; x++ {
		for y := -50; y < 150; y++ {
			c := 0
			for m := -1; m <= 1; m++ {
				for n := -1; n <= 1; n++ {
					bit, ok := image[point{x: x + n, y: y + m}]
					if !ok {
						bit = image[point{x: -50, y: -50}]
					}
					c |= bit << (8 - ((m+1)*3 + (n + 1)))
				}
			}
			if algorithm[c] == '#' {
				newImage[point{x, y}] = 1
			} else {
				newImage[point{x, y}] = 0
			}
		}
	}
	return newImage
}

func countLit(image map[point]int) (total int) {
	for _, c := range image {
		if c == 1 {
			total++
		}
	}
	return
}
