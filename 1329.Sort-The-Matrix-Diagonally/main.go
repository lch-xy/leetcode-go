package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	type TestCase struct {
		mat    [][]int
		output [][]int
	}
	testCases := []TestCase{
		{
			mat: [][]int{
				{3, 3, 1, 1},
				{2, 2, 1, 2},
				{1, 1, 1, 2},
			},
			output: [][]int{
				{1, 1, 1, 1},
				{1, 2, 2, 2},
				{1, 2, 3, 3},
			},
		},
	}

	for i, tc := range testCases {
		result := diagonalSort(tc.mat)
		if reflect.DeepEqual(result, tc.output) {
			fmt.Printf("Test case %d passed\n", i+1)
		} else {
			fmt.Printf("Test case %d failed\n", i+1)
		}
	}
}

func diagonalSort(mat [][]int) [][]int {
	row := len(mat)
	col := len(mat[0])

	list := make([]int, 0)
	for i := 0; i < row; i++ {
		addAll(i, 0, row, col, &list, mat)
		sort.Ints(list)
		fillGraph(i, 0, row, col, &list, &mat)
	}

	for j := 1; j < col; j++ {
		addAll(0, j, row, col, &list, mat)
		sort.Ints(list)
		fillGraph(0, j, row, col, &list, &mat)
	}
	return mat
}

func fillGraph(curRow, curCol, row, col int, list *[]int, mat *[][]int) {
	if curRow >= row || curCol >= col {
		return
	}
	curNumber := (*list)[0]
	*list = (*list)[1:]
	(*mat)[curRow][curCol] = curNumber
	fillGraph(curRow+1, curCol+1, row, col, list, mat)
}

func addAll(curRow, curCol, row, col int, list *[]int, mat [][]int) {
	if curRow >= row || curCol >= col {
		return
	}
	*list = append(*list, mat[curRow][curCol])
	addAll(curRow+1, curCol+1, row, col, list, mat)
}
