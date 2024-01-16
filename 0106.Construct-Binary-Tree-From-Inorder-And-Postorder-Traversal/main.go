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
		name      string
		inorder   []int
		postorder []int
		want      *TreeNode
	}{
		{
			name:      "Test Case 1",
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			want: &TreeNode{
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
		},
		{
			name:      "Test Case 2",
			inorder:   []int{2, 1, 3},
			postorder: []int{2, 3, 1},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
	}

	for _, tt := range tests {
		if got := buildTree(tt.inorder, tt.postorder); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("buildTree() = %v, want %v", got, tt.want)
		}
	}
}

// [9, 3, 15, 20, 7]
// [9, 15, 7, 20, 3]
// 主要还是找到左边和右边的区间
func buildTree(inorder []int, postorder []int) *TreeNode {
	return helper(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func helper(inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if inStart > inEnd || postStart > postEnd {
		return nil
	}
	val := postorder[postEnd]

	idx := -1
	for i := inStart; i <= inEnd; i++ {
		if val == inorder[i] {
			idx = i
			break
		}
	}
	// 左边长 (idx - 1) - inStart + 1 = idx - inStart
	// 右边长 inEnd - (idx + 1) + 1 = inEnd - idx
	// postStart+(idx-inStart) = postEnd-(inEnd-idx)
	left := helper(inorder, inStart, idx-1, postorder, postStart, postStart+(idx-inStart)-1)
	Right := helper(inorder, idx+1, inEnd, postorder, postStart+(idx-inStart), postEnd-1)

	return &TreeNode{val, left, Right}
}
