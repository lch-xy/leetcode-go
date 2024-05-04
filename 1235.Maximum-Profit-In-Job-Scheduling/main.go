package main

import (
	"fmt"
	"sort"
)

func main() {
	sss := tuples{{1, 3, 50}, {2, 4, 10}, {3, 5, 40}, {3, 6, 70}}
	sort.Sort(sss)
	fmt.Printf("sss: %v\n", sss)
	fmt.Printf("findLastNoConflictJob(sss, 3, 3): %v\n", findLastNoConflictJob(sss, 3, 3))
}

// one step: make tuple including startTime、endTime、profit
// two step: sorted by endTime
// three step: use binary search to find the last no conflict index
// four step: dp[i] = max(dp[i-1],dp[lastNoConflictIndex+1] + tuples[i-1].profit)
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	length := len(startTime)
	tt := tuples{}
	for i := 0; i < length; i++ {
		tt = append(tt,
			tuple{
				startTime: startTime[i],
				endTime:   endTime[i],
				profit:    profit[i],
			})
	}
	sort.Sort(tt)
	dp := make([]int, length+1)
	for i := 1; i <= length; i++ {
		lastNoConflictIndex := findLastNoConflictJob(tt, tt[i-1].startTime, i-1)
		// index mapping to dp should add 1
		// dp[lastNoConflictIndex+1]: the lastNoConflictIndex-th position in tuples
		dp[i] = max(dp[i-1], dp[lastNoConflictIndex+1]+tt[i-1].profit)
	}

	return dp[length]
}

// find last no conflict index
func findLastNoConflictJob(tuples []tuple, target, endIndex int) int {
	left, right := 0, endIndex+1
	for left < right {
		mid := left + (right-left)/2
		if tuples[mid].endTime > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right - 1
}

type tuples []tuple

type tuple struct {
	startTime int
	endTime   int
	profit    int
}

func (t tuples) Len() int {
	return len(t)
}

func (t tuples) Less(i, j int) bool {
	return t[i].endTime < t[j].endTime
}

func (t tuples) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
