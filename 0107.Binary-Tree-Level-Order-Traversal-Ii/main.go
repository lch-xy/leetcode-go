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
			want: [][]int{
				{15, 7},
				{9, 20},
				{3},
			},
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
				},
			},
			want: [][]int{
				{2, 3},
				{1},
			},
		},
	}

	for _, tt := range tests {
		if got := levelOrderBottom(tt.root); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("levelOrderBottom() = %v, want %v", got, tt.want)
		}
	}
}

// 和102题一模一样，只需要添加的时候往前添加即可
func levelOrderBottom(root *TreeNode) [][]int {
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
		res = append([][]int{curRes}, res...)
	}
	return res
}
