package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	res, pre := math.MaxInt32, -1
	helper(root, &pre, &res)
	return res
}

func helper(root *TreeNode, pre *int, res *int) {
	if root == nil {
		return
	}

	helper(root.Left, pre, res)
	if *pre != -1 {
		*res = min(*res, root.Val-*pre)
	}
	*pre = root.Val
	helper(root.Right, pre, res)
}
