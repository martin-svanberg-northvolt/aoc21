package main

import (
	"fmt"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

const DAYS = 256

func main() {
	input := lib.GetInput()
	daysLeft := make([]int64, 10)
	for _, j := range lib.StringsToInts(strings.Split(input[0], ",")) {
		daysLeft[int64(j)] += 1
	}
	for i := 0; i < DAYS; i++ {
		daysLeft[7] += daysLeft[0]
		daysLeft[9] += daysLeft[0]
		daysLeft = append(daysLeft[1:], 0)
	}
	var total int64
	for _, v := range daysLeft {
		total += v
	}
	fmt.Println(total)
}
