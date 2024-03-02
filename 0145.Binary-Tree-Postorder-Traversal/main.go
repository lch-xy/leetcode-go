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
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			want: []int{3, 2, 4, 1},
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
			want: []int{2, 1},
		},
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			want: []int{2, 1},
		},
	}

	for i, tc := range testCases {
		got := postorderTraversalIteration(tc.root)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("Test case %d: expected %v, got %v", i, tc.want, got)
		}
	}
}

func postorderTraversalIteration(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	stackOne := []*TreeNode{root}
	stackTwo := []*TreeNode{}
	for len(stackOne) > 0 {
		curNode := stackOne[len(stackOne)-1]
		stackOne = stackOne[:len(stackOne)-1]
		stackTwo = append(stackTwo, curNode)
		if curNode.Left != nil {
			stackOne = append(stackOne, curNode.Left)
		}
		if curNode.Right != nil {
			stackOne = append(stackOne, curNode.Right)
		}
	}

	for len(stackTwo) > 0 {
		curNode := stackTwo[len(stackTwo)-1]
		stackTwo = stackTwo[:len(stackTwo)-1]
		res = append(res, curNode.Val)
	}
	return res
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	helper(root, &res)
	return res
}
func helper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, res)
	helper(root.Right, res)
	*res = append(*res, root.Val)
}
