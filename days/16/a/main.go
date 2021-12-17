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
	payload []interface{}
}

type literal struct {
	n int
}

type operator struct {
	mode       int
	subPackets []packet
}

// shh...
var total = 0

func main() {
	input := lib.GetInput()[0]
	var bitInt big.Int
	bitInt.SetString(input, 16)
	bits := bitInt.Text(2)
	bits = strings.Repeat("0", len(bits)%4) + bits
	parsePackets(&bits, -1)
	fmt.Println(total)
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
		total += p.version
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
			litInterface := make([]interface{}, 1)
			litInterface[0] = lit
			p.payload = litInterface
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
			opInterface := make([]interface{}, 1)
			opInterface[0] = op
			p.payload = opInterface
		}
		packets = append(packets, p)
	}
	return packets
}
