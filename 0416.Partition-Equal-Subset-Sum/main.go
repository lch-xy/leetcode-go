package main

import "fmt"

func main() {
	testCases := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Case 1: Can partition",
			nums: []int{1, 5, 11, 5},
			want: true,
		},
		{
			name: "Case 2: Cannot partition",
			nums: []int{1, 2, 3, 5},
			want: false,
		},
		{
			name: "Case 3",
			nums: []int{1, 1},
			want: true,
		},
		{
			name: "Case 4",
			nums: []int{1, 2, 5},
			want: false,
		},
	}

	for _, tc := range testCases {
		got := canPartition(tc.nums)
		if got != tc.want {
			fmt.Printf("canPartition(%v) = %v; want %v", tc.nums, got, tc.want)
		}
	}
}

// 类似0/1背包问题，如果可以分为两个相等数数组，首先满足加起来是偶数
// 接下来可以算出目标数 target = sum / 2
func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	total := sum(nums)
	if total%2 != 0 {
		return false
	}

	target := total / 2

	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[target]
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
