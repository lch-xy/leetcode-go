package main

import "fmt"

func main() {
	testCases := []struct {
		board    [][]byte
		expected [][]byte
	}{
		{
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expected: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			},
		},
	}

	for i, tc := range testCases {
		result := solveSudoku(tc.board)
		if !compareSudoku(result, tc.expected) {
			fmt.Printf("Test case %d failed\n", i+1)
		}
	}
}

// use backtracking
func solveSudoku(board [][]byte) [][]byte {
	dfs(&board)
	return board
}

func dfs(board *[][]byte) bool {
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[0]); j++ {
			// replave '.' by [1,9]
			if string((*board)[i][j]) == "." {
				for k := 1; k <= 9; k++ {
					// check if valid
					if !isValid(board, i, j, k) {
						continue
					}
					(*board)[i][j] = byte('0' + k)
					// pruning , while find a result , return immediately
					if dfs(board) {
						return true
					}
					(*board)[i][j] = byte('.')
				}
				return false
			}
		}
	}
	return true
}

func isValid(board *[][]byte, row, col, k int) bool {
	// check col
	for i := 0; i < len(*board); i++ {
		if ((*board)[i][col]) == byte('0'+k) {
			return false
		}
	}
	// check row
	for i := 0; i < len((*board)[0]); i++ {
		if ((*board)[row][i]) == byte('0'+k) {
			return false
		}
	}
	startRow := row / 3 * 3
	startCol := col / 3 * 3
	// check 3x3 rectangle
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*board)[startRow+i][startCol+j] == byte('0'+k) {
				return false
			}
		}
	}
	return true
}

// Compare two Sudoku boards
func compareSudoku(board1, board2 [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board1[i][j] != board2[i][j] {
				return false
			}
		}
	}
	return true
}
