package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {

	type TestCase struct {
		// 输入数据
		changed  []int
		expected []int
	}
	testCases := []TestCase{
		{
			changed:  []int{1, 3, 4, 2, 6, 8},
			expected: []int{1, 3, 4},
		},
		{
			changed:  []int{6, 3, 0, 1},
			expected: []int{},
		},
		{
			changed:  []int{1},
			expected: []int{},
		},
		{
			changed:  []int{0, 0, 0, 0},
			expected: []int{0, 0},
		},
		{
			changed:  []int{1, 2, 3, 3},
			expected: []int{},
		},
	}

	for i, tc := range testCases {
		result := findOriginalArray(tc.changed)
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test case %d passed\n", i+1)
		} else {
			fmt.Printf("Test case %d failed: expected %v, got %v\n", i+1, tc.expected, result)
			return
		}
	}
}

func findOriginalArray(changed []int) []int {
	l := len(changed)
	if l%2 != 0 {
		return []int{}
	}
	sort.Ints(changed)
	res := []int{}

	// use array to save the count of num
	countArray := make([]int, changed[l-1]+1)
	for _, v := range changed {
		countArray[v]++
	}
	for _, v := range changed {
		// means "v" is the double num
		if countArray[v] == 0 {
			continue
		}
		// out of index or can't find the double num
		if v*2 > changed[l-1] || countArray[v*2] == 0 {
			return []int{}
		}
		// decrease the count
		countArray[v]--
		countArray[2*v]--
		res = append(res, v)
	}
	// no need to calculate the length
	// because we have already handle the abnormal case
	return res
}
