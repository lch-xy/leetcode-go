package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("findColumnWidth([][]int{{1}, {222}, {33}}): %v\n", findColumnWidth([][]int{{1}, {222}, {33}}))
}

func findColumnWidth(grid [][]int) []int {
	row := len(grid)
	col := len(grid[0])
	res := make([]int, 0)
	for i := 0; i < col; i++ {
		lenght := 0
		for j := 0; j < row; j++ {
			if lenght < len(strconv.Itoa(grid[j][i])) {
				lenght = len(strconv.Itoa(grid[j][i]))
			}
		}
		res = append(res, lenght)
	}
	return res
}
