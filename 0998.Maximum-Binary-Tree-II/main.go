package main

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// question：[5,2,4,null,1]
	// target ： [5,2,4,null,1,null,3]
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 1}
	res := insertIntoMaxTree(root, 3)
	println(res)
}

// 因为是前序遍历，左-中-右 的顺序，既然是加在最后面，那么肯定是在最右边
// 这题也可以先将TreeNode先序列化，再反序列化
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	v := root.Val
	if v < val {
		node := &TreeNode{Val: val}
		node.Left = root
		return node
	}
	if v > val {
		root.Right = insertIntoMaxTree(root.Right, val)
	}
	return root
}
