package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 1, 2}, 1},
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
	}

	for _, tc := range testCases {
		result := findMin(tc.nums)
		if result != tc.expected {
			fmt.Printf("findMin(%v) = %v; expected %v。", tc.nums, result, tc.expected)
		}
	}
}

// use binary search
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	// why left + 1 < right ?
	// because we want to stop the loop where difference between left and right valuse is 1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid
		} else if nums[mid] < nums[left] {
			right = mid
		} else {
			// if sorting is correct then return left
			return nums[left]
		}
	}
	// return the minimum value
	return min(nums[left], nums[right])
}
