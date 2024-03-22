package main

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)
	if nums[0] != 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 != nums[i] {
			return nums[i-1] + 1
		}
	}
	return nums[len(nums)-1] + 1
}
