package main

import (
	"fmt"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

type point struct {
	x, y, z int
}

type cuboid struct {
	p1, p2 point
}

func main() {
	input := lib.GetInput()
	cuboids := make([]cuboid, 0)
	instructions := make([]bool, 0)
	initArea := cuboid{p1: point{x: -50, y: -50, z: -50}, p2: point{x: 50, y: 50, z: 50}}
	for _, row := range input {
		splits := strings.Fields(row)

		on := splits[0] == "on"
		instructions = append(instructions, on)

		coordinates := strings.Split(splits[1], ",")
		xs := lib.StringsToInts(strings.Split(coordinates[0][2:], ".."))
		ys := lib.StringsToInts(strings.Split(coordinates[1][2:], ".."))
		zs := lib.StringsToInts(strings.Split(coordinates[2][2:], ".."))
		c := cuboid{
			p1: point{x: xs[0], y: ys[0], z: zs[0]},
			p2: point{x: xs[1], y: ys[1], z: zs[1]},
		}
		_, inArea := intersect(c, initArea)
		if inArea {
			cuboids = append(cuboids, c)
		}
	}
	fmt.Println(sum(cuboids, instructions))
}

func sum(cuboids []cuboid, instructions []bool) (total int) {
	for i := 0; i < len(cuboids); i++ {
		if instructions != nil && !instructions[i] {
			continue
		}
		head, tail := cuboids[i], cuboids[i+1:]
		total += volume(head) - sum(intersections(head, tail), nil)
	}
	return
}

func volume(c cuboid) int {
	return (c.p2.x - c.p1.x + 1) * (c.p2.y - c.p1.y + 1) * (c.p2.z - c.p1.z + 1)
}

func intersect(c1, c2 cuboid) (cuboid, bool) {
	p1 := point{
		x: lib.Max(c1.p1.x, c2.p1.x),
		y: lib.Max(c1.p1.y, c2.p1.y),
		z: lib.Max(c1.p1.z, c2.p1.z),
	}
	p2 := point{
		x: lib.Min(c1.p2.x, c2.p2.x),
		y: lib.Min(c1.p2.y, c2.p2.y),
		z: lib.Min(c1.p2.z, c2.p2.z),
	}
	if p1.x < p2.x && p1.y < p2.y && p1.z < p2.z {
		return cuboid{p1, p2}, true
	}
	return cuboid{}, false
}

func intersections(head cuboid, tail []cuboid) []cuboid {
	intersections := make([]cuboid, 0)
	for _, t := range tail {
		if v, ok := intersect(head, t); ok {
			intersections = append(intersections, v)
		}
	}
	return intersections
}
