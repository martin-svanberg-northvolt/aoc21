package lib

import "math"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxSlice(a []int) int {
	max := math.MinInt64
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinSlice(a []int) int {
	min := math.MaxInt64
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func Sum(arr []int) (out int) {
	for _, a := range arr {
		out += a
	}
	return
}

func Product(arr []int) (out int) {
	out = 1
	for _, a := range arr {
		out *= a
	}
	return
}

func Absi(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
