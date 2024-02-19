package main

import "fmt"

func main() {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{1, 1, 1, 1, 1},
			target: 3,
			want:   5,
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   3,
		},
		{
			nums:   []int{1, 2, 7, 9},
			target: 9,
			want:   0,
		},
		{
			nums:   []int{1, 0},
			target: 1,
			want:   2,
		},
	}

	for _, tt := range tests {
		got := findTargetSumWays(tt.nums, tt.target)
		if got != tt.want {
			fmt.Printf("findTargetSumWaysDp(%v, %v) = %v, want %v。", tt.nums, tt.target, got, tt.want)
		}
	}

	for _, tt := range tests {
		got := findTargetSumWaysDp(tt.nums, tt.target)
		if got != tt.want {
			fmt.Printf("findTargetSumWaysDp(%v, %v) = %v, want %v。", tt.nums, tt.target, got, tt.want)
		}
	}
}

// 正数 + 负数 = target
// 正数 - 负数 = sum
// 正数 = (sum + target)/ 2
// 就将题目转化成能否凑出(sum + target)/2
func findTargetSumWaysDp(nums []int, target int) int {

	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < target || (sum+target)%2 != 0 {
		return 0
	}

	newTarget := (sum + target) / 2

	if newTarget < 0 {
		return 0
	}

	dp := make([]int, newTarget+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := newTarget; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[newTarget]
}

func findTargetSumWays(nums []int, target int) int {
	res := 0
	helper(nums, target, 0, 0, &res)
	return res
}

func helper(nums []int, target, curIndex, total int, res *int) {
	if len(nums) == curIndex {
		if total == target {
			*res++
		}
		return
	}
	helper(nums, target, curIndex+1, total+nums[curIndex], res)
	helper(nums, target, curIndex+1, total-nums[curIndex], res)
}
