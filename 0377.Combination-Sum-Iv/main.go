package main

import "fmt"

func main() {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 2, 3}, 4, 7},
		{[]int{9}, 3, 0},
		{[]int{3, 2, 1}, 4, 7},
	}

	for _, tc := range testCases {
		got := combinationSum4(tc.nums, tc.target)
		if got != tc.want {
			fmt.Printf("combinationSum4(%v, %v) = %v; want %v", tc.nums, tc.target, got, tc.want)
		}
	}

	fmt.Println("")
	fmt.Println("")

	for _, tc := range testCases {
		got := combinationSum4Dp(tc.nums, tc.target)
		if got != tc.want {
			fmt.Printf("combinationSum4(%v, %v) = %v; want %v", tc.nums, tc.target, got, tc.want)
		}
	}
}

// 思路还是和背包问题差不多，把背包大小target装满
func combinationSum4Dp(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			// 要加判断 不然会越界
			if i-num >= 0 {
				dp[i] = dp[i] + dp[i-num]
			}
		}
	}
	return dp[target]
}

// 递归解法，会超时
func combinationSum4(nums []int, target int) int {
	res := 0
	helper(nums, target, &res)
	return res
}

func helper(nums []int, target int, res *int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*res++
		return
	}
	for i := 0; i < len(nums); i++ {
		helper(nums, target-nums[i], res)
	}
}
