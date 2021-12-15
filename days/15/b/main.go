package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"

	"github.com/martin-svanberg-northvolt/aoc21/lib"
)

func main() {
	input := lib.GetInput()
	board := make([][]int, 0)
	for _, row := range input {
		board = append(board, lib.StringsToInts(strings.Split(row, "")))
	}
	tileWidth, tileHeight := len(board[0]), len(board)
	fmt.Println(leastRiskyPath(board, tileWidth*5, tileHeight*5))
}

func leastRiskyPath(board [][]int, width, height int) int {
	tileWidth, tileHeight := len(board[0]), len(board)
	target := [2]int{width - 1, height - 1}
	risks := make(map[[2]int]int)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			risks[[2]int{x, y}] = math.MaxInt64
		}
	}
	queue := make(PriorityQueue, 0)
	heap.Push(&queue, &Item{
		value:    [2]int{0, 0},
		risk:     0,
		priority: heuristic([2]int{0, 0}, target),
	})
	heap.Init(&queue)

	for queue.Len() > 0 {
		current := heap.Pop(&queue).(*Item)
		pos := current.value
		x := pos[0]
		y := pos[1]
		if x < 0 || x >= width || y < 0 || y >= height {
			continue
		}
		risk := board[y%tileHeight][x%tileWidth] + (y / tileHeight) + (x / tileWidth)
		for risk > 9 {
			risk -= 9
		}
		if current.risk+risk < risks[pos] {
			risks[pos] = current.risk + risk
		} else {
			continue
		}
		if pos == target {
			return risks[target] - board[0][0]
		}
		for _, neighbor := range [][2]int{{x, y - 1}, {x - 1, y}, {x + 1, y}, {x, y + 1}} {
			item := &Item{
				value:    neighbor,
				risk:     risks[pos],
				priority: risks[pos] + heuristic(pos, target),
			}
			heap.Push(&queue, item)
		}
	}
	panic("No path")
}

func heuristic(p [2]int, target [2]int) int {
	return lib.Absi(p[0]-target[0]) + lib.Absi(p[1]-target[1])
}

type Item struct {
	value    [2]int
	risk     int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
