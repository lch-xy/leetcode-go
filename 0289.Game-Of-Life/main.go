package main

import (
	"fmt"
	"reflect"
)

func main() {
	testCases := []struct {
		board    [][]int
		expected [][]int
	}{
		{
			board: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
	}

	for i, tc := range testCases {
		gameOfLife(tc.board)
		if !reflect.DeepEqual(tc.board, tc.expected) {
			fmt.Printf("Test case %d failed:\nExpected output:\n%v\nActual output:\n%v\n", i+1, tc.expected, tc.board)
		}
	}
}

// 2:Survive in the last stage, die in the next stage
// 3:Die in the last stage, survive in the next stage
func gameOfLife(board [][]int) {
	dics := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	row := len(board)
	col := len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			cnt := caculate(board, dics, row, col, i, j)
			if cnt < 2 && board[i][j] == 1 {
				// if two or three live neighbors, the cell lives on to the next generation
				board[i][j] = 2
			} else if (cnt >= 2 && cnt <= 3) && board[i][j] == 1 {
				// if two or three live cells are around the live cell, the live cell is still alive
				// nothing to do
			} else if cnt > 3 && board[i][j] == 1 {
				// if more than three live cells are around the live cell, the live cell will die next generation
				board[i][j] = 2
			} else if cnt == 3 && board[i][j] == 0 {
				// if three live cells surround a dead cell, the dead cell will be revived
				board[i][j] = 3
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}

}

func caculate(board, dics [][]int, row, col, curRow, curCol int) int {
	cnt := 0
	for _, dic := range dics {
		newRow := curRow + dic[0]
		newCol := curCol + dic[1]
		if newRow < 0 || newRow >= row || newCol < 0 || newCol >= col {
			continue
		}
		if board[newRow][newCol] == 1 || board[newRow][newCol] == 2 {
			cnt++
		}
	}
	return cnt
}
