package main

import "fmt"

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{6, 5, 5}, 5},
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		// 其他测试用例...
	}

	for _, tc := range testCases {
		actual := majorityElementBMVA(tc.nums)
		if actual != tc.expected {
			fmt.Printf("For nums=%v; Expected %d, got %d。", tc.nums, tc.expected, actual)
		}
	}
}

// 摩尔投票算法（Boyer-Moore Voting Algorithm）
// 其实还需要遍历一下 看看是否满足大于一半 但这里题目说满足就没遍历了
func majorityElementBMVA(nums []int) int {
	candidate, count := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func majorityElement(nums []int) int {
	size := len(nums)
	cache := make(map[int]int)
	for i := 0; i < size; i++ {
		if _, ok := cache[nums[i]]; ok {
			cache[nums[i]]++
		} else {
			cache[nums[i]] = 1
		}

		if cache[nums[i]] > size/2 {
			return nums[i]
		}
	}
	return -1
}
