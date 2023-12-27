package main

import "math"

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// root = [3,1,4,null,null,2]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 2}
	recoverTree(root)

	println("=======================================>")

	// root = [1,3,null,null,2]
	//root := &TreeNode{Val: 1}
	//root.Left = &TreeNode{Val: 3}
	//root.Left.Right = &TreeNode{Val: 2}
	//recoverTree(root)

	println("=======================================>")

	// root = [2,3,1]
	//root := &TreeNode{Val: 2}
	//root.Left = &TreeNode{Val: 3}
	//root.Right = &TreeNode{Val: 1}
	//recoverTree(root)

	println("=======================================>")

}

// 使用Morris遍历算法 需要额外的空间复杂度
func recoverTree(root *TreeNode) {
	var first, second, pre *TreeNode
	firstTime := true
	pre = &TreeNode{Val: math.MinInt}
	for root != nil {
		if root.Left != nil {
			temp := root.Left
			// 找到左子树的最大的节点
			for temp.Right != nil && temp.Right != root {
				temp = temp.Right
			}
			// 如果为空说明还没有建桥
			// 那么把右节点指向root节点
			if temp.Right == nil {
				// 建桥，方便后面回溯
				temp.Right = root
				// 建桥后，向左遍历，建桥是用temp节点来找的，这里用root.Left即可
				root = root.Left
			} else {
				// 如果不为空就说明之前已经建过桥 这里要拆桥
				temp.Right = nil
				// 这里放处理的逻辑
				// 如果不是有序的，而且是第一次，就记录到first里
				if root.Val < pre.Val && firstTime {
					first = pre
					firstTime = false
				}
				if root.Val < pre.Val && !firstTime {
					second = root
				}
				// 把当前节点记录成pre节点
				pre = root
				// 我们左边遍历完后，其实是会返回到根节点的
				// 因为最右边的节点又跟当前的root节点连接，所以会返回当前的root节点
				// root节点下次再进行左节点的遍历的时候，发现已经建过桥了，这时候就会去拆桥
				// 拆完桥然后跟节点往右走，遍历右子树
				root = root.Right
			}
		} else {
			// 这里放我们的处理逻辑
			if root.Val < pre.Val && firstTime {
				first = pre
				firstTime = false
			}
			if root.Val < pre.Val && !firstTime {
				second = root
			}
			pre = root
			root = root.Right
		}
	}
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}
