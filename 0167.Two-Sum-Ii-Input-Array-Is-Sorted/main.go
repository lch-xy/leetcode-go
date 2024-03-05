package main

import (
	"fmt"
	"reflect"
)

func main() {
	testCases := []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{5, 25, 75}, 100, []int{2, 3}},
		// 其他测试用例...
	}

	for _, tc := range testCases {
		actual := twoSum(tc.numbers, tc.target)
		if !reflect.DeepEqual(actual, tc.expected) {
			fmt.Printf("For numbers=%v and target=%d; Expected %v, got %v。", tc.numbers, tc.target, tc.expected, actual)
		}
	}
}

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			break
		}
	}
	return []int{left + 1, right + 1}
}
