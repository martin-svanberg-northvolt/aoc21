package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	input := lib.StringsToInts(strings.Split(lib.GetInput()[0], ","))
	minFuel := math.MaxInt
	for _, i := range input {
		totalFuel := 0
		for _, align := range input {
			d := abs(align - i)
			totalFuel += d * (d + 1) / 2
		}
		if totalFuel < minFuel {
			minFuel = totalFuel
		}
	}
	fmt.Print(minFuel)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
