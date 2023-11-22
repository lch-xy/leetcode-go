package main

import "fmt"

func main() {
	graph := [][]int{
		[]int{1, 2},
		[]int{3},
		[]int{3},
		[]int{},
	}
	graph = [][]int{
		[]int{4, 3, 1},
		[]int{3, 2, 4},
		[]int{3},
		[]int{4},
		[]int{},
	}
	fmt.Print(allPathsSourceTarget(graph))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	path := []int{}
	traverse(graph, 0, &path, &res)
	return res
}

func traverse(graph [][]int, s int, path *[]int, res *[][]int) {
	*path = append(*path, s)

	n := len(graph)
	if s == n-1 {
		// 达到目标节点，记录到res中
		*res = append(*res, append([]int{}, *path...))
	}

	// 遍历graph的邻接表
	for _, v := range graph[s] {
		traverse(graph, v, path, res)
	}

	*path = (*path)[:len(*path)-1]
}
