package main

import (
	"fmt"

	"github.com/martin-svanberg-northvolt/aoc21/days/01/common"
	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	depths := common.ReadDepths(lib.GetFixturePath("input"))
	fmt.Print(common.CountDepthIncreases(depths))
}
