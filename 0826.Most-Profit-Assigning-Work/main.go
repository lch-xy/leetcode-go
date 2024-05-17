package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("maxProfitAssignment([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}): %v\n", maxProfitAssignment([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}))
	fmt.Printf("maxProfitAssignment([]int{68,35,52,47,86}, []int{67,17,1,81,3}, []int{92,10,85,84,82}): %v\n", maxProfitAssignment([]int{68, 35, 52, 47, 86}, []int{67, 17, 1, 81, 3}, []int{92, 10, 85, 84, 82}))

}

type tuple struct {
	d    int
	p    int
	maxP int // find the maximum profit that diffculty not exceeding d
}

// first step : create a tuple , sorted by d and fill in maxP
// second step : use binary search to find the maximum profit
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	m := len(worker)

	tuples := make([]tuple, n)
	for i := 0; i < n; i++ {
		tuples[i] = tuple{
			d: difficulty[i],
			p: profit[i],
		}
	}

	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].d < tuples[j].d
	})

	curMaxProfit := 0
	for i := 0; i < n; i++ {
		curMaxProfit = max(curMaxProfit, tuples[i].p)
		tuples[i].maxP = curMaxProfit
	}

	var getMaxProfit = func(ability int) int {
		left, right := 0, len(tuples)-1
		bestIndex := -1
		for left <= right {
			mid := left + (right-left)/2
			if tuples[mid].d <= ability {
				bestIndex = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if bestIndex == -1 {
			return 0
		}
		return tuples[bestIndex].maxP
	}

	res := 0
	for i := 0; i < m; i++ {
		res += getMaxProfit(worker[i])
	}

	return res
}
