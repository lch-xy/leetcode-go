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
		got := generateTreesDp(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("generateTrees(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

// dp[i][j]: 以[i,j]的子串，构建二叉搜索树的所以组合
func generateTreesDp(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	dp := make([][][]*TreeNode, n+1)
	for i := range dp {
		dp[i] = make([][]*TreeNode, n+1)
	}
	// 处理对角线 只有一种情况
	for i := 1; i <= n; i++ {
		dp[i][i] = append(dp[i][i], &TreeNode{i, nil, nil})
	}
	// 因为我们需要的最终答案是在右上角,所以我们遍历要斜着遍历
	// 前面两个for循环其实就是在构造斜着遍历
	// 后面三个for循环才是真正在干事情
	// 为什么从2开始，因为前面已经处理过length=1的情况了
	for length := 2; length <= n; length++ {
		for i := 1; i <= n-length+1; i++ {
			j := i + length - 1
			for root := i; root <= j; root++ {
				// 左右两个节点单独处理
				if root == i {
					for _, right := range dp[root+1][j] {
						dp[i][j] = append(dp[i][j], &TreeNode{root, nil, right})
					}
				} else if root == j {
					for _, left := range dp[i][root-1] {
						dp[i][j] = append(dp[i][j], &TreeNode{root, left, nil})
					}
				} else {
					for _, left := range dp[i][root-1] {
						for _, right := range dp[root+1][j] {
							dp[i][j] = append(dp[i][j], &TreeNode{root, left, right})
						}
					}
				}
			}
		}
	}
	return dp[1][n]
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
