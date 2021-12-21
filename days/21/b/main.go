package main

import (
	"fmt"
)

type player struct {
	pos, score int
}

type state struct {
	p1, p2 player
}

var cache map[state][2]int

func main() {
	cache = make(map[state][2]int)
	state := state{p1: player{pos: 4}, p2: player{pos: 10}}
	fmt.Println(count(state)[0])
}

func count(s state) [2]int {
	if v, ok := cache[s]; ok {
		return v
	}
	out := [2]int{0, 0}
	if s.p2.score >= 21 {
		out = [2]int{0, 1}
	} else {
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				for k := 1; k <= 3; k++ {
					d := i + j + k
					pos := (s.p1.pos+d-1)%10 + 1
					score := s.p1.score + pos
					c := count(state{p1: s.p2, p2: player{pos, score}})
					out = [2]int{out[0] + c[1], out[1] + c[0]}
				}
			}
		}
	}
	cache[s] = out
	return out
}
