package main

import (
	"fmt"
	"reflect"
)

func main() {
	type TestCase struct {
		left     int
		right    int
		expected []int
	}

	testCases := []TestCase{
		{
			left:     1,
			right:    22,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
		{
			left:     47,
			right:    85,
			expected: []int{48, 55, 66, 77},
		},
	}

	for i, tc := range testCases {
		result := selfDividingNumbers(tc.left, tc.right)
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test case %d passed\n", i+1)
		} else {
			fmt.Printf("Test case %d failed: expected %v, got %v\n", i+1, tc.expected, result)
			return
		}
	}
}
func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	for i := left; i <= right; i++ {
		cur := i
		flag := true
		for cur > 0 {
			remainder := cur % 10
			cur = cur / 10
			if remainder == 0 || i%remainder != 0 {
				flag = false
			}
		}
		if flag {
			res = append(res, i)
		}
	}
	return res
}
