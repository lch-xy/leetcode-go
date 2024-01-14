package main

import "fmt"

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
			name: "Test Case 1: Symmetric Binary Tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: true,
		},
		{
			name: "Test Case 2: Non-Symmetric Binary Tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		if got := isSymmetric(tt.root); got != tt.want {
			fmt.Printf("isSymmetric() = %v, want %v", got, tt.want)
		}
	}
}

func isSymmetric(root *TreeNode) bool {
	return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) || left.Val != right.Val {
		return false
	}
	// 对称是 左-右 右-左
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
