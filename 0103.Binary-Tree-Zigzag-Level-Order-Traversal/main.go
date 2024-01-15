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
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{
			name: "Test Case 1",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name: "Test Case 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			want: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}

	for _, tt := range tests {
		if got := zigzagLevelOrder(tt.root); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("zigzagLevelOrder() = %v, want %v", got, tt.want)
		}
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		level++
		curRes := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			var curNode *TreeNode
			if level%2 == 1 {
				curNode = queue[0]
				queue = queue[1:]
			} else {
				curNode = queue[len(queue)-1]
				queue = queue[:len(queue)-1]
			}

			curRes = append(curRes, curNode.Val)
			if level%2 == 1 {
				if curNode.Left != nil {
					queue = append(queue, curNode.Left)
				}
				if curNode.Right != nil {
					queue = append(queue, curNode.Right)
				}
			} else {
				if curNode.Right != nil {
					queue = append([]*TreeNode{curNode.Right}, queue...)
				}
				if curNode.Left != nil {
					queue = append([]*TreeNode{curNode.Left}, queue...)
				}
			}

		}
		res = append(res, curRes)
	}
	return res
}
