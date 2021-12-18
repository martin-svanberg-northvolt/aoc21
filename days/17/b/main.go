package main

import (
	"fmt"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	minX, maxX := 240, 292
	minY, maxY := -90, -57
	count := 0
	for velY := -minY; velY > minY+maxY; velY-- {
		for velX := maxX; velX > 0; velX-- {
			x, y := 0, 0
			vx := velX
			vy := velY
			for x < maxX && !(vx == 0 && (x < minX || x > maxX)) && vy >= minY {
				x += vx
				y += vy
				vx -= lib.Sgn(vx)
				vy -= 1
				if minX <= x && x <= maxX && minY <= y && y <= maxY {
					count++
					break
				}
			}
		}
	}
	fmt.Println(count)
}
