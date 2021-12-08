package main

import (
	"fmt"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	input := lib.GetInput()
	total := 0
	for _, row := range input {
		splits := strings.Split(row, "|")
		right := strings.Fields(splits[1])
		for _, item := range right {
			length := len(item)
			if (length >= 2 && length <= 4) || length == 7 {
				total += 1
			}
		}
	}
	fmt.Println(total)
}
