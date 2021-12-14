package main

import (
	"fmt"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

type insertion struct {
	where int
	char  string
}

func main() {
	input := lib.GetInput()
	template := input[0]
	rules := make([][2]string, 0)
	for _, row := range input[2:] {
		splits := strings.Split(row, " -> ")
		rules = append(rules, [2]string{splits[0], splits[1]})
	}
	for i := 0; i < 10; i++ {
		insertions := make([]insertion, 0)
		for j := range template {
			if j > len(template)-2 {
				continue
			}
			pair := template[j : j+2]
			for _, rule := range rules {
				if rule[0] == pair {
					insertions = append(insertions, insertion{where: j + 1, char: rule[1]})
					break
				}
			}
		}
		for j, insertion := range insertions {
			template = template[:insertion.where+j] + insertion.char + template[insertion.where+j:]
		}
	}
	fmt.Println(commonDifference(template))
}

func commonDifference(template string) int {
	counts := make(map[rune]int)
	for _, c := range template {
		counts[c] += 1
	}
	max, min := 0, 10000
	for _, v := range counts {
		max = lib.Max(max, v)
		min = lib.Min(min, v)
	}
	return max - min
}
