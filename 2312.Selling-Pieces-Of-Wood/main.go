package main

import "fmt"

func main() {
	testCases := []struct {
		name   string
		m, n   int
		prices [][]int
		want   int64
	}{
		{
			name:   "Example 1",
			m:      4,
			n:      6,
			prices: [][]int{{3, 2, 10}, {1, 4, 2}, {4, 1, 3}},
			want:   32, // 应替换为计算得出的正确结果值
		},
		// TODO: 添加更多测试用例
	}

	// 遍历测试用例，运行测试
	for _, tc := range testCases {
		got := sellingWood(tc.m, tc.n, tc.prices)
		if got != tc.want {
			fmt.Printf("%s: MaxProfit(%v, %v, %v) = %v, want %v", tc.name, tc.m, tc.n, tc.prices, got, tc.want)
		}
	}
}

func sellingWood(m int, n int, prices [][]int) int64 {
	cache := make([][]int64, m+1)
	for i := range cache {
		cache[i] = make([]int64, n+1)
		for j := range cache[i] {
			// -1 to indicate uncaculated subproblems
			cache[i][j] = -1
		}
	}

	tablePrice := make([][]int64, m+1)
	for i := range tablePrice {
		tablePrice[i] = make([]int64, n+1)
	}

	for _, price := range prices {
		tablePrice[price[0]][price[1]] = int64(price[2])
	}

	return dfs(cache, tablePrice, m, n)
}

// dfs to caculate the maximum price  with given m x n size
func dfs(cache [][]int64, tablePrice [][]int64, h, w int) int64 {
	if cache[h][w] != -1 {
		return cache[h][w]
	}

	maxPrice := tablePrice[h][w]

	// try every possible split point vertically
	for i := 1; i <= h/2; i++ {
		maxPrice = max(maxPrice, dfs(cache, tablePrice, i, w)+dfs(cache, tablePrice, h-i, w))
	}

	// try every possible split point horizontally
	for i := 1; i <= w/2; i++ {
		maxPrice = max(maxPrice, dfs(cache, tablePrice, h, i)+dfs(cache, tablePrice, h, w-i))
	}

	cache[h][w] = maxPrice

	return maxPrice
}
