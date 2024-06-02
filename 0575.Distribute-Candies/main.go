package main

import "sort"

func distributeCandies(candyType []int) int {
	sort.Ints(candyType)
	n := len(candyType)
	cnt := 1
	for i := 1; i < n; i++ {
		if candyType[i] != candyType[i-1] {
			cnt++
		}
	}
	return min(n/2, cnt)
}
