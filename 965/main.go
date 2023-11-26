package main

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	val := root.Val
	return helper(root, val)
}

func helper(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val != val {
		return false
	}
	return helper(root.Left, val) && helper(root.Right, val)
}
