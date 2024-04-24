package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	type TestCase struct {
		root   *TreeNode
		start  int
		output int
	}
	testCases := []TestCase{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 8,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			start:  2,
			output: 3,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			start:  2,
			output: 1,
		},
	}

	for i, tc := range testCases {
		result := amountOfTime(tc.root, tc.start)
		if result == tc.output {
			fmt.Printf("Test case %d passed\n", i+1)
		} else {
			fmt.Printf("Test case %d failed: expected %d, got %d\n", i+1, tc.output, result)
			return
		}
	}
}

// use dfs to build graph
func amountOfTime(root *TreeNode, start int) int {
	graph := make(map[int][]int)
	constructGraph(-1, root, &graph)

	// use bfs to caculate total time
	queue := []int{start}
	visited := make(map[int]struct{})
	res := -1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			visited[cur] = struct{}{}
			for _, value := range graph[cur] {
				if _, ok := visited[value]; !ok {
					queue = append(queue, value)
				}
			}
		}
		res++
	}
	return res
}

func constructGraph(parentVal int, child *TreeNode, graph *map[int][]int) {
	if child == nil {
		return
	}
	if parentVal != -1 {
		// build graph
		(*graph)[parentVal] = append((*graph)[parentVal], child.Val)
		(*graph)[child.Val] = append((*graph)[child.Val], parentVal)
	}
	constructGraph(child.Val, child.Left, graph)
	constructGraph(child.Val, child.Right, graph)

}
