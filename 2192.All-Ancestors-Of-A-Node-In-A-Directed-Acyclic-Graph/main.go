package main

import (
	"fmt"
	"sort"
)

func main() {

	// n := 8
	// edges := [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}
	n := 6
	edges := [][]int{{5, 1}, {2, 3}, {5, 3}, {0, 2}, {3, 1}, {5, 2}, {4, 0}}

	answer := getAncestors(n, edges)
	fmt.Println(answer)
	answer = getAncestorsTwo(n, edges)
	fmt.Println(answer)
}

// use adjacency matrix
func getAncestors(n int, edges [][]int) [][]int {
	graph := make([][]bool, n)
	for i := range graph {
		graph[i] = make([]bool, n)
	}
	for _, edge := range edges {
		graph[edge[1]][edge[0]] = true
	}

	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		queue := []int{i}
		curList := make([]int, 0)
		visited := make([]bool, n)
		for len(queue) != 0 {
			curNode := queue[0]
			queue = queue[1:]
			for i, v := range graph[curNode] {
				if v {
					if !visited[i] {
						visited[i] = true
						queue = append(queue, i)
						curList = append(curList, i)
					}
				}
			}
		}
		sort.Ints(curList)
		res = append(res, curList)
	}
	return res
}

// use adjacency list
func getAncestorsTwo(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[to] = append(graph[to], from)
	}

	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		queue := []int{i}
		curList := make([]int, 0)
		visited := make([]bool, n)
		for len(queue) != 0 {
			curNode := queue[0]
			queue = queue[1:]
			for _, v := range graph[curNode] {
				if !visited[v] {
					visited[v] = true
					queue = append(queue, v)
					curList = append(curList, v)
				}
			}
		}
		sort.Ints(curList)
		res = append(res, curList)
	}
	return res
}
