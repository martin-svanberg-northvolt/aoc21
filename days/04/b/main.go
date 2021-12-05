package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

type board []int

type result struct {
	b           board
	lastDraw    int
	unmarkedSum int
}

func main() {
	lines := lib.GetInput()
	draws := lib.StringsToInts(strings.Split(lines[0], ","))
	boards := parseBoards(lines)
	result := findLastBingoBoard(draws, boards)
	fmt.Print(result.lastDraw * result.unmarkedSum)
}

func parseBoards(lines []string) []board {
	boards := make([]board, 0)
	currentBoard := make([]int, 0, 25)
	for _, line := range lines[2:] {
		if strings.TrimSpace(line) == "" {
			boards = append(boards, currentBoard)
			currentBoard = make([]int, 0, 25)
			continue
		}
		spaces := regexp.MustCompile("  | ")
		currentBoard = append(currentBoard, lib.StringsToInts(spaces.Split(line, -1))...)
	}
	return boards
}

func unmarkedSum(b board, marked []int) (total int) {
	for _, n := range b {
		if !isMarked(n, marked) {
			total += n
		}
	}
	return
}

func findLastBingoBoard(draws []int, boards []board) result {
	var lastResult result
	marked := make([]int, 0)
	wonBoards := make(map[int]bool, 0)
	for _, draw := range draws {
		for i, b := range boards {
			if wonBoards[i] {
				continue
			}
			marked = append(marked, draw)
			if markedHorizontal(b, marked) || markedVertical(b, marked) {
				lastResult = result{
					b:           b,
					lastDraw:    draw,
					unmarkedSum: unmarkedSum(b, marked),
				}
				wonBoards[i] = true
			}
		}
	}
	return lastResult
}

func markedHorizontal(b board, marked []int) bool {
	for i := 0; i < 25; i += 5 {
		if isMarked(b[i], marked) &&
			isMarked(b[i+1], marked) &&
			isMarked(b[i+2], marked) &&
			isMarked(b[i+3], marked) &&
			isMarked(b[i+4], marked) {
			return true
		}
	}
	return false
}

func markedVertical(b board, marked []int) bool {
	for i := 0; i < 5; i += 1 {
		if isMarked(b[i], marked) &&
			isMarked(b[i+5], marked) &&
			isMarked(b[i+10], marked) &&
			isMarked(b[i+15], marked) &&
			isMarked(b[i+20], marked) {
			return true
		}
	}
	return false
}

func isMarked(n int, marked []int) bool {
	for _, m := range marked {
		if n == m {
			return true
		}
	}
	return false
}
