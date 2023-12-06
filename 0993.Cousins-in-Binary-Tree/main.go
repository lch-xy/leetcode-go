package main

import "container/list"

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//[1,2,3,null,4,null,5]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	println(isCousins(root, 5, 4))
}

// 使用二叉树的层序遍历来做，定义两个flag代表匹配到了x和y
// 我们直接在队列取出来的时候进行判断是否等于x or y即可
// 同一个for循环里代表同一层
// 既然存在于队列中 那么肯定已经通过了父节点对左右子树的判断了
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		isX := false
		isY := false
		for i := 0; i < size; i++ {
			node := queue.Front().Value.(*TreeNode)
			if node.Val == x {
				isX = true
			}
			if node.Val == y {
				isY = true
			}
			if node.Left != nil && node.Right != nil {
				left := node.Left.Val
				right := node.Right.Val
				if (left == x && right == y) || (left == y && right == x) {
					return false
				}
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			queue.Remove(queue.Front())
		}

		if isX && isY {
			return true
		}
	}
	return false
}
