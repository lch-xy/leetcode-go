package main

import "fmt"

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{4, 5, 6, 7, 0, 1, 4}, expected: 0},
		{nums: []int{0, 1, 4, 4, 5, 6, 7}, expected: 0},
		{nums: []int{2, 2, 2, 0, 1}, expected: 0},
		{nums: []int{3, 3, 1, 3}, expected: 1},
		{nums: []int{1, 3, 3}, expected: 1},
	}

	for _, tc := range testCases {
		result := findMin(tc.nums)
		if result != tc.expected {
			fmt.Printf("FAIL: 在数组 %v 中, 找到的最小元素为 %d，预期为 %d\n", tc.nums, result, tc.expected)
		}
	}
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	// we narrow the final result to between the values of nums[left] and nums[right]
	for left+1 < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := left + (right-left)/2
		if nums[left] > nums[mid] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid
		} else {
			// we can't determine the position of the minimum element
			// therefore we reduce the search space by ignoring both the right and the left element
			left++
			right--
		}
	}
	return min(nums[left], nums[right])
}
