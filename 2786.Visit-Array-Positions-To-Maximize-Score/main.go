package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxScore([]int{2, 3, 6, 1, 9, 2}, 5))
}

// cache[0]: the maximum even number
// cache[1]: the maximum odd number
// we can move to this position from the previous maximum even number or odd number
func maxScore(nums []int, x int) int64 {
	n := len(nums)
	cache := []int{math.MinInt32, math.MinInt32}
	cache[nums[0]%2] = nums[0]

	for i := 1; i < n; i++ {
		parity := nums[i] % 2
		cur := max(cache[parity]+nums[i], cache[1-parity]+nums[i]-x)
		cache[parity] = max(cur, cache[parity])
	}
	return max(int64(cache[0]), int64(cache[1]))
}
