package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	lines := lib.GetInput()
	x, y, aim := 0, 0, 0
	for _, line := range lines {
		splits := strings.Split(line, " ")
		cmd := splits[0]
		n, _ := strconv.Atoi(splits[1])
		if cmd == "forward" {
			x += n
			y += aim * n
		} else if cmd == "up" {
			aim -= n
		} else if cmd == "down" {
			aim += n
		}
	}
	fmt.Println(x * y)
}
