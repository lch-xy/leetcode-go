package main

import "fmt"

func main() {
	//nums := []int{1, 2, 3, 4, 3, 2, 5}
	//nums := []int{2, 2, 2, 2, 2, 2, 2}
	nums := []int{1, 3, 4}
	fmt.Print(resultsArray(nums, 2))
}

func resultsArray(nums []int, k int) []int {
	end := len(nums)
	if k == 1 {
		return nums
	}
	start := 0
	cur := 1
	point := -1
	res := []int{}
	for cur < end {
		if nums[cur] != nums[cur-1]+1 {
			point = cur - 1
		}
		if cur-start == k-1 {
			if start <= point && point <= cur {
				res = append(res, -1)
			} else {
				res = append(res, nums[cur])
			}
			start++
			cur++
		} else {
			cur++
		}
	}
	return res
}

func resultsArray2(nums []int, k int) []int {
	l := len(nums)
	res := make([]int, l-k+1)
	for i := 0; i < l; i++ {
		res[i] = -1
	}
	cnt := 0
	for i, v := range nums {
		if i == 0 || v == nums[i-1]+1 {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k {
			res[i-k+1] = v
		}
	}
	return res
}
