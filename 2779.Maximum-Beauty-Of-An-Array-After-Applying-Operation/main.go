package main

import "sort"

// question:
// 1.choose an index i from the array which has not been chosen before.
// 2.replace the element at nums[i] with any number within the range [nums[i]-k,nums[i]+k]
// order of arrays is not necessary for us , so first step we can sorted the arrays
// for example: [2,4,1,6] k=2
// sorted the arrays -> [1,2,4,6] -> [-1,3] [0,4] [2,6] [4,8]
// these interval means 1 can be replced -1,0,1,2,3 , 2 can be replaced 0,1,2,3,4 ......
// if the interval of [-1,3] and [0，4] can coincidence , means the number 2 can replaced the number 1
// if the left number's index is x ,the right number's index is y
// ==> nums[x] + k > nums[y] - k
func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)
	cnt, left := 0, 0
	for right, _ := range nums {
		for nums[right]-nums[left] > 2*k {
			left++
		}
		cnt = max(cnt, right-left+1)
	}
	return cnt
}
