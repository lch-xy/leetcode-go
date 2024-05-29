package main

import "fmt"

func main() {
	fmt.Printf("validPartition([]int{4, 4, 4, 5, 6}): %v\n", validPartition([]int{4, 4, 4, 5, 6}))
}

func validPartition(nums []int) bool {
	memo := make(map[int]bool)
	return dfs(nums, 0, &memo)
}

// Discuss in three different situations
func dfs(nums []int, curIndex int, memo *map[int]bool) bool {
	if curIndex >= len(nums) {
		return true
	}
	if _, ok := (*memo)[curIndex]; ok {
		return (*memo)[curIndex]
	}
	canFind := false
	if curIndex+1 < len(nums) && nums[curIndex] == nums[curIndex+1] {
		canFind = canFind || dfs(nums, curIndex+2, memo)
	}
	if curIndex+2 < len(nums) && nums[curIndex] == nums[curIndex+1] && nums[curIndex] == nums[curIndex+2] {
		canFind = canFind || dfs(nums, curIndex+3, memo)

	}
	if curIndex+2 < len(nums) && nums[curIndex]+1 == nums[curIndex+1] && nums[curIndex]+2 == nums[curIndex+2] {
		canFind = canFind || dfs(nums, curIndex+3, memo)

	}
	(*memo)[curIndex] = canFind
	return canFind
}
