package main

import (
	"fmt"
)

func main() {
	fmt.Printf("numberOfWeeks([]int{1, 2, 3}): %v\n", numberOfWeeks([]int{1, 2, 3}))
	fmt.Printf("numberOfWeeks([]int{5, 2, 1}): %v\n", numberOfWeeks([]int{5, 2, 1}))
	fmt.Printf("numberOfWeeks([]int{5, 7, 5, 7, 9, 7}): %v\n", numberOfWeeks([]int{5, 7, 5, 7, 9, 7}))
}

// if max(milestones) > sum(milestones) - max(milestones) + 1
// we will inevitably end up violating the rule about not working on the same project for two consecutive weeks.
// if max(milestones) <= sum(milestones) - max(milestones) + 1
// we can always alternate between the project with the most milestones and the other projects.
func numberOfWeeks(milestones []int) int64 {
	n := len(milestones)
	maxNumber := -1
	total := 0
	for i := 0; i < n; i++ {
		total += milestones[i]
		maxNumber = max(maxNumber, milestones[i])
	}

	rest := total - maxNumber

	if maxNumber > rest+1 {
		return int64(rest*2 + 1)
	}
	return int64(total)

}
