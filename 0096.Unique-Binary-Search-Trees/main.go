package main

import (
	"fmt"
)

func main() {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 5},
		{4, 14},
	}

	for _, tt := range tests {
		if got := numTreesDp(tt.n); got != tt.want {
			fmt.Printf("numTrees(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

// 思路和95题一样，求解dp[i][j]时要将全部的结果加起来
func numTreesDp(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for i := 1; i <= n-length+1; i++ {
			j := i + length - 1
			for root := i; root <= j; root++ {
				if root == i {
					dp[i][j] += dp[root+1][j]
				} else if root == j {
					dp[i][j] += dp[i][root-1]
				} else {
					dp[i][j] += dp[i][root-1] * dp[root+1][j]
				}
			}
		}
	}
	return dp[1][n]
}

// 思路和95题一样，主要加了个缓存memo，不然会超时
func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	memo := make(map[string]int)
	return helper(1, n, memo)
}

func helper(start, end int, memo map[string]int) int {
	if start > end {
		return 1
	}
	key := fmt.Sprintf("%d-%d", start, end)
	if val, ok := memo[key]; ok {
		return val
	}
	res := 0
	for root := start; root <= end; root++ {
		res += helper(start, root-1, memo) * helper(root+1, end, memo)
	}
	memo[key] = res
	return res
}
