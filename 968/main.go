package main

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// [0,null,0,null,0,null,0] 建个二叉树
	root := &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 0}
	root.Right.Right.Right = &TreeNode{Val: 0}
	println(minCameraCover(root))
}

func minCameraCover(root *TreeNode) int {
	var res = 0
	ptr := &res

	if helper(root, ptr) < 1 {
		return 1 + res
	}
	return res
}

// 从叶子节点向上遍历，所以用后序遍历入手
// 叶子节点：只能监控到自己的和父节点，共2个。
// 父节点：父节点最适合放监控，因为可以同时监控 2个子节点 + 自己 + 父节点 = 4个。

// return 0; 表示需要放置摄像头的节点
// return 1; 表示已经放置了摄像头的节点
// return 2; 代表被覆盖但是没有摄像头的节点
func helper(root *TreeNode, res *int) int {
	if root == nil {
		return 2
	}

	left := helper(root.Left, res)
	right := helper(root.Right, res)

	// 如果左右节点都是0 代表是叶子节点
	// 那么在他们的父节点上放监控最有性价比
	if left == 0 || right == 0 {
		*res++
		return 1
	}

	// 如果他们的子节点其中有一个被放置了监控
	// 说明自己可以不用管了，可以和空节点归纳到一起，返回2代表不用管了
	// 而他们的父节点，就可以理解成新的叶子节点，
	if left == 1 || right == 1 {
		return 2
	}
	// 如果是叶子节点，那么他的子节点为空都会返回2，他就会走到这返回0，代表是叶子节点
	return 0
}
