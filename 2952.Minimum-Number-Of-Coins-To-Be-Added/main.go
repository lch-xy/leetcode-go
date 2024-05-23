package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Printf("minimumAddedCoins([]int{1, 4, 10}, 19): %v\n", minimumAddedCoins([]int{1, 4, 10}, 19))
	fmt.Printf("minimumAddedCoins([]int{15, 1, 12}, 43): %v\n", minimumAddedCoins([]int{15, 1, 12}, 43))
}

// sorted arrays
// iterate arrays , add coin in to targetSum one by one , and compare the two number targetSum and coins[i]
// if coins[i] <= targetSum , means we can piece together the value of targetSum + coins[i]
// because we have already pieced together the value from 0 to targetSum - 1
// if coins[i] > targetSum , means we missing some values from targetSum - 1 to coins[i]
// so we should add coin that the value is targetSum
func minimumAddedCoins(coins []int, target int) int {
	sort.Ints(coins)
	targetSum := 1
	res := 0
	i := 0
	// why is <= not < ?
	// because targetSum default value 1 , only pieced together from 0 to targetSum -1
	for targetSum <= target {
		if i < len(coins) && coins[i] <= targetSum {
			targetSum += coins[i]
			i++
		} else {
			targetSum <<= 1
			res++
		}
	}
	return res
}
