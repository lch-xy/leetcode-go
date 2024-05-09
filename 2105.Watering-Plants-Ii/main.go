package main

import "fmt"

func main() {
	fmt.Printf("minimumRefill([]int{2, 2, 3, 3}, 5, 5): %v\n", minimumRefill([]int{2, 2, 3, 3}, 5, 5))
}

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	curA := capacityA
	curB := capacityB
	left, right := 0, len(plants)-1
	total := 0
	for left < right {
		if curA < plants[left] {
			total++
			curA = capacityA
		}
		curA -= plants[left]
		if curB < plants[right] {
			total++
			curB = capacityB
		}
		curB -= plants[right]
		left++
		right--
	}
	if left == right {
		if curA < plants[left] && curB < plants[right] {
			total++
		}
	}
	return total
}
