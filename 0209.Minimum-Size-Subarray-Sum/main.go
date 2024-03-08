package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		target   int
		nums     []int
		expected int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{7, []int{1, 2, 3}, 0},
		{11, []int{1, 2, 3, 4, 5}, 3},
		{3, []int{}, 0},
		// 边界测试用例，非常大的 target
		{100, []int{1, 2, 3, 4, 5, 50, 10, 20, 30}, 4},
		// 边界测试用例，只包含一个数字的数组
		{3, []int{10}, 1},
		// 边界测试用例，target 为数组所有元素的总和
		{15, []int{1, 2, 3, 4, 5}, 5},
	}

	for _, tc := range testCases {
		actual := minSubArrayLen(tc.target, tc.nums)
		if actual != tc.expected {
			fmt.Printf("minSubArrayLen(%d, %v) = %d; expected %d。", tc.target, tc.nums, actual, tc.expected)
		}
	}
}

// use sliding window
func minSubArrayLen(target int, nums []int) int {
	size := len(nums)
	slow, fast := 0, 0
	sum := 0
	res := size + 1
	for fast < size {
		sum += nums[fast]
		// when the sum is greater than or equal to the target value
		// it indicates that wt should shrink the window
		// therefore incrementing the "slow" point
		for slow < size && sum >= target {
			res = min(res, fast-slow+1)
			sum -= nums[slow]
			slow++
		}
		fast++
	}
	if res == size+1 {
		return 0
	}
	return res
}
