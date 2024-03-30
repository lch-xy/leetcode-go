package main

import (
	"fmt"
	"math"
)

func main() {
	// Define test cases slice
	testCases := []struct {
		nums     []int // 整数数组
		expected int   // 预期的山形三元组的元素和
	}{
		{nums: []int{1, 3, 1, 2, 1}, expected: 4},
		// You can add more test cases here
	}

	// Iterate through test cases and run tests
	for i, tc := range testCases {
		result := minimumSum(tc.nums)
		if result != tc.expected {
			fmt.Printf("Test case %d failed: Expected output %d, got %d\n", i+1, tc.expected, result)
		}
	}
}

// iterator all possibilities
func minimumSum(nums []int) int {
	length := len(nums)
	left, mid, rifht := 0, 0, 0
	res := math.MaxInt32
	for left = 0; left < length; left++ {
		for mid = left + 1; mid < length; mid++ {
			for rifht = mid + 1; rifht < length; rifht++ {
				if nums[left] < nums[mid] && nums[mid] > nums[rifht] {
					if res > nums[left]+nums[mid]+nums[rifht] {
						res = nums[left] + nums[mid] + nums[rifht]
					}
				}
			}
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
