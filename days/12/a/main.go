package main

import (
	"fmt"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

type edges map[string][]string

func main() {
	input := lib.GetInput()
	edges := make(edges)
	for _, row := range input {
		splits := strings.Split(row, "-")
		a, b := splits[0], splits[1]
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}
	pathTaken := make([]string, 1)
	pathTaken[0] = "start"
	fmt.Println(countPaths("start", "end", edges, pathTaken))
}

func countPaths(from, to string, edges edges, pathTaken []string) (total int) {
	if from == to {
		return 1
	}
	nextSteps := edges[from]
	for _, next := range nextSteps {
		if isSmallCave(next) && hasVisited(next, pathTaken) {
			continue
		}
		total += countPaths(next, to, edges, append(pathTaken, next))
	}
	return
}

func isSmallCave(cave string) bool {
	return strings.ToLower(cave) == cave
}

func hasVisited(cave string, pathTaken []string) bool {
	for _, p := range pathTaken {
		if p == cave {
			return true
		}
	}
	return false
}
