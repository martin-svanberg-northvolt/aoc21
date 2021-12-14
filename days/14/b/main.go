package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	input := lib.GetInput()
	template := input[0]
	rules := make(map[uint16]byte, 0)
	for _, row := range input[2:] {
		splits := strings.Split(row, " -> ")
		a, b, c := splits[0][0], splits[0][1], splits[1][0]
		rules[pairKey(a, b)] = c
	}
	pairs := make(map[uint16]int)
	for i := range template {
		if i == len(template)-1 {
			break
		}
		pairs[pairKey(template[i], template[i+1])] += 1
	}
	for step := 0; step < 40; step++ {
		next := make(map[uint16]int)
		for pair, count := range pairs {
			insertion := rules[pair]
			next[pairKey(left(pair), insertion)] += count
			next[pairKey(insertion, right(pair))] += count
		}
		pairs = next
	}
	fmt.Println(commonDifference(pairs))
}

func commonDifference(pairs map[uint16]int) int {
	counts := make(map[byte]int)
	for p, v := range pairs {
		counts[right(p)] += v
	}
	max, min := 0, math.MaxInt64
	for _, v := range counts {
		max = lib.Max(max, v)
		min = lib.Min(min, v)
	}
	return max - min
}

func pairKey(a, b byte) uint16 {
	return (uint16(a) << 8) | uint16(b)
}

func left(rule uint16) byte {
	return byte((rule >> 8) & 0xFF)
}

func right(rule uint16) byte {
	return byte(rule & 0xFF)
}
