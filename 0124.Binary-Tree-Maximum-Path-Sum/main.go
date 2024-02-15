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
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: &TreeNode{
				Val: -10,
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
			want: 42,
		},
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
			want: 6,
		},
		{
			root: &TreeNode{
				Val: -3,
			},
			want: -3,
		},
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			want: 48,
		},
	}

	for i, tc := range testCases {
		got := maxPathSum(tc.root)
		if got != tc.want {
			fmt.Printf("Test case %d: expected %d, got %d", i, tc.want, got)
		}
	}
}

// 解题思路：递归的计算以当前节点的最大值
// 还有一种思路就是转化为无向图，然后遍历图的每个节点计算最大值
func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	calculate(root, &res)
	return res
}

// 计算二叉树中的最大增益，如果节点为空，返回 0
func calculate(curNode *TreeNode, res *int) int {
	if curNode == nil {
		return 0
	}
	// 如果小于值0 就直接忽略
	left := max(0, calculate(curNode.Left, res))
	right := max(0, calculate(curNode.Right, res))
	// 后续遍历的时候
	*res = max(*res, left+right+curNode.Val)
	return curNode.Val + max(left, right)
}
