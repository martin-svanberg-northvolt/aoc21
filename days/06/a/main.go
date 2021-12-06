package main

import (
	"fmt"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

const DAYS = 80

func main() {
	input := lib.GetInput()
	fishes := lib.StringsToInts(strings.Split(input[0], ","))
	for i := 0; i < DAYS; i++ {
		for j, fish := range fishes {
			if fish == 0 {
				fishes[j] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[j] -= 1
			}
		}
	}
	fmt.Println(len(fishes))
}
