package main

import (
	"fmt"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	depths := lib.StringsToInts(lib.GetInput())
	fmt.Print(countDepthIncreases(depths))
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
