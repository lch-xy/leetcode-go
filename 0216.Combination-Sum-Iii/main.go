package main

import (
	"fmt"
	"reflect"
)

func main() {
	tests := []struct {
		k    int
		n    int
		want [][]int
	}{
		// {3, 7, [][]int{{1, 2, 4}}},
		// {3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{4, 24, [][]int{{1, 6, 8, 9}, {2, 5, 8, 9}, {2, 6, 7, 9}, {3, 4, 8, 9}, {3, 5, 7, 9}, {3, 6, 7, 8}, {4, 5, 6, 9}, {4, 5, 7, 8}}},
	}

	for _, tt := range tests {
		got := combinationSum3(tt.k, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("combinationSum3(%v, %v) = %v, want %v", tt.k, tt.n, got, tt.want)
		}
	}
}

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	dfs(&res, make([]int, 0), 1, n, k)
	return res
}

func dfs(res *[][]int, curList []int, index, remain, k int) {
	if remain < 0 || len(curList) > k {
		return
	}
	if remain == 0 && len(curList) == k {
		*res = append(*res, append([]int{}, curList...))
		return
	}
	for i := index; i <= 9; i++ {
		dfs(res, append(curList, i), i+1, remain-i, k)
	}
}
