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
		n    int
		want []*TreeNode
	}{
		{0, nil},
		{1, []*TreeNode{{Val: 1}}},
		{2, []*TreeNode{{Val: 1, Right: &TreeNode{Val: 2}}, {Val: 2, Left: &TreeNode{Val: 1}}}},
	}

	for _, tt := range tests {
		got := generateTrees(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("generateTrees(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTree(1, n)
}

// 返回[start,end]之间的所有的组合
// 主要思路就是递归的得到左右子树的所有构建方式，然后遍历左右子树和根节点进行组合
// 然后将所有结果返回姐上一层
func generateTree(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	res := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftNodes := generateTree(start, i-1)
		rightNodes := generateTree(i+1, end)
		for _, left := range leftNodes {
			for _, right := range rightNodes {
				currentTree := &TreeNode{i, left, right}
				res = append(res, currentTree)
			}
		}
	}
	return res
}
