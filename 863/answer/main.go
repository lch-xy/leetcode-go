package main

import (
	"container/list"
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//root = [3,5,1,6,2,0,8,null,null,7,4],
	// 根据上面数组创建一颗TreeNode
	root := &TreeNode{3, nil, nil}
	root.Left = &TreeNode{5, nil, nil}
	root.Right = &TreeNode{1, nil, nil}
	root.Left.Left = &TreeNode{6, nil, nil}
	root.Left.Right = &TreeNode{2, nil, nil}
	root.Right.Left = &TreeNode{0, nil, nil}
	root.Right.Right = &TreeNode{8, nil, nil}
	root.Left.Right.Left = &TreeNode{7, nil, nil}
	root.Left.Right.Right = &TreeNode{4, nil, nil}
	fmt.Print(distanceK(root, root.Left.Right, 3))

}

func distanceK(root, target *TreeNode, k int) (ans []int) {
	if root == nil {
		return []int{}
	}

	graph := map[*TreeNode][]*TreeNode{}
	var fillGraph func(node *TreeNode)
	// 遍历构建相互映射关系
	fillGraph = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			graph[node] = append(graph[node], node.Left)
			graph[node.Left] = append(graph[node.Left], node)
			fillGraph(node.Left)
		}
		if node.Right != nil {
			graph[node] = append(graph[node], node.Right)
			graph[node.Right] = append(graph[node.Right], node)
			fillGraph(node.Right)

		}
	}
	fillGraph(root)

	res := []int{}
	queue := list.New()
	queue.PushBack(target)
	visited := make(map[int]bool)
	visited[target.Val] = true

	// bfs遍历，走k步
	for k >= 0 && queue.Len() != 0 {
		if k == 0 {
			for queue.Len() > 0 {
				node := queue.Front().Value.(*TreeNode)
				queue.Remove(queue.Front())
				res = append(res, node.Val)
			}
			return res
		}
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())

			for _, neighbor := range graph[node] {
				if !visited[neighbor.Val] {
					visited[neighbor.Val] = true
					queue.PushBack(neighbor)
				}
			}
		}
		k--
	}
	return res
}
