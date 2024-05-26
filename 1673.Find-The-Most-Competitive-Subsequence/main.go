package main

import "fmt"

func main() {
	fmt.Printf("mostCompetitive([]int{3, 5, 2, 6}, 2): %v\n", mostCompetitive([]int{3, 5, 2, 6}, 2))
}

// monotone-stack 
func mostCompetitive(nums []int, k int) []int {
	stack := make([]int, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		// len(stack)+n-i > k : means only n-i numbers remain , add all the remaining number into stack will still not exceded k numbers
		for len(stack) != 0 && stack[len(stack)-1] >= nums[i] && len(stack)+n-i > k {
			stack = stack[:len(stack)-1]
		}
		// only need k nunbers , so lager numbers not need to enter the stack
		if len(stack) < k {
			stack = append(stack, nums[i])
		}
	}
	return stack
}
