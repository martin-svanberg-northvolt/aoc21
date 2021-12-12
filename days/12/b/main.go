package main

import (
	"fmt"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

type edges map[string][]string

type pathTaken struct {
	path           []string
	smallCaveTwice bool
}

func main() {
	input := lib.GetInput()
	edges := make(edges)
	for _, row := range input {
		splits := strings.Split(row, "-")
		a, b := splits[0], splits[1]
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}
	pathTaken := pathTaken{
		path:           make([]string, 1),
		smallCaveTwice: false,
	}
	pathTaken.path[0] = "start"
	fmt.Println(countPaths("start", "end", edges, pathTaken))
}

func countPaths(from, to string, edges edges, p pathTaken) (total int) {
	if from == to {
		return 1
	}
	nextSteps := edges[from]
	for _, next := range nextSteps {
		if isSmallCave(next) && hasVisited(next, p) {
			if !p.smallCaveTwice && next != "end" && next != "start" {
				total += countPaths(next, to, edges,
					pathTaken{
						path:           append(p.path, next),
						smallCaveTwice: true,
					},
				)
			}
			continue
		}
		total += countPaths(next, to, edges,
			pathTaken{
				path:           append(p.path, next),
				smallCaveTwice: p.smallCaveTwice,
			},
		)
	}
	return
}

func isSmallCave(cave string) bool {
	return strings.ToLower(cave) == cave
}

func hasVisited(cave string, pathTaken pathTaken) bool {
	for _, p := range pathTaken.path {
		if p == cave {
			return true
		}
	}
	return false
}
