package main

import (
	"fmt"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

type vector struct {
	x, y, z int
}

func main() {
	input := lib.GetInput()
	orientations := [][]int{
		{-1, 0, 0, 0, -1, 0, 0, 0, 1},
		{-1, 0, 0, 0, 0, -1, 0, -1, 0},
		{-1, 0, 0, 0, 0, 1, 0, 1, 0},
		{-1, 0, 0, 0, 1, 0, 0, 0, -1},
		{0, -1, 0, -1, 0, 0, 0, 0, -1},
		{0, -1, 0, 0, 0, -1, 1, 0, 0},
		{0, -1, 0, 0, 0, 1, -1, 0, 0},
		{0, -1, 0, 1, 0, 0, 0, 0, 1},
		{0, 0, -1, -1, 0, 0, 0, 1, 0},
		{0, 0, -1, 0, -1, 0, -1, 0, 0},
		{0, 0, -1, 0, 1, 0, 1, 0, 0},
		{0, 0, -1, 1, 0, 0, 0, -1, 0},
		{0, 0, 1, -1, 0, 0, 0, -1, 0},
		{0, 0, 1, 0, -1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1, 0, -1, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, -1, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, -1, -1, 0, 0},
		{0, 1, 0, 0, 0, 1, 1, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0, -1},
		{1, 0, 0, 0, -1, 0, 0, 0, -1},
		{1, 0, 0, 0, 0, -1, 0, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, -1, 0},
		{1, 0, 0, 0, 1, 0, 0, 0, 1},
	}
	scanners := make([][]vector, 0)
	for _, row := range input {
		if strings.Contains(row, "---") {
			scanners = append(scanners, make([]vector, 0))
		} else if strings.Contains(row, ",") {
			splits := strings.Split(row, ",")
			scanners[len(scanners)-1] = append(scanners[len(scanners)-1], vector{
				x: lib.MustAtoi(splits[0]),
				y: lib.MustAtoi(splits[1]),
				z: lib.MustAtoi(splits[2]),
			})
		}
	}
	oriented := make([]int, 1, len(scanners))
	oriented[0] = 0
	unoriented := make(map[int]bool, len(scanners))
	total := make(map[vector]bool, 500)
	for i := range scanners {
		unoriented[i] = true
	}
	for len(unoriented) > 0 {
		reference := scanners[oriented[0]]
		oriented = oriented[1:]
		for unorientedIdx := range unoriented {
			unorientedScanner := scanners[unorientedIdx]
			for _, t := range orientations {
				transformed := make([]vector, len(unorientedScanner))
				for i := range unorientedScanner {
					transformed[i] = mul(t, unorientedScanner[i])
				}
				translation, ok := findTranslation(reference, transformed)
				if ok {
					for j := range unorientedScanner {
						v := sub(transformed[j], translation)
						scanners[unorientedIdx][j] = v
						total[v] = true
					}
					delete(unoriented, unorientedIdx)
					oriented = append(oriented, unorientedIdx)
					break
				}
			}
		}
	}
	fmt.Println(len(total))
}

func findTranslation(s1, s2 []vector) (vector, bool) {
	ds := make(map[vector]int, len(s1)*len(s2))
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			k := sub(v2, v1)
			ds[k] += 1
			if ds[k] == 12 {
				return k, true
			}
		}
	}
	return vector{}, false
}

func mul(mat []int, v vector) vector {
	return vector{
		x: mat[0]*v.x + mat[1]*v.y + mat[2]*v.z,
		y: mat[3]*v.x + mat[4]*v.y + mat[5]*v.z,
		z: mat[6]*v.x + mat[7]*v.y + mat[8]*v.z,
	}
}

func sub(v1, v2 vector) vector {
	return vector{
		x: v1.x - v2.x,
		y: v1.y - v2.y,
		z: v1.z - v2.z,
	}
}
