package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Printf("totalCost([]int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}, 11, 2): %v\n", totalCost([]int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}, 11, 2))
	fmt.Printf("totalCost([]int{1,2,4,1}, 3, 3): %v\n", totalCost([]int{1, 2, 4, 1}, 3, 3))
	// h := &SortHeap{}
	// heap.Init(h)
	// heap.Push(h, 2)
	// heap.Push(h, 1)
	// heap.Push(h, -1)
	// heap.Push(h, 3)
	// fmt.Printf("h.IntSlice[0]: %v\n", h.IntSlice[0])
	// fmt.Printf("h.Pop(): %v\n", h.Pop())
}

// use PriorityQueue to pop the minimums node
func totalCost(costs []int, k int, candidates int) int64 {
	res := int64(0)
	sortHeap := &nums{}
	heap.Init(sortHeap)

	leftIndex := 0
	rightIndex := len(costs) - 1
	for leftIndex < rightIndex && candidates > 0 {
		heap.Push(sortHeap, pair{costs[leftIndex], 0})
		heap.Push(sortHeap, pair{costs[rightIndex], 1})
		candidates--
		leftIndex++
		rightIndex--
	}

	for i := 0; i < k; i++ {
		minPair := heap.Pop(sortHeap).(pair)
		res += int64(minPair.num)
		if minPair.tp == 0 && leftIndex <= rightIndex {
			heap.Push(sortHeap, pair{costs[leftIndex], 0})
			leftIndex++
		}
		if minPair.tp == 1 && leftIndex <= rightIndex {
			heap.Push(sortHeap, pair{costs[rightIndex], 1})
			rightIndex--
		}
	}
	return res
}

type nums []pair

// tp： it's to distinguish left and right nodes
type pair struct {
	num int
	tp  int
}

func (n nums) Len() int {
	return len(n)
}

func (n nums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n nums) Less(i, j int) bool {
	if n[i].num == n[j].num {
		return n[i].tp < n[j].tp
	}
	return n[i].num < n[j].num
}

func (n *nums) Push(x any) {
	*n = append(*n, x.(pair))
}

func (n *nums) Pop() any {
	l := len(*n)
	x := (*n)[l-1]
	*n = (*n)[:l-1]
	return x
}

// time out
// func totalCost(costs []int, k int, candidates int) int64 {
// 	res := 0
// 	visited := make(map[int]bool)
// 	for i := 0; i < k; i++ {
// 		leftCandidates := 0
// 		rightCandidates := 0
// 		minData := math.MaxInt32
// 		minIndex := -1
// 		for j := 0; j < len(costs); j++ {
// 			if leftCandidates == candidates {
// 				break
// 			}
// 			if visited[j] {
// 				continue
// 			}
// 			if minData > costs[j] {
// 				minData = costs[j]
// 				minIndex = j
// 			}
// 			leftCandidates++
// 		}

// 		for j := len(costs) - 1; j >= 0; j-- {
// 			if rightCandidates == candidates {
// 				break
// 			}
// 			if visited[j] {
// 				continue
// 			}
// 			if minData > costs[j] {
// 				minData = costs[j]
// 				minIndex = j
// 			}
// 			rightCandidates++
// 		}
// 		res += minData
// 		visited[minIndex] = true
// 	}
// 	return int64(res)
// }
