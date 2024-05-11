package main

import "fmt"

func main() {
	fmt.Printf("garbageCollection([]string{\"G\", \"P\", \"GP\", \"GG\"}, []int{2, 4, 3}): %v\n", garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}))
}

// simulate the driving path of garbage rucks
// the garbage trucks will only drive when there is such garbage at the next station
func garbageCollection(garbage []string, travel []int) int {
	garbageM := 0
	garbageP := 0
	garbageG := 0
	driveTime := 0
	handleTime := 0
	for i := 0; i < len(garbage); i++ {
		cntM := caculate(garbage[i], "M")
		cntG := caculate(garbage[i], "G")
		cntP := caculate(garbage[i], "P")
		if i != 0 {
			if cntM > 0 {
				for j := garbageM; j < i; j++ {
					driveTime += travel[j]
				}
				garbageM = i
			}
			if cntG > 0 {
				for j := garbageG; j < i; j++ {
					driveTime += travel[j]
				}
				garbageG = i
			}
			if cntP > 0 {
				for j := garbageP; j < i; j++ {
					driveTime += travel[j]
				}
				garbageP = i
			}
		}
		handleTime = handleTime + cntM + cntP + cntG
	}
	return driveTime + handleTime
}

// caculate count of target string
func caculate(source string, target string) int {
	cnt := 0
	for _, v := range source {
		if string(v) == target {
			cnt++
		}
	}
	return cnt
}
