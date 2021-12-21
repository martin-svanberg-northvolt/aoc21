package main

import "fmt"

type player struct {
	pos, score int
}

func main() {
	p1 := player{pos: 4}
	p2 := player{pos: 10}
	d := 0
	for {
		d += 3
		p1.pos = ((p1.pos-1)+3*d+7)%10 + 1
		p1.score += p1.pos
		if p1.score >= 1000 {
			break
		}
		p1, p2 = p2, p1
	}
	fmt.Println(p2.score * d)
}
