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
	return common.CountDepthIncreases(depths)
}
