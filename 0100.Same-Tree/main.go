package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 创建一些测试用例
	testCases := []struct {
		p        *TreeNode
		q        *TreeNode
		expected bool
	}{
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: true,
		},
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			q: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
	}

	// 运行测试用例
	for _, tc := range testCases {
		result := isSameTree(tc.p, tc.q)
		if result != tc.expected {
			fmt.Printf("isSameTree(%v, %v) = %v; expected %v", tc.p, tc.q, result, tc.expected)
		}
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
