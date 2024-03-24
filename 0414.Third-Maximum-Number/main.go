package main

import "math"

// use math.MinInt64 instead of math.MinInt32
func thirdMax(nums []int) int {
	a1, a2, a3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num == a1 || num == a2 || num == a3 {
			continue
		}
		if num > a1 {
			a3 = a2
			a2 = a1
			a1 = num
		} else if num > a2 {
			a3 = a2
			a2 = num
		} else if num > a3 {
			a3 = num
		}
	}

	if a3 == math.MinInt64 {
		return a1
	}
	return a3
}
