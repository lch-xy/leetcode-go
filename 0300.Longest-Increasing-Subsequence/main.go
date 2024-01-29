package main

import "fmt"

func main() {
	tests := []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{10, 9, 2, 5, 3, 7, 101, 18},
			expect: 4,
		},
		{
			nums:   []int{0, 1, 0, 3, 2, 3},
			expect: 4,
		},
		{
			nums:   []int{7, 7, 7, 7, 7, 7, 7},
			expect: 1,
		},
	}

	for _, test := range tests {
		result := lengthOfLIS(test.nums)
		if result != test.expect {
			fmt.Printf("lengthOfLIS(%v) = %d; expect %d", test.nums, result, test.expect)
		}
	}
}

// [10,9,2,5,3,7,101,18]
// dp[i]:[0,i]内，以nums[i]结束的最长递增子序列的长度
// 第一个for循环是算dp[i]
// 第二个for循环是找到dp[i]的最大值
func lengthOfLIS(nums []int) int {
	size := len(nums)
	dp := make([]int, size)
	for i := range dp {
		dp[i] = 1
	}
	maxResult := 1
	for i := 0; i < size; i++ {
		for j := 0; j <= i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				maxResult = max(maxResult, dp[i])
			}
		}
	}
	return maxResult
}
