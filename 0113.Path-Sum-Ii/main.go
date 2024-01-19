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
		root      *TreeNode
		targetSum int
		res       *[][]int
		cur       []int
		want      *[][]int
	}{
		{
			name: "Test Case 1",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: 22,
			res:       &[][]int{},
			cur:       []int{},
			want:      &[][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
	}

	for _, tt := range tests {
		helper(tt.root, tt.targetSum, tt.res, tt.cur)
		if !reflect.DeepEqual(tt.res, tt.want) {
			fmt.Printf("helper() = %v, want %v", tt.res, tt.want)
		}
	}
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	helper(root, targetSum, &res, []int{})
	return res
}

func helper(root *TreeNode, targetSum int, res *[][]int, cur []int) {
	if root == nil {
		return
	}

	cur = append(cur, root.Val)
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		// 在 Go 语言中，所有的函数参数都是值传递
		// 对于引用类型（如切片、映射和通道等），虽然它们的指针被复制，但是它们指向的底层数据结构是共享的。
		// 这里要重新创建一个新的切片
		curCopy := make([]int, len(cur))
		copy(curCopy, cur)
		*res = append(*res, curCopy)
		return
	}

	helper(root.Left, targetSum-root.Val, res, cur)
	helper(root.Right, targetSum-root.Val, res, cur)
}
