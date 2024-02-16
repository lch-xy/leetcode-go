package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: 25,
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			want: 1026,
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 9,
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			want: 531,
		},
	}

	for i, tc := range testCases {
		got := sumNumbers(tc.root)
		if got != tc.want {
			fmt.Printf("Test case %d: expected %d, got %d", i, tc.want, got)
		}
	}
}

func sumNumbers(root *TreeNode) int {
	res := 0
	caculate(root, &res, 0)
	return res
}

func caculate(root *TreeNode, res *int, cur int) {
	cur = 10*cur + root.Val
	// 如果左右节点都为空 说明是叶子节点，纳入计算
	if root.Left == nil && root.Right == nil {
		*res += cur
		return
	}
	if root.Left != nil {
		caculate(root.Left, res, cur)
	}
	if root.Right != nil {
		caculate(root.Right, res, cur)
	}
}
