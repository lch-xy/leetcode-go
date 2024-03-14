package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	start := nums[0]
	end := nums[0]
	for i := 1; i < len(nums); i++ {
		if end+1 == nums[i] {
			end = nums[i]
			continue
		}
		if start == end {
			res = append(res, fmt.Sprint(start))
		} else {
			res = append(res, fmt.Sprint(start, "->", end))
		}
		start = nums[i]
		end = nums[i]
	}
	// to add the last interval
	if start == end {
		res = append(res, fmt.Sprint(start))
	} else {
		res = append(res, fmt.Sprint(start, "->", end))
	}
	return res
}
