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
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: [][]int{{1}, {2, 3}, {4, 5}},
		},
		{
			name: "Test Case 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			want: [][]int{{1}, {2, 3}, {4, 5}},
		},
	}

	for _, tt := range tests {
		if got := levelOrder(tt.root); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("Function() = %v, want %v", got, tt.want)
		}
	}
}

// 直接用队列进行层序遍历即可
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curRes := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			curNode := queue[0]
			curRes = append(curRes, curNode.Val)
			queue = queue[1:]
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		res = append(res, curRes)
	}
	return res
}
