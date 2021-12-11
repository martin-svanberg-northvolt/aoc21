package main

import (
	"fmt"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

type point struct {
	X, Y int
}

func main() {
	input := lib.GetInput()
	energies := make([][]int, 0)
	flashes := 0
	for _, row := range input {
		energies = append(energies, lib.StringsToInts(strings.Split(row, "")))
	}
	for step := 0; ; step++ {
		for i, row := range energies {
			for j := range row {
				energies[i][j] += 1
			}
		}
		flashedAt := make(map[point]bool)
		anyFlashes := true
		for anyFlashes {
			anyFlashes = false
			for y, row := range energies {
				for x, energy := range row {
					point := point{X: x, Y: y}
					if energy > 9 && !flashedAt[point] {
						incAdjacent(energies, x, y)
						flashes += 1
						flashedAt[point] = true
						anyFlashes = true
					}
				}
			}
		}
		for y, row := range energies {
			for x, energy := range row {
				if energy > 9 {
					energies[y][x] = 0
				}
			}
		}
		allFlash := true
		for y, row := range energies {
			for x := range row {
				if !flashedAt[point{X: x, Y: y}] {
					allFlash = false
				}
			}
		}
		if allFlash {
			fmt.Println(step + 1)
			break
		}
	}
}

func incAdjacent(energies [][]int, x, y int) {
	adjacent := [][]int{
		{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1},
		{x - 1, y}, {x + 1, y},
		{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
	}
	width, height := len(energies[0]), len(energies)
	for _, a := range adjacent {
		ax, ay := a[0], a[1]
		if ax >= 0 && ax < width && ay >= 0 && ay < height {
			energies[ay][ax] += 1
		}
	}
}
