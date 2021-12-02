package main

import (
	"fmt"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	depths := lib.StringsToInts(lib.GetInput())
	windows := windowed(depths)
	fmt.Print(countDepthIncreases(windows))
}

func windowed(depths []int) []int {
	out := make([]int, 0)
	for i, depth := range depths {
		if i < 2 {
			continue
		}
		out = append(out, depths[i-2]+depths[i-1]+depth)
	}
	return out
}

func countDepthIncreases(depths []int) int {
	lastDepth := -1
	out := 0
	for _, depth := range depths {
		if lastDepth != -1 {
			if lastDepth < depth {
				out += 1
			}
			out += 0
		}
		lastDepth = depth
	}
	return out
}
