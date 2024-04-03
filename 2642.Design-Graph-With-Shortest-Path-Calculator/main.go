package main

import (
	"fmt"
	"math"
)

func main() {
	g := Constructor(5, [][]int{{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3}})
	fmt.Println(g.ShortestPath(3, 2))
	fmt.Println(g.ShortestPath(0, 3))

	g.AddEdge([]int{1, 3, 4})

	fmt.Println(g.ShortestPath(0, 3))
}

type Graph struct {
	cnt             int
	adjacencyMatrix [][]int
}

func Constructor(n int, edges [][]int) Graph {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = math.MaxInt32
		}
	}

	for _, v := range edges {
		matrix[v[0]][v[1]] = v[2]
	}

	graph := Graph{
		cnt:             n,
		adjacencyMatrix: matrix,
	}

	return graph
}

func (this *Graph) AddEdge(edge []int) {
	this.adjacencyMatrix[edge[0]][edge[1]] = edge[2]
}

// use dijkstra's algorithm to find the shortest path
func (this *Graph) ShortestPath(node1 int, node2 int) int {
	distance := make([]int, this.cnt)
	visited := make([]bool, this.cnt)

	for i := range distance {
		distance[i] = math.MaxInt32
	}
	distance[node1] = 0

	for k := 0; k < this.cnt; k++ {
		closed := -1

		for i := 0; i < this.cnt; i++ {
			if !visited[i] && (closed == -1 || (distance[closed] > distance[i])) {
				closed = i
			}
		}
		if closed == -1 {
			break
		}
		visited[closed] = true
		for i := 0; i < this.cnt; i++ {
			if distance[closed] != math.MaxInt32 && this.adjacencyMatrix[closed][i] != math.MaxInt32 {
				distance[i] = min(distance[i], distance[closed]+this.adjacencyMatrix[closed][i])
			}
		}
	}

	if distance[node2] == math.MaxInt32 {
		return -1
	}
	return distance[node2]
}
