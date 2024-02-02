package main

import (
	"fmt"
	"math"
)

func main() {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Test 1",
			n:    1,
			want: 0,
		},
		{
			name: "Test 2",
			n:    2,
			want: 1,
		},
		{
			name: "Test 3",
			n:    10,
			want: 16,
		},
	}

	for _, tt := range tests {
		if got := getMoneyAmount(tt.n); got != tt.want {
			fmt.Printf("getMoneyAmount() = %v, want %v", got, tt.want)
		}
	}
}

func getMoneyAmount(n int) int {
	// 为什么要+2，因为会存在k-1和k+1的情况，所以默认多开辟2个空间
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	// 为什么length从2开始？
	// 如果length为1的话，说明就一个元素，其实就等于0
	for length := 2; length <= n; length++ {
		// 为什么是i <= n-length+1？
		// 当 i=1 length=2时，我们要的数字其实是[1,2],所以i+length-1=2
		// i + length - 1 <= n 可以推迟 i <= n - length + 1
		for i := 1; i <= n-length+1; i++ {
			j := i + length - 1
			dp[i][j] = math.MaxInt32
			for k := i; k <= j; k++ {
				// Minimax算法
				dp[i][j] = min(dp[i][j], max(dp[i][k-1], dp[k+1][j])+k)
			}
		}
	}
	return dp[1][n]
}
