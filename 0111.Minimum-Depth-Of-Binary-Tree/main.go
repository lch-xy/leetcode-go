package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 如果有一个节点是空 那就不能参与比较，不然0肯定是最小的
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	minDepth := 0
	if leftDepth == 0 {
		minDepth = rightDepth
	} else if rightDepth == 0 {
		minDepth = leftDepth
	} else {
		minDepth = min(leftDepth, rightDepth)
	}
	return minDepth + 1
}
