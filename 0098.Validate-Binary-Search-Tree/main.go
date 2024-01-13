package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "Test Case 1: Valid Binary Search Tree",
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: true,
		},
		{
			name: "Test Case 2: Invalid Binary Search Tree",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			want: false,
		}, {
			name: "Test Case 3: Invalid Binary Search Tree with same values",
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		if got := isValidBST(tt.root); got != tt.want {
			fmt.Printf("isValidBST() = %v, want %v", got, tt.want)
		}

	}
}

// 根据二叉搜索树的定义求解即可
func isValidBST(root *TreeNode) bool {
	return helper(root.Left, math.MinInt, root.Val) && helper(root.Right, root.Val, math.MaxInt)
}

func helper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}
