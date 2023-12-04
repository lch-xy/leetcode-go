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
	root.Right.Left = &TreeNode{Val: 6}
	//root.Right.Right = &TreeNode{Val: 7}
	//root.Left.Left.Left = &TreeNode{Val: 8}
	//root.Left.Left.Right = &TreeNode{Val: 9}
	//root.Left.Right.Left = &TreeNode{Val: 10}
	//root.Left.Right.Right = &TreeNode{Val: 11}
	println(maxProduct(root))
}

// 解法1：
// 先递归计算总数，再进行第二次递归计算乘积
func maxProduct(root *TreeNode) int {
	const M = 1e9 + 7
	res := 0
	point_res := &res
	cnt := helper(root, point_res, 0)
	helper(root, point_res, cnt)
	return int(*point_res % M)
}

func helper(root *TreeNode, res *int, total int) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left, res, total)
	right := helper(root.Right, res, total)

	curCnt := left + right + root.Val
	// 计算sum的时候 这一步是不会进来的，所以可以共用一个递归函数
	if curCnt*(total-curCnt) > *res {
		*res = curCnt * (total - curCnt)
	}
	return curCnt
}
