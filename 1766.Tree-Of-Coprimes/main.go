package main

import (
	"fmt"
	"reflect"
)

func main() {
	testCases := []struct {
		nums     []int
		edges    [][]int
		expected []int
	}{
		{
			nums:     []int{2, 3, 3, 2},
			edges:    [][]int{{0, 1}, {1, 2}, {1, 3}},
			expected: []int{-1, 0, 0, 1},
		},
	}

	for i, tc := range testCases {
		result := getCoprimes(tc.nums, tc.edges)
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test case %d passed\n", i+1)
		} else {
			fmt.Printf("Test case %d failed: expected %v, got %v\n", i+1, tc.expected, result)
			return
		}
	}
}

// how to make algorithm more fast?
// we can caculate the result of gcd from 0 to 50 in advanced
func getCoprimes(nums []int, edges [][]int) []int {
	length := len(nums)
	graph := make([][]int, length)
	for i := range graph {
		graph[i] = make([]int, 0)
	}
	for _, edge := range edges {
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}

	// ancestors's struct: {value,{index,deap}}
	// to store node's ancestors
	ancestors := make(map[int][][2]int)
	for _, num := range nums {
		ancestors[num] = make([][2]int, 0)
	}

	res := make([]int, length)

	dfs(ancestors, graph, nums, &res, 0, -1, 0)

	return res
}

func dfs(ancestors map[int][][2]int, graph [][]int, nums []int, res *[]int, node, parent, depth int) {
	curVal := nums[node]
	ancestor := -1
	maxDepth := -1
	// to find the closest one
	for key, value := range ancestors {
		if gcd(key, curVal) == 1 {
			for _, tumple := range value {
				curIndex := tumple[0]
				curDepth := tumple[1]
				if curDepth > maxDepth {
					maxDepth = curDepth
					ancestor = curIndex
				}
			}
		}
	}
	(*res)[node] = ancestor

	ancestors[curVal] = append(ancestors[curVal], [2]int{node, depth})

	for _, child := range graph[node] {
		// avoid dead loop
		if child != parent {
			dfs(ancestors, graph, nums, res, child, node, depth+1)
		}
	}
	ancestors[curVal] = ancestors[curVal][:len(ancestors[curVal])-1]
}

func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

// error solution , can accept some easy example
// this solution can't find the closest node , because we never store depth
// func getCoprimes(nums []int, edges [][]int) []int {
// 	length := len(nums)
// 	graph := make([][]int, length)
// 	for i := range graph {
// 		graph[i] = make([]int, 0)
// 	}
// 	for _, edge := range edges {
// 		graph[edge[1]] = append(graph[edge[1]], edge[0])
// 	}
// 	res := make([]int, length)
// 	for i := 0; i < length; i++ {
// 		curNumber := nums[i]
// 		if len(graph[i]) == 0 {
// 			res[i] = -1
// 			continue
// 		}
// 		res[i] = findIndex(graph, nums, curNumber, i)
// 	}
// 	return res
// }

// func gcd(a, b int) int {
// 	for b != 0 {
// 		temp := b
// 		b = a % b
// 		a = temp
// 	}
// 	return a
// }

// func findIndex(graph [][]int, nums []int, startNumber, curIndex int) int {
// 	for _, v := range graph[curIndex] {
// 		curNumber := nums[v]
// 		if gcd(startNumber, curNumber) == 1 {
// 			return v
// 		}
// 		return findIndex(graph, nums, startNumber, v)
// 	}
// 	return -1
// }
