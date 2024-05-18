package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("maxDivScore([]int{4, 7, 9, 3, 9}, []int{5, 2, 3}): %v\n", maxDivScore([]int{4, 7, 9, 3, 9}, []int{5, 2, 3}))
}

func maxDivScore(nums []int, divisors []int) int {
	resDivisor := math.MaxInt32
	resCnt := 0
	for i := 0; i < len(divisors); i++ {
		curDivisor := divisors[i]
		curCnt := 0
		for j := 0; j < len(nums); j++ {
			if nums[j]%curDivisor == 0 {
				curCnt++
			}
		}
		if resCnt == curCnt && resDivisor > curDivisor {
			resDivisor = curDivisor
		}
		if resCnt < curCnt {
			resCnt = curCnt
			resDivisor = curDivisor
		}
	}
	return resDivisor
}
