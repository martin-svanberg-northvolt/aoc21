package common

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func CountDepthIncreases(depths []int) int {
	lastDepth := -1
	diffs := 0
	for _, depth := range depths {
		if lastDepth != -1 {
			diffs += diffDepths(lastDepth, depth)
		}
		lastDepth = depth
	}
	return diffs
}

func diffDepths(prev int, next int) int {
	if prev < next {
		return 1
	}
	return 0
}

func ReadDepths(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	depths := make([]int, 0, 1000)
	for scanner.Scan() {
		line := scanner.Text()
		depth, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, depth)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return depths
}
