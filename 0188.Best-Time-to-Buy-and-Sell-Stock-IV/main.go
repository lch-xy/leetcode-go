package main

func main() {
	println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}

// 解法和第123题一样
func maxProfit(k int, prices []int) int {
	len := len(prices)
	if len == 0 || k <= 0 {
		return 0
	}
	dp := make([][]int, len)
	for i := 0; i < len; i++ {
		// 这里要定义k+1，
		dp[i] = make([]int, k+1)
	}
	for j := 1; j <= k; j++ {
		// 假设第一天买入
		diff := prices[0]
		for i := 1; i < len; i++ {
			diff = min(diff, prices[i-1]-dp[i-1][j-1])
			dp[i][j] = max(dp[i-1][j], prices[i]-diff)
		}
	}
	return dp[len-1][k]
}
