package main

import (
	"fmt"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	lines := lib.GetInput()
	totals := make([]int, 12)
	for _, line := range lines {
		for i, c := range []byte(line) {
			totals[i] += int(c - '0')
		}
	}
	gamma, epsilon := 0, 0
	for i := 0; i < 12; i++ {
		gamma <<= 1
		epsilon <<= 1
		if totals[i]-len(lines)/2 >= 0 {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}
	fmt.Print(gamma * epsilon)
}
