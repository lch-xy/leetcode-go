package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	tests := []struct {
		n     int
		edges [][]int
		want  []int
	}{
		{
			n:     6,
			edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			want:  []int{3, 4},
		},
		{
			n:     4,
			edges: [][]int{{1, 0}, {1, 2}, {1, 3}},
			want:  []int{1},
		},
	}

	for _, tt := range tests {
		got := findMinHeightTrees(tt.n, tt.edges)
		if !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("findMinHeightTrees(%v, %v) = %v, want %v", tt.n, tt.edges, got, tt.want)
		}
	}
}

// 使用拓扑排序来找到最小高度树的所有根节点，时间复杂度是 O(n)，因为每个节点和边都只被访问了一次。
// 每次都从图中删除当前队列中的所有节点（即当前层的所有叶子节点），并将这些节点的邻居节点的度减 1。如果邻居节点的度变为 1，那么就将邻居节点添加到队列中。
// 当图中的节点数量小于等于 2 时，队列中剩下的节点就是最小高度树的所有根节点。
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	graph := make([][]int, n)
	queue := []int{}

	for _, v := range edges {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	for i := 0; i < n; i++ {
		// 如果长度为1，说明是叶子节点
		if len(graph[i]) == 1 {
			queue = append(queue, i)
		}
	}

	// 因为在一个无环图（或树）中，最小高度树的根节点一定是图的中心节点。而图的中心节点最多有两个。
	for n > 2 {
		size := len(queue)
		n -= size
		// 将叶子结点全部摘除
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 反向也要摘除，而且如果它的邻居节点的入度减为1时，也需要加入到队列里面
			for _, v := range graph[cur] {
				graph[v] = remove(graph[v], cur)
				if len(graph[v]) == 1 {
					queue = append(queue, v)
				}
			}
		}
	}
	// 最后队列里剩下的就是根节点
	return queue
}

func remove(slice []int, val int) []int {
	for i, v := range slice {
		if v == val {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// 先构建成一个有向图，然后dfs遍历每个节点，计算机出每个节点最小的高度
// 但是无法ac，会超时
func findMinHeightTreesDfs(n int, edges [][]int) []int {
	graph := make([][]bool, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		graph[i] = make([]bool, n)
	}

	// 建立有向图
	for i := 0; i < len(edges); i++ {
		graph[edges[i][0]][edges[i][1]] = true
		graph[edges[i][1]][edges[i][0]] = true
	}
	res := []int{}
	minHeight := math.MaxInt
	for i := 0; i < n; i++ {
		height := dfs(graph, visited, i)
		if height < minHeight {
			minHeight = height
			res = []int{i}
		} else if height == minHeight {
			res = append(res, i)
		}
	}
	return res
}

func dfs(graph [][]bool, visited []bool, i int) int {
	if visited[i] {
		return 0
	}
	res := 0

	visited[i] = true
	for j := 0; j < len(graph[i]); j++ {
		if graph[i][j] {
			res = max(res, dfs(graph, visited, j))
		}
	}
	visited[i] = false
	return res + 1
}
