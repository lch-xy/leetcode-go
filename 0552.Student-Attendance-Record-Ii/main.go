package main

import "fmt"

func main() {
	testCases := []struct {
		n        int
		expected int
	}{
		{1, 3},
		{2, 8},
		{3, 19},
		{10101, 183236316},
		// 添加更多测试用例
	}

	for _, tc := range testCases {
		result := checkRecord(tc.n)
		if result != tc.expected {
			fmt.Printf("calculate(%v) = %v; expected %v", tc.n, result, tc.expected)
		}
	}
}

// dp[i][j][k]: i代表第几天，j代表缺席天数，k代表连续迟到天数
func checkRecord(n int) int {
	mod := 1000000007

	dp := make([][][]int, n)

	for i := range dp {
		dp[i] = make([][]int, 2)
		for j := range dp[i] {
			dp[i][j] = make([]int, 3)
		}
	}

	dp[0][0][0], dp[0][1][0], dp[0][0][1] = 1, 1, 1

	for i := 1; i < n; i++ {
		// Absent，缺勤
		// 当前 ‘A’ 可以从3种类型得到， 1：很正常 2：迟到1次 3：迟到2次。
		// 因为如果不是连续三次迟到，都会被重置
		dp[i][1][0] = (dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2]) % mod

		// Late，迟到
		// 因为 ‘L’ 是需要连续的，所以只能从上一个状态流转过来
		dp[i][0][1] = dp[i-1][0][0]
		dp[i][0][2] = dp[i-1][0][1]
		dp[i][1][1] = dp[i-1][1][0]
		dp[i][1][2] = dp[i-1][1][1]

		// Present，到场
		// 当前 ‘P’ 且 0 缺席，可以从3种类型得到， 1：很正常 2：迟到1次 3：迟到2次。
		dp[i][0][0] = (dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2]) % mod
		// 当前 ‘P’ 且 1 缺席，这里多加了一个dp[i][1][0]，因为也可以由上面的3中状态转移而来
		dp[i][1][0] = (dp[i][1][0] + dp[i-1][1][0] + dp[i-1][1][1] + dp[i-1][1][2]) % mod
	}

	// 计算 n-1天时，缺席少于2，迟到连续少于3
	sum := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			sum = (sum + dp[n-1][i][j]) % mod
		}
	}

	return sum
}
