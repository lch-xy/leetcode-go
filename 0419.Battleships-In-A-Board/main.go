package main

func countBattleships(board [][]byte) int {
	row := len(board)
	col := len(board[0])
	cnt := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'X' {
				if (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
					cnt++
				}
			}
		}
	}
	return cnt
}
