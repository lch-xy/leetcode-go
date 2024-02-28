package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		m, n, maxMove, startRow, startColumn int
		expected                             int
	}{
		{2, 2, 2, 0, 0, 6},
		{1, 3, 3, 0, 1, 12},
		// 添加更多测试用例
	}

	for _, tc := range testCases {
		result := findPathsDfs(tc.m, tc.n, tc.maxMove, tc.startRow, tc.startColumn)
		if result != tc.expected {
			fmt.Printf("findPaths(%v, %v, %v, %v, %v) = %v; expected %v。", tc.m, tc.n, tc.maxMove, tc.startRow, tc.startColumn, result, tc.expected)
		}
	}

	for _, tc := range testCases {
		result := findPathsDp(tc.m, tc.n, tc.maxMove, tc.startRow, tc.startColumn)
		if result != tc.expected {
			fmt.Printf("findPathsDp(%v, %v, %v, %v, %v) = %v; expected %v。", tc.m, tc.n, tc.maxMove, tc.startRow, tc.startColumn, result, tc.expected)
		}
	}
}

// dp[i][j][k]表示从起始位置移动 k 步后到达网格位置(i, j)的路径数量。
func findPathsDp(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, maxMove+1)
		}
	}

	direction := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	for k := 1; k <= maxMove; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				for _, v := range direction {
					ni, nj := i+v[0], j+v[1]
					// 朝上下左右移动一格，如果当前还在格子内，说明(i,j)可以通过(ni,nj)移动得到
					// 所以需要把所有的情况都加起来
					if ni >= 0 && ni < m && nj >= 0 && nj < n {
						dp[i][j][k] = (dp[i][j][k] + dp[ni][nj][k-1]) % int(1e9+7)
					} else {
						dp[i][j][k]++
					}
				}
			}
		}
	}
	return dp[startRow][startColumn][maxMove]
}

func findPathsDfs(m int, n int, maxMove int, startRow int, startColumn int) int {
	cache := make([][][]int, m+1)
	for i := range cache {
		cache[i] = make([][]int, n+1)
		for j := range cache[i] {
			cache[i][j] = make([]int, maxMove+1)
			for k := range cache[i][j] {
				cache[i][j][k] = -1
			}
		}
	}
	return dfs(m, n, maxMove, startRow, startColumn, cache)
}

func dfs(m, n, maxMove, startRow, startColumn int, cache [][][]int) int {
	if startRow < 0 || m <= startRow || startColumn < 0 || n <= startColumn {
		return 1
	}
	if maxMove == 0 {
		return 0
	}
	if cache[startRow][startColumn][maxMove] != -1 {
		return cache[startRow][startColumn][maxMove]
	}
	one := dfs(m, n, maxMove-1, startRow-1, startColumn, cache)
	two := dfs(m, n, maxMove-1, startRow+1, startColumn, cache)
	three := dfs(m, n, maxMove-1, startRow, startColumn-1, cache)
	four := dfs(m, n, maxMove-1, startRow, startColumn+1, cache)
	res := (one + two + three + four) % int(1e9+7)
	cache[startRow][startColumn][maxMove] = res
	return res
}
