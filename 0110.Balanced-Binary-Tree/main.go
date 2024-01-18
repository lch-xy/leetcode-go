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
			name: "Test Case 1",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: true,
		},
		{
			name: "Test Case 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			want: false,
		},
		{
			name: "Test Case 3",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		if got := isBalanced(tt.root); got != tt.want {
			fmt.Printf("isBalanced() = %v, want %v", got, tt.want)
		}
	}
}

// 一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1 。
func isBalanced(root *TreeNode) bool {
	isBalanced, _ := helper(root)
	return isBalanced
}

func helper(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	isLeftBalanced, left := helper(root.Left)
	isRightBalanced, right := helper(root.Right)
	return !(left-right > 1 || right-left > 1) && isLeftBalanced && isRightBalanced, max(left, right) + 1
}

// 第一版代码 感觉递归次数有点多
// func isBalanced(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	left := helper(root.Left)
// 	right := helper(root.Right)
// 	return !(left-right > 1 || right-left > 1)  && isBalanced(root.Left) && isBalanced(root.Right)
// }

// func helper(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return max(helper(root.Left), helper(root.Right)) + 1
// }
