package main

import "fmt"

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{-4, -3, -1, 0, 1, 4, 5}, expected: 3},
		{nums: []int{-2, 0, 1, 3, 5, 7}, expected: 4},
		{nums: []int{-10, -5, 0, 5, 10}, expected: 2},
		{nums: []int{5, 10, 12, 13}, expected: 4},
	}

	for i, tc := range testCases {
		maxCount := maximumCount(tc.nums)
		if maxCount != tc.expected {
			fmt.Printf("Test case %d failed: expected %d, got %d\n", i+1, tc.expected, maxCount)
			return
		}
		fmt.Printf("Test case %d passed\n", i+1)
	}
}
func maximumCount(nums []int) int {
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return len(nums)
	}
	cnt1 := 0
	cnt2 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			cnt1++
		} else if nums[i] > 0 {
			cnt2++
		}
	}
	return max(cnt1, cnt2)
}
