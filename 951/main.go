package main

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root1 := &TreeNode{3, nil, nil}
	root1.Left = &TreeNode{5, nil, nil}
	root1.Right = &TreeNode{1, nil, nil}
	root1.Left.Left = &TreeNode{6, nil, nil}
	root1.Left.Right = &TreeNode{2, nil, nil}

	root2 := &TreeNode{3, nil, nil}
	root2.Left = &TreeNode{1, nil, nil}
	root2.Right = &TreeNode{5, nil, nil}
	root2.Right.Left = &TreeNode{2, nil, nil}
	root2.Right.Right = &TreeNode{6, nil, nil}

	println(flipEquiv(root1, root2))
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	// 都为空说明已经遍历到底了,此时应该返回 true 的
	if root1 == nil && root2 == nil {
		return true
	}
	// 其中一个为空或者当前两个根节点都不相等，直接返回false
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	// 不需要交换节点，直接进行递归
	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}
