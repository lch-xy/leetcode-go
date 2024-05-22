package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}): %v\n", findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
}

func findWinners(matches [][]int) [][]int {
	n := len(matches)
	winCache := make(map[int]int)
	loseCache := make(map[int]int)
	for i := 0; i < n; i++ {
		winCache[matches[i][0]]++
		loseCache[matches[i][1]]++
	}
	onceLose := make([]int, 0)
	for key, value := range loseCache {
		delete(winCache, key)
		if value == 1 {
			onceLose = append(onceLose, key)
		}
	}
	winner := make([]int, 0)
	for key, _ := range winCache {
		winner = append(winner, key)
	}
	sort.Ints(winner)
	sort.Ints(onceLose)
	return [][]int{winner, onceLose}
}
