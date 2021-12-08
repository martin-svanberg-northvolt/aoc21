package main

import (
	"fmt"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	input := lib.GetInput()
	total := uint(0)
	for _, row := range input {
		splits := strings.Split(row, "|")
		left := strings.Fields(splits[0])
		right := strings.Fields(splits[1])
		segments := make([]uint, 7)
		for i := 0; i < len(segments); i++ {
			segments[uint(i)] = 0b1111111
		}
		for _, field := range left {
			length := len(field)
			item := encode(field)
			if length == 2 {
				segments[2] &= item
				segments[5] &= item
			} else if length == 3 {
				segments[0] &= item
				segments[2] &= item
				segments[5] &= item
			} else if length == 4 {
				segments[1] &= item
				segments[2] &= item
				segments[3] &= item
				segments[5] &= item
			} else if length == 5 {
				segments[0] &= item
				segments[3] &= item
				segments[6] &= item
			} else if length == 6 {
				segments[0] &= item
				segments[1] &= item
				segments[5] &= item
				segments[6] &= item
			}
		}
		for u := 0; u < 7; u++ {
			for i, values := range segments {
				if (values & (values - 1)) == 0 {
					for j := range segments {
						if i != j {
							segments[j] &= ^values
						}
					}
				}
			}
		}

		for i, item := range right {
			total += pow(10, uint(3-i)) * decode(item, invert(segments))
		}
	}
	fmt.Println(total)
}

func encode(s string) (n uint) {
	for _, c := range s {
		n |= 1 << uint(c-'a')
	}
	return n
}

func decode(s string, segments map[uint]uint) uint {
	numbers := map[uint]uint{
		0b1110111: 0,
		0b0100100: 1,
		0b1011101: 2,
		0b1101101: 3,
		0b0101110: 4,
		0b1101011: 5,
		0b1111011: 6,
		0b0100101: 7,
		0b1111111: 8,
		0b1101111: 9,
	}
	encoded := encode(s)
	decoded := uint(0)
	for i := uint(0); encoded > 0; i++ {
		decoded |= (encoded & 1) * 1 << segments[1<<i]
		encoded >>= 1
	}
	return numbers[decoded]
}

func invert(m []uint) map[uint]uint {
	out := make(map[uint]uint)
	for k, v := range m {
		out[v] = uint(k)
	}
	return out
}

func pow(a, b uint) uint {
	p := uint(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
