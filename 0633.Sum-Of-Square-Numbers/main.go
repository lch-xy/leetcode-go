package main

import "math"

func judgeSquareSum(c int) bool {
	d := math.Sqrt(float64(c))
	end := int(math.Ceil(d))
	start := 0
	for start < end {
		if start*start+end*end > c {
			end--
			continue
		} else if start*start+end*end < c {
			start++
			continue
		} else {
			return true
		}
	}
	return false
}
