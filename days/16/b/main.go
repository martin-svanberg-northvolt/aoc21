package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

type packet struct {
	version int
	typeId  int
	payload interface{}
}

type literal struct {
	n int
}

type operator struct {
	mode       int
	subPackets []packet
}

func main() {
	input := lib.GetInput()[0]
	var bitInt big.Int
	bitInt.SetString(input, 16)
	bits := bitInt.Text(2)
	bits = strings.Repeat("0", len(bits)%4) + bits
	fmt.Println(eval(parsePackets(&bits, -1)[0]))
}

func eval(p packet) int {
	if lit, ok := p.payload.(literal); ok {
		return lit.n
	}
	if op, ok := p.payload.(operator); ok {
		r := make([]int, 0)
		for _, e := range op.subPackets {
			r = append(r, eval(e))
		}
		switch p.typeId {
		case 0:
			return lib.Sum(r)
		case 1:
			return lib.Product(r)
		case 2:
			return lib.MinSlice(r)
		case 3:
			return lib.MaxSlice(r)
		case 5:
			if r[0] > r[1] {
				return 1
			} else {
				return 0
			}
		case 6:
			if r[0] < r[1] {
				return 1
			} else {
				return 0
			}
		case 7:
			if r[0] == r[1] {
				return 1
			} else {
				return 0
			}
		}
	}
	panic("Unreachable")
}

func take(bits *string, n int) int {
	if len(*bits) < n {
		return -1
	}
	i, _ := strconv.ParseInt((*bits)[:n], 2, 64)
	*bits = (*bits)[n:]
	return int(i)
}

func parsePackets(bits *string, max int) []packet {
	if max == -1 {
		max = math.MaxInt64
	}
	packets := make([]packet, 0)
	for i := 0; len(*bits) > 0 && (i < max); i++ {
		var p packet
		p.version = take(bits, 3)
		p.typeId = take(bits, 3)
		if p.version < 0 || p.typeId < 0 {
			break
		}
		if p.typeId == 4 {
			var lit literal
			b := take(bits, 5)
			lit.n = b & 0b1111
			for b>>4 == 1 {
				b = take(bits, 5)
				lit.n <<= 4
				lit.n |= b & 0b1111
			}
			p.payload = lit
		} else {
			var op operator
			op.mode = take(bits, 1)
			if op.mode == 0 {
				length := take(bits, 15)
				r := (*bits)[:length]
				take(bits, length)
				op.subPackets = parsePackets(&r, -1)
			} else {
				length := take(bits, 11)
				op.subPackets = parsePackets(bits, length)
			}
			p.payload = op
		}
		packets = append(packets, p)
	}
	return packets
}
