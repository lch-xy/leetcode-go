package main

import (
	"fmt"
)

func main() {
	fmt.Printf("countTestedDevices([]int{1, 1, 2, 1, 3}): %v\n", countTestedDevices([]int{1, 1, 2, 1, 3}))
	fmt.Printf("countTestedDevices([]int{0,1,2}): %v\n", countTestedDevices([]int{0, 1, 2}))
}

func countTestedDevices(batteryPercentages []int) int {
	cnt := 0
	for i := 0; i < len(batteryPercentages); i++ {
		batteryPercentages[i] -= cnt
		if batteryPercentages[i] > 0 {
			batteryPercentages[i]++
			cnt++
		}
	}
	return cnt
}
