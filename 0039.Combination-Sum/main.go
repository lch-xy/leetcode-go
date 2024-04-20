package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	helper(candidates, target, 0, 0, []int{}, &res)
	return res
}

func helper(candidates []int, target, sum, index int, curList []int, res *[][]int) {
	if sum > target {
		return
	}
	if sum == target {
		newList := make([]int, len(curList))
		copy(newList, curList)
		*res = append(*res, newList)
		return
	}
	for i := index; i < len(candidates); i++ {
		helper(candidates, target, sum+candidates[i], i, append(curList, candidates[i]), res)
	}
}
