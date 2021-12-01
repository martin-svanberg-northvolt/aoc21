package main

import (
	"fmt"

	"github.com/martin-svanberg-northvolt/aoc21/01/common"
	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	fmt.Println(solution())
}

func solution() int {
	depths := common.ReadDepths(lib.GetFixturePath("input"))
	windows := windowed(depths)
	return common.CountDepthIncreases(windows)
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
