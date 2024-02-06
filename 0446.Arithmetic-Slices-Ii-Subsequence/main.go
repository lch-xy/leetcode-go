package main

import "fmt"

func main() {
	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{"Test Case 1", []int{2, 4, 6, 8, 10}, 7},
		{"Test Case 2", []int{7, 7, 7}, 1},
		{"Test Case 2", []int{7, 7, 7, 7, 7}, 16},
		{"Test Case 3", []int{3, -1, -5, -9}, 3},
		{"Test Case 4", []int{1, 2, 3, 4}, 3},
		{"Test Case 5", []int{1}, 0},
	}

	for _, tc := range testCases {
		got := numberOfArithmeticSlices(tc.nums)
		if got != tc.want {
			fmt.Printf("numberOfArithmeticSlices(%v) = %v; want %v", tc.nums, got, tc.want)
		}
	}
}

// dp[i][diff]：以nums[i]结尾，diff作为公差的组合有几个
// 我们从第二个数开始遍历，因为第一个没有公差
// 我们计算的res的时候，其实是根据前一个数来算的
// 例如我们算到dp[3]的时候，我们要用到的数其实是dp[2][diff]
// 因为以nums[2]结尾，diff作为公差的组合有几个，我们把nums[3]加进去就行了
func numberOfArithmeticSlices(nums []int) int {
	size := len(nums)
	dp := make([]map[int]int, size)
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	res := 0
	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			dp[i][diff] += dp[j][diff] + 1
			if dp[j][diff] > 0 {
				res += dp[j][diff]
			}
		}
	}
	return res
}
