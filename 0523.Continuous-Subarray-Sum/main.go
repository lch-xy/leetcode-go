package main

import "fmt"

func main() {
	testCases := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{23, 2, 4, 6, 7}, 6, true},
		{[]int{23, 2, 6}, 6, false},
		{[]int{23, 2, 6, 4, 7}, 13, false},
		{[]int{23, 6, 9}, 6, false},
		{[]int{0, 0}, 1, true},
	}

	for _, tc := range testCases {
		result := checkSubarraySum(tc.nums, tc.k)
		if result != tc.expected {
			fmt.Printf("checkSubarraySum(%v, %v) = %v; expected %v。", tc.nums, tc.k, result, tc.expected)
		}
	}
}

// 若数字a和b分别除以数字c，若得到的余数相同，那么(a-b)必定能够整除c
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 || k == 0 {
		return false
	}
	hmap := make(map[int]int)
	// base case
	hmap[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		mod := sum % k
		// 如果存在，那么就说明从上一个相同模值的索引到当前索引，构成的子数组的和是k的倍数
		if prevIndex, ok := hmap[mod]; ok {
			if i-prevIndex > 1 {
				return true
			}
		} else {
			hmap[mod] = i
		}

	}
	return false
}

// func checkSubarraySum(nums []int, k int) bool {
// 	if len(nums) < 2 {
// 		return false
// 	}
// 	if k == 0 {
// 		return false
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		total := nums[i]
// 		for j := i + 1; j < len(nums); j++ {
// 			total += nums[j]
// 			if total%k == 0 || total == 0 {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
