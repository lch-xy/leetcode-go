package main

import "fmt"

func main() {
	fmt.Println(maxOperations([]int{1, 1, 1, 1, 1, 1}))
}

func maxOperations(nums []int) int {
	denominator := nums[0] + nums[1]
	cnt := 0
	for i := 1; i*2 <= len(nums); i++ {
		if nums[i*2-2]+nums[i*2-1] == denominator {
			cnt++
			continue
		}
		break
	}
	return cnt
}
