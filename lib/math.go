package lib

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Sum(arr []int) (out int) {
	for _, a := range arr {
		out += a
	}
	return
}

func Absi(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
