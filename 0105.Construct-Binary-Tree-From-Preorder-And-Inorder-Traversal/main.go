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
		name     string
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{
			name:     "Test Case 1",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
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
			name:     "Test Case 2",
			preorder: []int{1, 2, 3},
			inorder:  []int{2, 1, 3},
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
		if got := buildTree(tt.preorder, tt.inorder); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("buildTree() = %v, want %v", got, tt.want)
		}
	}
}

// preorder: [3, 9, 20, 15, 7]
// inorder: [9, 3, 15, 20, 7]
// 主要还是找到左边和右边的区间
func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func helper(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	val := preorder[preStart]

	idx := -1
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == val {
			idx = i
			break
		}
	}
	// 找到了inorder中 root节点的位置 idx，可以算出左子树的长度idx-inStart
	left := helper(preorder, preStart+1, preStart+(idx-inStart), inorder, inStart, idx-1)
	right := helper(preorder, preStart+(idx-inStart)+1, preEnd, inorder, idx+1, inEnd)
	return &TreeNode{val, left, right}
}
