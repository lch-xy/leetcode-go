package main

import (
	"math"
)

func main() {
	prices := []int{3, 2, 6, 5, 0, 3}
	println(maxProfit(prices))
	println(maxProfitDp(prices))
}

// dp[i,k] 在第i天交易k次最大的收益
// 如果我们什么都不错做，那就继承前一天的最大收益即可：dp[i,k] = dp[i-1,k]
// 我们也可以选择在第i天将股票卖出，那么在[0，i-1]天里一定还有一次买入
// 如果是第 0 天买入，那收益就是 prices[i] - prices[0]
// 如果是第 1 天买入，那收益就是 prices[i] - prices[1] + dp[0,k-1]（还要加上前一天的最大收益）
// 如果是第 2 天买入，那收益就是 prices[i] - prices[2] + dp[1,k-1]（还要加上前一天的最大收益）
// 如果是第 j 天买入，那收益就是 prices[i] - prices[j] + dp[j-1,k-1]（还要加上前一天的最大收益）
// 因为prices[i]是固定的，所以我们只需要求 prices[j] - dp[j-1,k-1] 的最小值
// 如果第 j 天就是最后我们要选择的买入点，它使得最后的收益最高，dp[j][k-1] 和 dp[j-1][k-1] 一定是相等的。
// 因为第 j 天一定是一个低点而第 j - 1 天是个高点，第 j 天为了得到更高收益肯定选择不操作，所以和第 j - 1 天的收益是一样的
// 所以状态方程我们可以写成prices[j] - dp[j,k-1]，但其实写成prices[j] - dp[j-1,k-1]也可以的
// 为什么从0开始，正好和index对应上，然后只需要返回dp[len-1,k]即可
func maxProfitDp(prices []int) int {
	len := len(prices)
	if len == 0 {
		return 0
	}
	k := 2
	dp := make([][]int, len)
	for i := 0; i < len; i++ {
		dp[i] = make([]int, k+1)
	}
	// k=0时，默认就是收益肯定是0所以需要赋初始值
	for j := 1; j <= k; j++ {
		// 默认取第一天的值
		diff := prices[0]
		for i := 1; i < len; i++ {
			diff = min(diff, prices[i]-dp[i-1][j-1])
			dp[i][j] = max(dp[i-1][j], prices[i]-diff)
		}
	}
	return dp[len-1][k]
}

// hold1：持有第一支股票的最大收益
// sale1：卖出第一支股票的最大收益
// hold2：持有第二支股票的最大收益
// sale2：卖出第二支股票的最大收益

// 第k天持有第一支股票 = max(当前买入，之前就买过了)
// hold1[k] = max(-p,hold1[k-1])
// 第k天卖出第一支股票 = max(今天卖出去了，之前就卖过了)
// sale1[k] = max(hold[k-1]+p,sale1[k-1]])
// 第k天持有第二支股票 = max(今天重新买入第二支，之前就买过了)
// hold2[k] = max(sale[k-1]-p,hold2[k-1])
// 第k天卖出第二支股票 = max(今天卖出去了，之前就卖过了)
// sale2[k] = max(hold2[k-1]+p,sale2[k-1])

// hold1 -> sale1 -> hold2 -> sale2
// 如果要进行下一个操作，那么就需要前一个的状态，只有买入了第一只股票后，才能卖出第一支股票
// 同时这个状态不仅来自于之前的操作，也可以不操作继承前面的状态
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	hold1, sale1, hold2, sale2 := math.MinInt, 0, math.MinInt, 0
	for i := 0; i < len(prices); i++ {
		price := prices[i]
		hold1_old := hold1
		sale1_old := sale1
		hold2_old := hold2
		sale2_old := sale2

		hold1 = max(-price, hold1_old)
		sale1 = max(hold1_old+price, sale1_old)
		hold2 = max(sale1_old-price, hold2_old)
		sale2 = max(hold2_old+price, sale2_old)
	}
	return max(sale1, sale2)
}
