package main

import (
	"fmt"
	"math"
)

func main() {
	tests := []struct {
		matrix [][]int
		k      int
		want   int
	}{
		{
			matrix: [][]int{
				{1, 0, 1},
				{0, -2, 3},
			},
			k:    2,
			want: 2,
		},
		{
			matrix: [][]int{
				{2, 2, -1},
			},
			k:    3,
			want: 3,
		},
		{
			matrix: [][]int{
				{5, -4, 1},
				{-3, 2, 3},
				{6, -4, 0},
			},
			k:    8,
			want: 8,
		},
	}

	for _, tt := range tests {
		got := maxSumSubmatrix(tt.matrix, tt.k)
		if got != tt.want {
			fmt.Printf("maxSumSubmatrix(%v, %v) = %v, want %v", tt.matrix, tt.k, got, tt.want)
		}
	}
}

// 第一步类似于0304.Range-Sum-Query-2D-Immutable
// dp[i][j]: 从(0,0)到(i-1,j-1)围成的面积
func maxSumSubmatrix(matrix [][]int, k int) int {
	len1 := len(matrix)
	len2 := len(matrix[0])
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i := 1; i < len2; i++ {
		dp[0][i] = dp[0][i-1] + matrix[0][i]
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	// 算出dp数组后需要遍历计算矩形的面积，找到小于k的最大面积
	maxSum := math.MinInt32
	for row1 := 1; row1 <= len1; row1++ {
		for row2 := row1; row2 <= len1; row2++ {
			for col1 := 1; col1 <= len2; col1++ {
				for col2 := col1; col2 <= len2; col2++ {
					sum := dp[row2][col2] - dp[row1-1][col2] - dp[row2][col1-1] + dp[row1-1][col1-1]
					if row2 >= row1 && col2 >= col1 && sum <= k {
						maxSum = max(maxSum, sum)
					}
				}
			}
		}
	}

	return maxSum
}
