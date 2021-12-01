package main

import (
	"fmt"

	"github.com/martin-svanberg-northvolt/aoc21/days/01/common"
	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	depths := common.LinesToInts(lib.GetInput())
	windows := windowed(depths)
	fmt.Print(common.CountDepthIncreases(windows))
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
