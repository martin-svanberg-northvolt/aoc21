package main

import (
	"fmt"

	"github.com/martin-svanberg-northvolt/aoc21/01/common"
)

func main() {
	depths := common.ReadDepths("./input")
	windows := windowed(depths)
	fmt.Println(common.CountDepthDiffs(windows))
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
