package main

import "fmt"

func main() {
	fmt.Printf("findPeaks([]int{1, 4, 3, 8, 5}): %v\n", findPeaks([]int{1, 4, 3, 8, 5}))
}

func findPeaks(mountain []int) []int {
	res := make([]int, 0)
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			res = append(res, i)
		}
	}
	return res
}
