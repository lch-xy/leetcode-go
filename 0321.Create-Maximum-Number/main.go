package main

import (
	"fmt"
	"reflect"
)

func main() {
	tests := []struct {
		nums1  []int
		nums2  []int
		k      int
		expect []int
	}{
		{
			nums1:  []int{3, 4, 6, 5},
			nums2:  []int{9, 1, 2, 5, 8, 3},
			k:      5,
			expect: []int{9, 8, 6, 5, 3},
		},
		{
			nums1:  []int{6, 7},
			nums2:  []int{6, 0, 4},
			k:      5,
			expect: []int{6, 7, 6, 0, 4},
		},
		{
			nums1:  []int{3, 9},
			nums2:  []int{8, 9},
			k:      3,
			expect: []int{9, 8, 9},
		},
	}

	for _, test := range tests {
		result := maxNumber(test.nums1, test.nums2, test.k)
		if !reflect.DeepEqual(result, test.expect) {
			fmt.Printf("maxNumber(%v, %v, %d) = %v; expect %v", test.nums1, test.nums2, test.k, result, test.expect)
		}
	}

}

// 解题思路：将k分为i和k-i，分别从nums1和nums2中找到对对应的数，然后合并
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	res := make([]int, k)
	for i := 1; i <= k; i++ {
		if i > len(nums1) || k-i > len(nums2) {
			continue
		}
		numOne := getMax(nums1, i)
		numsTwo := getMax(nums2, k-i)

		mergeNumber := merge(numOne, numsTwo)
		if !greater(res, 0, mergeNumber, 0) {
			res = mergeNumber
		}
	}

	return res
}

// 判断nums1是否大于nums2
// 先判断大小，再根据长度
func greater(nums1 []int, index1 int, nums2 []int, index2 int) bool {
	for index1 < len(nums1) && index2 < len(nums2) && nums1[index1] == nums2[index2] {
		index1++
		index2++
	}
	// 要加上index1 < len(nums1)，不然nums1遍历完后会越界
	return index2 == len(nums2) || (index1 < len(nums1) && nums1[index1] > nums2[index2])
}

// 合并两个数组
func merge(numOne, numTwo []int) []int {
	index1 := 0
	index2 := 0
	res := []int{}
	for index1 < len(numOne) || index2 < len(numTwo) {
		if greater(numOne, index1, numTwo, index2) {
			res = append(res, numOne[index1])
			index1++
		} else {
			res = append(res, numTwo[index2])
			index2++
		}
	}
	return res
}

// [3,4,6,5]  k=2
// stack : [3,4] index = 2
func getMax(nums []int, k int) []int {
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		// 如果不单调，栈不为空且栈内元素+剩余元素>k
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] && len(stack)+len(nums)-i-1 >= k {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return stack[:k]
}
