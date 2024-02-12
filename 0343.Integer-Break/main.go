package main

import (
	"fmt"
)

func main() {
	fmt.Print(integerBreak(8))
}

// dp[i]：将i拆分成k个正整数，使它们乘积最大化
// 如果分成2个数乘积：j * (i - j)
// 如果分成3个及以上数乘积：j * dp[i - j]，因为dp数组就代表至少分成2个数
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}
