package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func maxAncestorDiff(root *TreeNode) int {
	res := 0
	helper(root.Left, root.Val, root.Val, &res)
	helper(root.Right, root.Val, root.Val, &res)
	return res
}

// record the maximum value and minimum value and caculate the max diff
// why we should record the minimum value?
// because the maximum value may be in the leaf node
func helper(curNode *TreeNode, minVal, maxVal int, res *int) {
	if curNode == nil {
		return
	}
	*res = max(*res, abs(minVal, curNode.Val), abs(maxVal, curNode.Val), abs(maxVal, minVal))
	helper(curNode.Left, min(minVal, curNode.Val), max(maxVal, curNode.Val), res)
	helper(curNode.Right, min(minVal, curNode.Val), max(maxVal, curNode.Val), res)
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
