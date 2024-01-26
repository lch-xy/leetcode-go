package main

import (
	"fmt"
	"math"
)

func main() {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},    // 1^2
		{2, 2},    // 1^2 + 1^2
		{3, 3},    // 1^2 + 1^2 + 1^2
		{4, 1},    // 2^2
		{5, 2},    // 2^2 + 1^2
		{9, 1},    // 3^2
		{10, 2},   // 3^2 + 1^2
		{12, 3},   // 2^2 + 2^2 + 2^2
		{13, 2},   // 3^2 + 2^2
		{100, 1},  // 10^2
		{9999, 4}, // 99^2 + 5^2 + 2^2 + 1^2
	}

	for _, tt := range tests {
		if got := numSquaresDp(tt.n); got != tt.want {
			fmt.Printf("numSquaresDp(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

// 类似于背包问题一样，[1,n]个物品，背包大小为i
// dp[i]:[0,n]个物品里，能凑到背包大小为i所用的物品最少得量
func numSquaresDp(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i <= n; i++ {
		// i+j*j <= n防止越界，到此为止
		for j := 1; i+j*j <= n; j++ {
			dp[i+j*j] = min(dp[i+j*j], dp[i]+1)
		}
	}
	return dp[n]
}

// 递归解法超时，无法ac
func numSquares(n int) int {
	res := math.MaxInt32
	maxNumber := int(math.Sqrt(float64(n)))
	helper(n, maxNumber, 0, &res)
	return res
}

func helper(target int, maxNumber int, cnt int, res *int) {
	for i := maxNumber; i > 0; i-- {
		newTarget := target - i*i
		if newTarget < 0 {
			continue
		}
		newCnt := cnt + 1
		if newTarget == 0 {
			*res = min(*res, newCnt)
			return
		}
		helper(newTarget, i, cnt+1, res)
	}
}
