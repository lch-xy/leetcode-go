package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Printf("cherryPickup([][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}): %v\n", cherryPickup([][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}))
	fmt.Printf("cherryPickup([][]int{{1}}): %v\n", cherryPickup([][]int{{1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1}}))
}

// question : (0,0) -> (n-1,n-1) -> (0,0)
// intuitively , we usually use two dimensions array to solve this problem
// although , we can pick up the maxmium cherries from (0,0) to (n-1,n-1)
// we can not pick up the maxmium cherries from (0,0) to (n-1,n-1) and (n-1,n-1) to (0,0)
// like this case :
// 0  0  1  1  1  0  0
// 0  0  0  0  1  0  0
// 0  0  0  0  1  0  1
// 1  0  0  0  1  0  0
// 0  0  0  0  1  0  0
// 0  0  0  0  1  0  0
// 0  0  0  0  1  1  1
// (0,0) -> (n-1,n-1) : (0,0) -> (4,0) -> (4,6) -> (6,6) we can pick up 11 cherries
// (n-1,n-1) -> (0,0) : (6,6) -> (6,0) -> (0,0) we can pick up 1 cherries
// according to the previous writing , we can pick up 12 cherries
// but if we use this path (0,0) -> (4,2) -> (6,2) -> (6,6) -> (4,6) -> (4,3) -> (0,3) -> (0,0) we can pick up 13 cherries

// we can transfer angle to see this question
// transfer the question to two person start from (n-1,n-1) to (0,0) independently, and they have same steps
// dp[x1,y1,x2] : use three dimensions to represent the two person's position (x1,y1) and (x2,x1+y1-x2)
func cherryPickup(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	memo := make([][][]int, row)
	for i := range memo {
		memo[i] = make([][]int, col)
		for j := range memo[i] {
			memo[i][j] = make([]int, row)
			for k := range memo[i][j] {
				memo[i][j][k] = math.MinInt32
			}
		}
	}

	return max(0, dp(row-1, col-1, row-1, &memo, &grid))
}

func dp(x1, y1, x2 int, memo *[][][]int, grid *[][]int) int {
	y2 := x1 + y1 - x2
	if x1 < 0 || y1 < 0 || x2 < 0 || y2 < 0 {
		return -1
	}
	if (*grid)[x1][y1] < 0 || (*grid)[x2][y2] < 0 {
		return -1
	}
	if x1 == 0 && y1 == 0 {
		return (*grid)[x1][y1]
	}

	if (*memo)[x1][y1][x2] != math.MinInt32 {
		return (*memo)[x1][y1][x2]
	}

	(*memo)[x1][y1][x2] = max(dp(x1-1, y1, x2-1, memo, grid), dp(x1-1, y1, x2, memo, grid), dp(x1, y1-1, x2-1, memo, grid), dp(x1, y1-1, x2, memo, grid))

	// if result >= 0 means the path don't have thorns
	if (*memo)[x1][y1][x2] >= 0 {
		(*memo)[x1][y1][x2] += (*grid)[x1][y1]
		if x1 != x2 {
			(*memo)[x1][y1][x2] += (*grid)[x2][y2]
		}
	}
	return (*memo)[x1][y1][x2]
}

// time out
// func cherryPickup(grid [][]int) int {
// 	res := 0
// 	goToDestnation(&grid, 0, 0, 0, &res)
// 	return res
// }

// func goToDestnation(grid *[][]int, row, col, count int, res *int) {
// 	if row > len((*grid))-1 || col > len((*grid)[0])-1 || (*grid)[row][col] == -1 {
// 		return
// 	}
// 	if (*grid)[row][col] == -1 {
// 		return
// 	}
// 	temp := (*grid)[row][col]
// 	if temp == 1 {
// 		count++
// 		(*grid)[row][col] = 0
// 	}
// 	if row == len((*grid))-1 && col == len((*grid)[0])-1 {
// 		goBackOrigin(grid, row, col, count, res)
// 	}
// 	goToDestnation(grid, row+1, col, count, res)
// 	goToDestnation(grid, row, col+1, count, res)
// 	(*grid)[row][col] = temp
// }

// func goBackOrigin(grid *[][]int, row, col, count int, res *int) {
// 	if row < 0 || col < 0 {
// 		return
// 	}
// 	if (*grid)[row][col] == -1 {
// 		return
// 	}
// 	if row == 0 && col == 0 {
// 		*res = max(*res, count)
// 		return
// 	}
// 	temp := (*grid)[row][col]
// 	if temp == 1 {
// 		count++
// 		(*grid)[row][col] = 0
// 	}
// 	goBackOrigin(grid, row-1, col, count, res)
// 	goBackOrigin(grid, row, col-1, count, res)
// 	(*grid)[row][col] = temp
// }
