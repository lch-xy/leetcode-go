package main

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 0}
	println(distributeCoins(root))
}

func distributeCoins(root *TreeNode) int {
	res := 0
	rpt := &res
	helper(root, rpt)
	return res
}

// 返回当前节点需要获得或者交出银币的个数
// 我们只需要关注当前节点，剩下的递归函数会帮我们完成
func helper(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, res)
	right := helper(root.Right, res)

	// 核心代码在后续遍历的位置，我们从叶子节点往上开始遍历

	// 我们在计算次数的时候，需要用上绝对值
	// 无论是差了多少硬币或者多出来多少硬币，都是需要移动的
	// 多出来的是要往父节点移动的，少的也是要父节点给我们的
	*res += abs(left) + abs(right) + root.Val - 1
	// 每个节点可以给出的硬币个数是左右子节点分别可以给出的个数加上当前节点值并减1
	// 负数表示还差多少硬币
	// 整数表示多了多少隐蔽
	return left + right + root.Val - 1
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
