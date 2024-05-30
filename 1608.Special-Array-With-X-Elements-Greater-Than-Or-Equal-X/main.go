package main

// one step : count the number of occurrences of each number
// two step : traverse arrays
func specialArray(nums []int) int {
	n := len(nums)
	counter := make([]int, 1001)
	for i := 0; i < n; i++ {
		counter[nums[i]]++
	}
	total := n
	for i := 0; i <= 1000; i++ {
		if i == total {
			return i
		}
		total -= counter[i]
	}
	return -1
}
