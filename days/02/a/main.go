package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	lines := lib.GetInput()
	x, y := 0, 0
	for _, line := range lines {
		splits := strings.Split(line, " ")
		cmd := splits[0]
		n, _ := strconv.Atoi(splits[1])
		if cmd == "forward" {
			x += n
		} else if cmd == "up" {
			y -= n
		} else if cmd == "down" {
			y += n
		}
	}
	fmt.Println(x * y)
}
