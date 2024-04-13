package main

import (
	"fmt"
)

func main() {
	type TestCase struct {
		n        int
		edges    [][]int
		expected int
	}

	testCases := []TestCase{
		{
			n:        3,
			edges:    [][]int{{0, 1}, {1, 2}},
			expected: 0,
		},
		{
			n:        3,
			edges:    [][]int{{0, 1}, {1, 2}, {0, 2}},
			expected: 0,
		},
		{
			n:        4,
			edges:    [][]int{{0, 2}, {1, 3}, {1, 2}},
			expected: -1,
		},
	}

	for i, tc := range testCases {
		result := findChampion(tc.n, tc.edges)
		if result == tc.expected {
			fmt.Printf("Test case %d passed\n", i+1)
		} else {
			fmt.Printf("Test case %d failed: expected %d, got %d\n", i+1, tc.expected, result)
			return
		}
	}
}

// to find a node that in-degree is 0
// means no one stronger than it
// if only one node meet the condition, return it.
// otherwise return -1
func findChampion(n int, edges [][]int) int {
	inDegrees := make([]int, n)

	for _, edge := range edges {
		inDegrees[edge[1]]++
	}

	champion := -1
	countZeroInDegreeNode := 0
	for i, v := range inDegrees {
		if v == 0 {
			countZeroInDegreeNode++
			champion = i
		}
	}
	if countZeroInDegreeNode == 1 {
		return champion
	}
	return -1
}

// timeout
// 输入：n = 4, edges = [[0,2],[1,3],[1,2]]
// 输出：-1
// - 0 1 2 3
// 0 0 0 1 0
// 1 0 0 1 1
// 2 0 0 0 0
// 3 0 0 0 0
// func findChampion(n int, edges [][]int) int {
// 	graph := make([][]int, n)
// 	for i := range graph {
// 		graph[i] = make([]int, n)

// 	}
// 	for _, edge := range edges {
// 		graph[edge[0]][edge[1]] = 1
// 	}

// 	// to fill graph and find champion
// 	champion := -1
// 	for i := 0; i < n; i++ {
// 		// to find underdog
// 		cache := make(map[int]struct{})
// 		toFindNode(graph, i, cache)
// 		if len(cache) == n-1 {
// 			champion = i
// 		}
// 	}

// 	return champion
// }

// func toFindNode(graph [][]int, index int, cache map[int]struct{}) {
// 	if len(graph[index]) == 0 {
// 		return
// 	}
// 	for i, v := range graph[index] {
// 		if v == 1 {
// 			cache[i] = struct{}{}
// 			toFindNode(graph, i, cache)
// 		}
// 	}
// }
