package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
}

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i, v := range nums {
		j := i + 1
		cnt := 0
		for nums[j%n] <= v && cnt <= n {
			j++
			cnt++
		}
		if cnt > n {
			res[i] = -1
		} else {
			res[i] = nums[j%n]
		}
	}
	return res
}
