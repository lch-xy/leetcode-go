package main

import "fmt"

func main() {
	fmt.Printf("missingRolls([]int{1, 2, 3, 4}, 6, 4): %v\n", missingRolls([]int{1, 2, 3, 4}, 6, 4))
	fmt.Printf("missingRolls([]int{3, 2, 4, 3}, 4, 2): %v\n", missingRolls([]int{3, 2, 4, 3}, 4, 2))
	fmt.Printf("missingRolls([]int{4,5,6,2,3,6,5,4,6,4,5,1,6,3,1,4,5,5,3,2,3,5,3,2,1,5,4,3,5,1,5}, 4, 40): %v\n", missingRolls([]int{4, 5, 6, 2, 3, 6, 5, 4, 6, 4, 5, 1, 6, 3, 1, 4, 5, 5, 3, 2, 3, 5, 3, 2, 1, 5, 4, 3, 5, 1, 5}, 4, 40))
}

func missingRolls(rolls []int, mean int, n int) []int {
	total := 0
	for i := 0; i < len(rolls); i++ {
		total += rolls[i]
	}
	remain := float64(len(rolls)*mean-total+mean*n) / float64(n)
	if remain > 6 || remain < 1 {
		return []int{}
	}
	res := make([]int, 0)
	decrease := float64(0)
	for i := 0; i < n-1; i++ {
		decrease += remain - float64(int(remain))
		res = append(res, int(remain)+int(decrease))
		total += int(remain) + int(decrease)
		if decrease >= 1 {
			decrease--
		}
	}
	res = append(res, len(rolls)*mean-total+mean*n)
	return res
}
