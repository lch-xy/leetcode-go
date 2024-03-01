package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	testCases := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			want: []int{1, 2, 3},
		},
		{
			root: nil,
			want: []int{},
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			want: []int{1},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			want: []int{1, 2},
		},
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			want: []int{1, 2},
		},
	}

	for i, tc := range testCases {
		got := preorderTraversalIteration(tc.root)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("Test case %d: expected %v, got %v", i, tc.want, got)
		}
	}
}


func preorderTraversalIteration(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curNode.Val)
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
	}
	return res
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	helper(root, &res)
	return res
}

func helper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	helper(root.Left, res)
	helper(root.Right, res)
}
