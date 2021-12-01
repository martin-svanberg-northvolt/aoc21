package main

import (
	"fmt"

	"github.com/martin-svanberg-northvolt/aoc21/days/01/common"
	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	depths := common.LinesToInts(lib.GetInput())
	fmt.Print(common.CountDepthIncreases(depths))
}
