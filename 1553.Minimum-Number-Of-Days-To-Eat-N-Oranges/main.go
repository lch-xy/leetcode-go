package main

import (
	"fmt"
)

func main() {

	fmt.Printf("minDays(6): %v\n", minDays(6))
	fmt.Printf("minDays(10): %v\n", minDays(10))
	fmt.Printf("minDays(56): %v\n", minDays(56))

	fmt.Println("==================================================")

	fmt.Printf("minDaysDfs(6): %v\n", minDaysDfs(6))
	fmt.Printf("minDaysDfs(10): %v\n", minDaysDfs(10))
	fmt.Printf("minDaysDfs(56): %v\n", minDaysDfs(56))
}

// the searching route like a tree, so we can adjust angle to solving question
// the same as caculating the minimum distance that root node to leaf node
// using bfs to deal with it
// using cache to avoid duplicate caculations
func minDays(n int) int {
	if n == 0 {
		return 0
	}
	cache := make(map[int]struct{})
	queue := []int{n}
	step := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if curNode == 0 {
				return step
			}
			if curNode%2 == 0 {
				if _, ok := cache[curNode-curNode/2]; !ok {
					queue = append(queue, curNode-curNode/2)
					cache[curNode-curNode/2] = struct{}{}
				}
			}
			if curNode%3 == 0 {
				if _, ok := cache[curNode-2*curNode/3]; !ok {
					queue = append(queue, curNode-2*curNode/3)
					cache[curNode-2*curNode/3] = struct{}{}

				}
			}
			if _, ok := cache[curNode-1]; !ok {
				queue = append(queue, curNode-1)
				cache[curNode-1] = struct{}{}
			}
		}
		step++
	}
	return -1
}

func minDaysDfs(n int) int {
	cache := make(map[int]int)
	return dfs(n, &cache)
}

// normal dp will time out
// how to purning?
// not to caculate eating oranges one by one ,
// directly subtract n%2 or n%3 , pretending to eat one by one
// avoiding large number of sub-problems
func dfs(n int, cache *map[int]int) int {
	if n <= 1 {
		return n
	}
	if _, ok := (*cache)[n]; ok {
		return (*cache)[n]
	}
	minCnt := 1 + min(n%2+dfs(n-(n%2)-n/2, cache), n%3+dfs(n-(n%3)-2*(n/3), cache))
	(*cache)[n] = minCnt
	return minCnt
}

// timeout
// func minDays(n int) int {
// 	dp := make([]int, n+1)
// 	for i := 1; i <= n; i++ {
// 		minCnt := dp[i-1]
// 		if i%2 == 0 {
// 			minCnt = min(minCnt, dp[i-i/2])
// 		}
// 		if i%3 == 0 {
// 			minCnt = min(minCnt, dp[i-2*(i/3)])
// 		}
// 		dp[i] = minCnt + 1
// 	}
// 	return dp[n]
// }
