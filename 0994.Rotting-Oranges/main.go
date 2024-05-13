package main

import "fmt"

func main() {
	type TestCase struct {
		grid   [][]int
		output int
	}
	testCases := []TestCase{
		{
			grid:   [][]int{{1, 2}},
			output: 1,
		},
		{
			grid:   [][]int{{0}},
			output: 0,
		},
		{
			grid:   [][]int{{1}},
			output: -1,
		},
		{
			grid:   [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			output: 4,
		},
		{
			grid:   [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			output: -1,
		},
		{
			grid:   [][]int{{0, 2}},
			output: 0,
		},
	}

	for i, tc := range testCases {
		result := orangesRotting(tc.grid)
		if result == tc.output {
			fmt.Printf("Test case %d passed\n", i+1)
		} else {
			fmt.Printf("Test case %d failed: expected %d but got %d\n", i+1, tc.output, result)
		}
	}
}

// use bfs 
func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	queue := [][]int{}
	freshOranges := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				freshOranges++
			}
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
				visited[i][j] = true
			}
		}
	}
	if len(queue) == 0 && freshOranges == 0 {
		return 0
	}

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	step := -1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curNode := queue[0]
			queue = queue[1:]
			curRow := curNode[0]
			curCol := curNode[1]
			for _, v := range directions {
				newRow := curRow + v[0]
				newCol := curCol + v[1]
				if newRow < 0 || newRow >= m || newCol < 0 || newCol >= n || visited[newRow][newCol] {
					continue
				}
				if grid[newRow][newCol] == 1 {
					queue = append(queue, []int{newRow, newCol})
					visited[newRow][newCol] = true
					freshOranges--
				}
			}
		}
		step++
	}
	if freshOranges > 0 {
		return -1
	}
	return step
}
