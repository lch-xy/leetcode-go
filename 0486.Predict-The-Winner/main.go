package main

import "fmt"

func main() {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1},
			want: true,
		},
		{
			nums: []int{1, 5, 2},
			want: false,
		},
		{
			nums: []int{1, 5, 233, 7},
			want: true,
		},
		{
			nums: []int{1, 1, 1},
			want: true,
		},
		{
			nums: []int{2, 4, 55, 6, 8},
			want: false,
		},
		{
			nums: []int{1, 5, 2, 4, 6},
			want: true,
		},
	}

	for _, tt := range tests {
		got := predictTheWinner(tt.nums)
		if got != tt.want {
			fmt.Printf("predictTheWinner(%v) = %v, want %v。", tt.nums, got, tt.want)
		}
	}

	for _, tt := range tests {
		got := predictTheWinnerDp(tt.nums)
		if got != tt.want {
			fmt.Printf("predictTheWinnerDp(%v) = %v, want %v。", tt.nums, got, tt.want)
		}
	}
}

// dp[i][j]: nums[i] -> nums[j] 能够拿到最大的净分差
// nums[1,2,3,4,5,6,7]   dp[i+1][j] , dp[i][j-1] -> dp[i][j]
// 遍历顺序如下:
// -   0   1   2   3   4   5   6
// 0   1   8  14  19  23  26  28
// 1   0   2   9  15  20  24  27
// 2   0   0   3  10  16  21  25
// 3   0   0   0   4  11  17  22
// 4   0   0   0   0   5  12  18
// 5   0   0   0   0   0   6  13
// 6   0   0   0   0   0   0   7

func predictTheWinnerDp(nums []int) bool {
	size := len(nums)
	if size == 1 {
		return true
	}
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
		dp[i][i] = nums[i]
	}

	// 我们在计算dp[i][j]的时候，必须要知道dp[i+1][j] 和 dp[i][j-1]
	// 所以在遍历的时候，应该是斜着遍历的
	for length := 2; length <= size; length++ {
		for i := 0; i <= size-length; i++ {
			j := i + length - 1
			// 如果我选择左边nums[i]，那么就需要减去右边的最大净利润
			// 如果我选择右边nums[j]，那么就需要减去左边的最大净利润
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}

	return dp[0][size-1] >= 0
}

func predictTheWinner(nums []int) bool {
	return scoreDiff(nums, 0, len(nums)-1) >= 0
}

// 递归函数返回的是首位两个玩家的差值
// 如果当前玩家选了首部的数字，那么剩下的数组部分首尾两个玩家得分差就是 nums[start] - scoreDiff(nums, start+1, end)
// 如果当前玩家选了尾部的数字，那么剩下的数组部分首尾两个玩家得分差就是 nums[end] - scoreDiff(nums, start, end-1)
// 这两者之前我们去最大的值
func scoreDiff(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	return max(nums[start]-scoreDiff(nums, start+1, end), nums[end]-scoreDiff(nums, start, end-1))
}
