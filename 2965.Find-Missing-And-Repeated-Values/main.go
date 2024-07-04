package main

import "fmt"

func main() {
	fmt.Println(findMissingAndRepeatedValues([][]int{{1, 3}, {2, 2}}))
}

func findMissingAndRepeatedValues(grid [][]int) []int {
	counter := make([]int, 2501)
	row := len(grid)
	col := len(grid[0])

	res := make([]int, 2)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			counter[grid[i][j]]++
			if counter[grid[i][j]] > 1 {
				res[0] = grid[i][j]
			}
		}
	}
	for i := 1; i <= 2500; i++ {
		if counter[i] == 0 {
			res[1] = i
			break
		}
	}
	return res
}
