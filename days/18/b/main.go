package main

import (
	"fmt"
	"unicode"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

const PAIR = -1

func main() {
	input := lib.GetInput()
	maxMagn := 0
	sfns := make([][]int, 0)
	for _, row := range input {
		sfns = append(sfns, parseSfn(row))
	}
	for i, n := range sfns {
		for j, k := range sfns {
			if i == j {
				continue
			}
			sum := reduce(addSfn(n, k))
			mag := magnitude(sum)
			if mag > maxMagn {
				maxMagn = mag
			}
		}
	}
	fmt.Println(maxMagn)
}

func parseSfn(s string) []int {
	out := make([]int, 0)
	num := ""
	for i, c := range s {
		if c == '[' {
			out = append(out, PAIR)
		} else if unicode.IsDigit(c) {
			num += string(c)
		} else if c == ',' || i == len(s)-1 {
			out = append(out, lib.MustAtoi(num))
			num = ""
		}
	}
	return out
}

func addSfn(a, b []int) []int {
	out := make([]int, len(a)+len(b)+1)
	out[0] = PAIR
	copy(out[1:], a)
	copy(out[1+len(a):], b)
	return out
}

func explode(sfn []int, i int) []int {
	if sfn[i] != PAIR {
		panic("explode invariant broken")
	}
	left, right := sfn[i+1], sfn[i+2]
	for j := i; j >= 0; j-- {
		if sfn[j] != PAIR {
			sfn[j] += left
			break
		}
	}
	for j := i + 3; j < len(sfn); j++ {
		if sfn[j] != PAIR {
			sfn[j] += right
			break
		}
	}
	sfn = append(append(sfn[:i], 0), sfn[i+3:]...)
	return sfn
}

func split(sfn []int, i int) []int {
	if sfn[i] == PAIR {
		panic("split invariant broken")
	}
	out := make([]int, len(sfn)+2)
	copy(out, sfn[:i])
	out[i] = PAIR
	out[i+1] = sfn[i] / 2
	out[i+2] = sfn[i]/2 + sfn[i]%2
	copy(out[i+3:], sfn[i+1:])
	return out
}

func reduce(sfn []int) []int {
	action := true
	for action {
		action = false
		pairStack := make([]int, 0)
		for i := 0; i < len(sfn); i++ {
			n := sfn[i]
			if n == PAIR {
				pairStack = append(pairStack, 0)
			} else if len(pairStack) > 0 {
				pairStack[len(pairStack)-1] += 1
				for pairStack[len(pairStack)-1] == 2 {
					pairStack = pairStack[:len(pairStack)-1]
					if len(pairStack) == 0 {
						break
					}
					pairStack[len(pairStack)-1] += 1
				}
			}
			if len(pairStack) > 4 {
				sfn = explode(sfn, i)
				action = true
				break
			}
		}
		if !action {
			for i, n := range sfn {
				if n >= 10 {
					sfn = split(sfn, i)
					action = true
					break
				}
			}
		}
	}
	return sfn
}

func magnitude(n []int) int {
	clone := make([]int, len(n))
	copy(clone, n)
	for i := 0; len(clone) > 1; i = (i + 1) % len(clone) {
		if i < len(clone)-2 && clone[i] == PAIR && clone[i+1] != PAIR && clone[i+2] != PAIR {
			clone[i] = 3*clone[i+1] + 2*clone[i+2]
			clone = append(clone[:i+1], clone[i+3:]...)
		}
	}
	return clone[0]
}
