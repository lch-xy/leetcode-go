package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("minOperations([]int{8, 5, 9, 9, 8, 4}): %v\n", minOperations([]int{8, 5, 9, 9, 8, 4}))
	fmt.Printf("minOperations([]int{8,10,16,18,10,10,16,13,13,16}): %v\n", minOperations([]int{8, 10, 16, 18, 10, 10, 16, 13, 13, 16}))

}

func minOperations(nums []int) int {
	sort.Ints(nums)
	setNums := []int{nums[0]}

	// remove duplicates
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		setNums = append(setNums, nums[i])
	}

	lengthOne := len(nums)
	lengthTwo := len(setNums)
	res := lengthOne

	// use a sliding window to count the number of operations
	for i, j := 0, 0; i < lengthTwo; i++ {
		// (setNums[j]-setNums[i] < lengthOne) means want to fill in the number between [setNums[i],setNums[j]]
		for j < lengthTwo && setNums[j]-setNums[i] < lengthOne {
			j++
		}
		// (j-i) means how many numbers there are between [setNums[i],setNums[j]]
		// other numbers need to replaced with number in [setNums[i],setNums[j]]
		res = min(res, lengthOne-(j-i))
	}
	return res
}
