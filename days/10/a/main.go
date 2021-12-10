package main

import (
	"fmt"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	input := lib.GetInput()
	total := 0
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
					total += bracketScore(char)
					continue outer
				}
			}
		}
	}
	fmt.Println(total)
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
		return 3
	} else if right == ']' {
		return 57
	} else if right == '}' {
		return 1197
	} else if right == '>' {
		return 25137
	}
	panic("Invalid bracket")
}
