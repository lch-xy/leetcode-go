package main

import (
	"fmt"
	"sort"
)

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{0, 0}, 1},
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{}, 0},
		{[]int{10}, 1},
		{[]int{9, 1, -3, 2, 4, 8, 3, -1, 6, -2, -4, 7}, 4},
		{[]int{1, 2, 0, 1}, 3},
	}

	for _, tc := range testCases {
		actual := longestConsecutive(tc.nums)
		if actual != tc.expected {
			fmt.Printf("For nums=%v; Expected %d, got %d。", tc.nums, tc.expected, actual)
		}
	}
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	res := 1
	cnt := 1
	for i := 1; i < len(nums); i++ {
		for i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
		if i < len(nums) && nums[i] == nums[i-1]+1 {
			cnt++
		} else {
			cnt = 1
		}
		res = max(res, cnt)
	}
	return res
}
