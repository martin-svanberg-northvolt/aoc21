package main

import (
	"fmt"
	"strconv"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	lines := lib.GetInput()
	oxygen := findRating(lines, func(a int, b int) bool { return a >= b })
	co2 := findRating(lines, func(a int, b int) bool { return a < b })
	fmt.Println(oxygen * co2)
}

func findRating(input []string, comparator func(int, int) bool) int64 {
	zeros := make([]string, 0, len(input))
	ones := make([]string, 0, len(input))
	for i := 0; i < 12; i++ {
		for _, row := range input {
			if row[i] == '0' {
				zeros = append(zeros, row)
			} else {
				ones = append(ones, row)
			}
		}
		if comparator(len(ones), len(zeros)) {
			input = ones
		} else {
			input = zeros
		}
		zeros = zeros[:0]
		ones = ones[:0]
		if len(input) == 1 {
			num, _ := strconv.ParseInt(input[0], 2, 64)
			return num
		}
	}
	panic("No rating found")
}
