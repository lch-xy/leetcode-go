package main

import "fmt"

func main() {
	testCases := []struct {
		prices []int
		want   int
	}{
		{[]int{1, 2, 3, 0, 2}, 3},
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tc := range testCases {
		got := maxProfit(tc.prices)
		if got != tc.want {
			fmt.Printf("maxProfit(%v) = %d; want %d", tc.prices, got, tc.want)
		}
	}
}

// dp[i][1]表示第i天结束时，手上有股票的最大利润。
func maxProfit(prices []int) int {
	size := len(prices)
	dp := make([][2]int, size+1)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < size; i++ {
		// 表示第i天结束时，手上没有股票，可能是因为第i天没有进行任何操作，或者因为第i天卖出了股票
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

		if i > 1 {
			// 因为有个冷冻期，所以是dp[i-2][0]
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		} else {
			dp[i][1] = max(dp[i-1][1], -prices[i])
		}
	}
	return dp[size-1][0]
}
