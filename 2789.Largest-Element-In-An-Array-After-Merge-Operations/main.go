package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		name   string
		nums   []int
		expect int64
	}{
		{
			name:   "Example 1",
			nums:   []int{3, 7},
			expect: 10,
		},
		{
			name:   "Example 2",
			nums:   []int{2, 4, 5},
			expect: 11,
		},
		{
			name:   "Example 3",
			nums:   []int{9, 8, 1, 0, 1, 9, 9, 8},
			expect: 37,
		},
		{
			name:   "Example 4",
			nums:   []int{1, 2, 3, 4, 5},
			expect: 15,
		},
		{
			name:   "Example 5",
			nums:   []int{5, 4, 3, 2, 1},
			expect: 5,
		},
	}

	for _, tc := range testCases {
		result := maxArrayValue(tc.nums)
		if result != tc.expect {
			fmt.Printf("Expected %d, but got %d", tc.expect, result)
		}
	}
}

func maxArrayValue(nums []int) int64 {
	lastIndex := len(nums) - 1
	res := 0
	cur := nums[lastIndex]
	// the question tells us nums[i] <= nums[i+1]
	// so only need to traverse from back to front once
	for i := lastIndex - 1; i >= 0; i-- {
		if cur >= nums[i] {
			cur += nums[i]
		} else {
			cur = nums[i]
		}
		res = max(res, cur)
	}
	return int64(res)
}
