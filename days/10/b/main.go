package main

import (
	"fmt"
	"sort"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	input := lib.GetInput()
	scores := make([]int, 0)
	incompletes := make([][]rune, 0)
outer:
	for _, line := range input {
		stack := make([]rune, 0)
		for _, char := range line {
			if char == '(' || char == '[' || char == '{' || char == '<' {
				stack = append(stack, char)
			} else {
				popped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if matchingBracket(popped) != char {
					continue outer
				}
			}
		}
		incompletes = append(incompletes, stack)
	}
	for _, stack := range incompletes {
		lineScore := 0
		for len(stack) > 0 {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			lineScore *= 5
			lineScore += bracketScore(matchingBracket(popped))
		}
		scores = append(scores, lineScore)
	}
	sort.Ints(scores)
	fmt.Println(scores[(len(scores)-1)/2+1])
}

func matchingBracket(left rune) rune {
	if left == '(' {
		return ')'
	} else if left == '[' {
		return ']'
	} else if left == '{' {
		return '}'
	} else if left == '<' {
		return '>'
	}
	panic("Invalid bracket")
}

func bracketScore(right rune) int {
	if right == ')' {
		return 1
	} else if right == ']' {
		return 2
	} else if right == '}' {
		return 3
	} else if right == '>' {
		return 4
	}
	panic("Invalid bracket")
}
