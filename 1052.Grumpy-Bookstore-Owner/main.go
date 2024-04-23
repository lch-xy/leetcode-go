package main

import "fmt"

func main() {
	fmt.Printf("maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3): %v\n", maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
}

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	length := len(grumpy)
	cnt := 0
	for i := 0; i < length; i++ {
		if grumpy[i] == 0 {
			cnt += customers[i]
		}
	}
	total := 0
	if length <= minutes {
		for i := 0; i < length; i++ {
			if grumpy[i] == 1 {
				total += customers[i]
			}
		}
	} else {
		for i := 0; i <= length-minutes; i++ {
			tmp := 0
			for j := 0; j < minutes; j++ {
				if grumpy[i+j] == 1 {
					tmp += customers[i+j]
				}
			}
			if tmp > total {
				total = tmp
			}
		}
	}
	return cnt + total
}
