package main

import (
	"fmt"
	"reflect"
)

func main() {
	tests := []struct {
		edges [][]int
		want  []int
	}{
		{
			edges: [][]int{{1, 2}, {1, 3}, {2, 3}},
			want:  []int{2, 3},
		},
		{
			edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}},
			want:  []int{4, 1},
		},
		{
			edges: [][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}},
			want:  []int{2, 1},
		},
	}

	for _, tt := range tests {
		got := findRedundantDirectedConnection(tt.edges)
		if !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("findRedundantDirectedConnection(%v) = %v, want %v", tt.edges, got, tt.want)
		}
	}
}

// 并查集
func findRedundantDirectedConnection(edges [][]int) []int {
	size := len(edges)
	uf := initUnionFind(size)
	parent := make([]int, size+1)
	for i := range parent {
		parent[i] = i
	}
	conflict := -1
	cycle := -1
	for i, edge := range edges {
		node1, node2 := edge[0], edge[1]
		// 对于每一条边，如果它的终点 node2 的父节点不是 node2 自身
		// 那么就说明 node2 已经有一个父节点，所以这条边就是冲突的边，我们将 conflict 设置为这条边的索引。
		if parent[node2] != node2 {
			conflict = i
		} else {
			parent[node2] = node1
			// 这两个节点如果根节点相同，那么再连接这两个节点就会成环
			if uf.find(node1) == uf.find(node2) {
				cycle = i
			} else {
				uf.union(node1, node2)
			}
		}
	}

	if conflict < 0 {
		// 如果不存在冲突的边（conflict < 0），那么就返回导致环路的边。
		return []int{edges[cycle][0], edges[cycle][1]}
	} else {
		conflictEdge := edges[conflict]
		if cycle > 0 {
			// 因为如果发现冲突就不会执行 union 操作，不然就成环了
			// 所以返回的是冲突的边和它的父节点
			return []int{parent[conflictEdge[1]], conflictEdge[1]}
		} else {
			// 冲突没有成环，返回当前这条边
			return []int{conflictEdge[0], conflictEdge[1]}
		}
	}
}

type UnionFind struct {
	ancestor []int
}

func initUnionFind(size int) *UnionFind {
	ancestor := make([]int, size+1)
	for i := range ancestor {
		ancestor[i] = i
	}
	return &UnionFind{ancestor: ancestor}
}

func (uf *UnionFind) find(x int) int {
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
	}
	return uf.ancestor[x]
}

func (uf *UnionFind) union(x int, y int) bool {
	xRoot := uf.find(x)
	yRoot := uf.find(y)
	if xRoot == yRoot {
		return false
	}
	uf.ancestor[xRoot] = yRoot
	return true
}
