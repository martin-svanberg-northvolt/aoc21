package main

import (
	"fmt"

	"github.com/martin-svanberg-northvolt/aoc21/01/common"
)

func main() {
	depths := common.ReadDepths("./input")
	fmt.Println(common.CountDepthDiffs(depths))
}
