package main

import "fmt"

func main() {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 1, 5, 8}, 167},
		{[]int{1, 5}, 10},
		{[]int{9, 76, 64, 21}, 62016},
		{[]int{}, 0},
		{[]int{3, 1, 5, 8, 4, 2}, 238},
	}

	for _, tt := range tests {
		got := maxCoins(tt.nums)
		if got != tt.want {
			fmt.Printf("maxCoins(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

// dp[i][j]表示戳破i到j之间（开区间，不包括i和j，因为我们手动加了两个1进去）所有气球能得到的最大硬币数。
// nums = [1,3,1,5,8,1]
// 当长度为3时，我们有以下的子区间：[1,3,1]，[3,1,5]，[1,5,8]，[5,8,1]。我们可以计算出dp[0][2] = 3，dp[1][3] = 15，dp[2][4] = 40，dp[3][5] = 40。
// 当长度为4时，我们有以下的子区间：[1,3,1,5]，[3,1,5,8]，[1,5,8,1]。我们可以计算出dp[0][3] = 30，dp[1][4] = 135，dp[2][5] = 120。
// 当长度为5时，我们有以下的子区间：[1,3,1,5,8]，[3,1,5,8,1]。我们可以计算出dp[0][4] = 159，dp[1][5] = 167。
// 当长度为6时，我们只有一个子区间：[1,3,1,5,8,1]。我们可以计算出dp[0][5] = 167。
func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	size := len(nums)

	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
	}
	for length := 3; length <= size; length++ {
		for i := 0; i <= size-length; i++ {
			// j的最大值应该是size-1
			// 所以我们可以反推i的最大值 size - 1 = i + length - 1
			// ==> i = size - length
			j := i + length - 1
			// 这是因为k是在i和j之间的一个位置
			for k := i + 1; k < j; k++ {
				// 为什么是nums[i]*nums[k]*nums[j]而不是nums[k-1]*nums[k]*nums[k+1]
				// 因为dp[i][k]和dp[k][j]代表[i,k]和[k,j]之间的气球已经被打爆了，只剩下i和j了
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[k]*nums[j])
			}
		}
	}
	return dp[0][size-1]
}
