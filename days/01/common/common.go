package common

import (
	"log"
	"strconv"
)

func CountDepthIncreases(depths []int) int {
	lastDepth := -1
	diffs := 0
	for _, depth := range depths {
		if lastDepth != -1 {
			diffs += diffDepths(lastDepth, depth)
		}
		lastDepth = depth
	}
	return diffs
}

func diffDepths(prev int, next int) int {
	if prev < next {
		return 1
	}
	return 0
}

func LinesToInts(lines []string) []int {
	ints := make([]int, 0)
	for _, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, depth)
	}
	return ints
}
