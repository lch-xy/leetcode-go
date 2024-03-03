package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	vals  []int
	index int
}

// 解题思路：先通过二叉树的前序遍历将所有的节点的值放到vals数组中
// index为当前要取出的值的索引下标
func Constructor(root *TreeNode) BSTIterator {
	res := []int{}
	index := math.MaxInt32
	helper(root, &res, &index)
	return BSTIterator{res, 0}
}

func helper(root *TreeNode, res *[]int, index *int) {
	if root == nil {
		return
	}
	if root.Val < *index {
		index = &root.Val
	}
	helper(root.Left, res, index)
	*res = append(*res, root.Val)
	helper(root.Right, res, index)
}

func (this *BSTIterator) Next() int {
	if len(this.vals) > this.index {
		cur := this.vals[this.index]
		this.index++
		return cur
	}
	return -1
}

func (this *BSTIterator) HasNext() bool {
	return len(this.vals) > this.index
}
