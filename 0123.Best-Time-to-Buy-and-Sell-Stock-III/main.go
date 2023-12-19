package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{3, 2, 6, 5, 0, 3}
	result := maxProfit(prices)
	fmt.Println(result)
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
