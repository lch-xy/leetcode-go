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
		want *TreeNode
	}{
		{
			name: "Test Case 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 5,
								Right: &TreeNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		flatten(tt.root)
		if !reflect.DeepEqual(tt.root, tt.want) {
			fmt.Printf("flatten() = %v, want %v", tt.root, tt.want)
		}
	}
}

// morris
func flattenMorris(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			next := cur.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			// 直接将最右边的节点右子树 指向根的右节点
			pre.Right = cur.Right
			// 开始处理左右节点，next就是左节点，直接赋给右节点
			cur.Left, cur.Right = nil, next
		}
		cur = cur.Right
	}
}

// 只需要处理左边不为空的情况
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	// 如果左子树不为空，那么将左子树插入到右子树的位置
	if root.Left != nil {
		left := root.Left
		// 找到左子树的最右侧节点
		for left.Right != nil {
			left = left.Right
		}

		left.Right = root.Right
		root.Right = root.Left
		// 一定要将左节点置空
		root.Left = nil
	}
}
