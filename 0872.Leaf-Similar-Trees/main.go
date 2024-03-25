package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return helper(root1) == helper(root2)
}

// use dfs to build a string
func helper(root *TreeNode) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return string(root.Val)
	}
	left := helper(root.Left)
	right := helper(root.Right)
	if left == "" {
		return right
	}
	if right == "" {
		return left
	}
	return left + "#" + right
}
