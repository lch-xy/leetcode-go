package main

import "fmt"

func main() {
	type testCase struct {
		grid [][]int
		want int
	}
	tests := []testCase{
		{
			grid: [][]int{
				{0, 0, 1},
				{1, 0, 1},
				{0, 0, 0},
			},
			want: 1,
		},
		{
			grid: [][]int{
				{0, 0, 0},
				{1, 0, 0},
				{1, 1, 0},
			},
			want: 2,
		},
	}

	for _, tc := range tests {
		got := findChampion(tc.grid)
		fmt.Printf("grid: %v, want: %v, got: %v\n", tc.grid, tc.want, got)
	}
}

func findChampion(grid [][]int) int {
	row := len(grid)
	startChampion := 0
	loop := true
	// to find champion level by level
	for loop {
		temp := startChampion
		for i := 0; i < row; i++ {
			if startChampion == i {
				continue
			}
			if grid[startChampion][i] != 1 {
				startChampion = i
				break
			}
		}
		// no team is stronger than startChampion , break loop
		if temp == startChampion {
			loop = false
		}
	}
	return startChampion
}
