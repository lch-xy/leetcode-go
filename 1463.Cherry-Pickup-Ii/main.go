package main

func cherryPickup(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	cache := make([][][]int, row)
	for i := range cache {
		cache[i] = make([][]int, col)
		for j := range cache[i] {
			cache[i][j] = make([]int, col)
			for k := range cache[i][j] {
				cache[i][j][k] = -1
			}
		}
	}

	return max(0, dp(0, 0, col-1, row, col, &cache, grid))
}

func dp(x1, y1, y2, row, col int, cache *[][][]int, grid [][]int) int {
	if x1 >= row || y1 >= col || y2 >= col || y1 < 0 || y2 < 0 {
		return 0
	}
	if (*cache)[x1][y1][y2] != -1 {
		return (*cache)[x1][y1][y2]
	}
	curCherries := 0
	if y1 != y2 {
		curCherries += grid[x1][y1] + grid[x1][y2]
	} else {
		curCherries += grid[x1][y1]
	}

	maxCherries := 0
	// find maxCherries in nine situations
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			maxCherries = max(maxCherries, curCherries+dp(x1+1, y1+i, y2+j, row, col, cache, grid))
		}
	}
	(*cache)[x1][y1][y2] = maxCherries
	return maxCherries
}
