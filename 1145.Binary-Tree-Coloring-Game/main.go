package main

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	//root.Right.Left = &TreeNode{Val: 6}
	//root.Right.Right = &TreeNode{Val: 7}
	//root.Left.Left.Left = &TreeNode{Val: 8}
	//root.Left.Left.Right = &TreeNode{Val: 9}
	//root.Left.Right.Left = &TreeNode{Val: 10}
	//root.Left.Right.Right = &TreeNode{Val: 11}
	println(btreeGameWinningMove(root, 5, 1))
}

// 先找到x所在的节点，计算左右节点的数量
// 获胜节点只可能出现在父节点或者两个子节点上，因为向上或者向下的路只有一条
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	node := getNode(root, x)
	if node == nil {
		return false
	}

	left := helper(node.Left)
	right := helper(node.Right)

	// 计算获胜节点数的临界值
	half := (n - 1) / 2
	// 可以获胜的方案
	//  1、左节点 > 临界值
	//  2、有节点 > 临界值
	//  3、当前节点所包含的节点个数 <= 临界值
	if half < left || half < right || (left+right+1) <= half {
		return true
	}
	return false
}

// 获取节点位置
func getNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := getNode(root.Left, val)
	if left != nil {
		return left
	}
	rigth := getNode(root.Right, val)
	if rigth != nil {
		return rigth
	}
	return nil
}

// 节点数
func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left)
	right := helper(root.Right)

	return left + right + 1
}
