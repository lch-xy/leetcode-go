package main

import "fmt"

func main() {
	fmt.Println(countPairsOfConnectableServers([][]int{{0, 1, 1}, {1, 2, 5}, {2, 3, 13}, {3, 4, 9}, {4, 5, 2}}, 1))
	fmt.Println(countPairsOfConnectableServers([][]int{{0, 6, 3}, {6, 5, 3}, {0, 3, 1}, {3, 2, 7}, {3, 1, 6}, {3, 4, 2}}, 3))
	fmt.Println(countPairsOfConnectableServers([][]int{{1, 0, 2}, {2, 1, 4}, {3, 2, 4}, {4, 0, 3}, {5, 1, 4}, {6, 2, 2}, {7, 6, 4}, {8, 1, 2}, {9, 8, 3}}, 1))
}

type Node struct {
	to       int
	weighted int
}

var _signalSpeed int
var graph [][]Node

// this is a question that caculate how many pair of node will be connected via i node
// use dfs to find all node that meet condition
// how to caculate all result?
// for example: if node i have three child , count of nodes that satisfy the condition are [3 ,4 ,5] respectively
// when we search the first child and find 3 nodes , because we not search another part , so the result is 0
// when we search the second child and find 4 nodes , we can find 4 * 3 routes
// when we search the third child and find 5 nodes , we should add (4+3)*5 routes
func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	_signalSpeed = signalSpeed
	n := len(edges)
	graph = make([][]Node, n+1)
	for i := 0; i < n; i++ {
		_from, _to, _weighted := edges[i][0], edges[i][1], edges[i][2]
		graph[_from] = append(graph[_from], Node{
			to:       _to,
			weighted: _weighted,
		})
		graph[_to] = append(graph[_to], Node{
			to:       _from,
			weighted: _weighted,
		})
	}
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		count := 0
		for _, node := range graph[i] {
			to, weighted := node.to, node.weighted
			researched := dfs(to, i, weighted)
			res[i] += count * researched
			count += researched
		}
	}
	return res
}

func dfs(child int, parent int, weighted int) int {
	res := 0
	if weighted%_signalSpeed == 0 {
		res = 1
	}
	for _, node := range graph[child] {
		_to, _weighted := node.to, node.weighted
		if _to != parent {
			res += dfs(_to, child, weighted+_weighted)
		}
	}
	return res
}
