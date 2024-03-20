package main

import "fmt"

func main() {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{nums: []int{1, 4, 3, 7, 4, 5}, k: 3, expected: 15},
		{nums: []int{5, 5, 1, 1}, k: 0, expected: 10},
	}

	for i, tc := range testCases {
		result := maximumScore(tc.nums, tc.k)
		if result != tc.expected {
			fmt.Printf("测试用例 %d 失败：期望输出 %d，实际输出 %d\n", i+1, tc.expected, result)
		}
	}
}

func maximumScore(nums []int, k int) int {
	length := len(nums)
	left := make([]int, length)
	right := make([]int, length)
	// init default value
	for i := range left {
		left[i] = -1
	}
	for i := range right {
		right[i] = length
	}

	stack := make([]int, 0)

	// use monotonic stack
	// to find the previous smaller elements
	for i := 0; i < length; i++ {
		for len(stack) != 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// clear up
	stack = make([]int, 0)

	// to find the next smaller elements
	for i := length - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	maxValue := 0
	for i := 0; i < length; i++ {
		// check position
		if left[i]+1 <= k && k <= right[i]-1 {
			maxValue = max(maxValue, nums[i]*(right[i]-left[i]-1))
		}
	}

	return maxValue
}
