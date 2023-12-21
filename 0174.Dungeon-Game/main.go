package main

import "math"

func main() {
	dungeon := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	println(calculateMinimumHP(dungeon))
}

// dp[i,j]：从左上角到当前位置的血量，如果我们正这取写，其实是很难算的
// 但是我们知道如果要活着，到左下角的时候血量一定要大于等于1，那我们可以考虑从后往前推。
// 通过题目可知我们只能向右或者想做移动，如果保证不死的话，血量一定要大于等于1
// dp[i,j] + hungeon[i][j] >= dp[i+1,j]
// dp[i,j] + hungeon[i][j] >= dp[i,j+1]
// 所以最小值是 f[i,j]的最小值是 min(dp[i+1,j] - hungeon[i][j] , dp[i,j+1] - hungeon[i][j])
func calculateMinimumHP(dungeon [][]int) int {
	x := len(dungeon)
	y := len(dungeon[0])

	dp := make([][]int, x+1)
	for i := 0; i <= x; i++ {
		dp[i] = make([]int, y+1)
	}

	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	// 初始化目标dp[x-1,y-1]的右边和下边的值，以便可以进行dp计算，也就是最低1滴血
	dp[x][y-1] = 1
	dp[x-1][y] = 1
	for i := x - 1; i >= 0; i-- {
		for j := y - 1; j >= 0; j-- {
			dp[i][j] = min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
			dp[i][j] = max(1, dp[i][j])
		}
	}
	return dp[0][0]
}
