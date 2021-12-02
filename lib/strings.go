package lib

import (
	"log"
	"strconv"
)

func StringsToInts(lines []string) []int {
	ints := make([]int, 0)
	for _, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			log.Println(err)
		}
		ints = append(ints, depth)
	}
	return ints
}
