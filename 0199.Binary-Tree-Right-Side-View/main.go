package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	// 测试rightSideView函数
	expected := []int{1, 3, 4}
	if res := rightSideView(root); !reflect.DeepEqual(res, expected) {
		fmt.Printf("Expected %v, but got %v", expected, res)
	}

	// 创建一个空的二叉树
	root = nil
	expected = []int{}
	if res := rightSideView(root); !reflect.DeepEqual(res, expected) {
		fmt.Printf("Expected %v, but got %v", expected, res)
	}

	// 创建一个只有一个节点的二叉树
	root = &TreeNode{Val: 1}
	expected = []int{1}
	if res := rightSideView(root); !reflect.DeepEqual(res, expected) {
		fmt.Printf("Expected %v, but got %v", expected, res)
	}
}

// 解题思路：使用二叉树的层序遍历即可，我们倒着遍历，这样直接取第一个值放入res数组中即可
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		rightNode := queue[0]
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
		}
		res = append(res, rightNode.Val)
	}

	return res
}
