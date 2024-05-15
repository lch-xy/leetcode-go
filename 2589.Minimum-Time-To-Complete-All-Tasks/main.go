package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("findMinimumTime([][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}}): %v\n", findMinimumTime([][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}}))
	fmt.Printf("findMinimumTime([][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}}): %v\n", findMinimumTime([][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}}))
}

// first step : sorted tasks by end time asc and start time asc
// second step : create visited arrays to record the time that had already beed spent
// third step : expend the window as large as possible to make next interval has more numbers in previous interval
func findMinimumTime(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i][1] == tasks[j][1] {
			return tasks[i][0] < tasks[j][0]
		}
		return tasks[i][1] < tasks[j][1]
	})

	visited := make([]int, 2001)
	res := 0
	for i := 0; i < len(tasks); i++ {
		task := tasks[i]
		start := task[0]
		end := task[1]
		duration := task[2]
		for j := start; j <= end; j++ {
			if visited[j] == 1 {
				duration--
				continue
			}
			if end-j+1 == duration {
				visited[j] = 1
				duration--
				res++
			}
		}
	}
	return res
}
