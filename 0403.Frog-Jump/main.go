package main

import "fmt"

func main() {
	testCases := []struct {
		name   string
		stones []int
		want   bool
	}{
		{"normal", []int{0, 1, 3, 5, 6, 8, 12, 17}, true},
		{"normal", []int{0, 1, 2, 3, 4, 8, 9, 11}, false},
		{"edge", []int{0, 1}, true},
		{"edge", []int{0, 2}, false},
		{"exception", []int{}, false},
	}

	for _, tc := range testCases {
		got := canCross(tc.stones)
		if got != tc.want {
			fmt.Printf("canCross(%v) = %v; want %v", tc.stones, got, tc.want)
			fmt.Println("")
		}
	}
}

// dp[i]里存储的是到跳到当前位置后，所用的步数是多少
// 那我往后跳的时候，就可以根据当前的步数进行三种选择了
// 也可以使用[]map[int]bool来进行优化，因为contains函数查询是O(n)的时间复杂度
func canCross(stones []int) bool {
	if len(stones) == 0 {
		return false
	}
	size := len(stones)

	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, 0)
	}

	dp[0] = []int{0}
	for i := 1; i < size; i++ {
		// 需要遍历之前所有的情况，因为可以从任何一个石头跳过来
		for j := 0; j < i; j++ {
			for _, v := range dp[j] {
				if stones[j]+v+1 == stones[i] {
					if !contains(dp[i], v+1) {
						dp[i] = append(dp[i], v+1)
					}
				} else if stones[j]+v == stones[i] {
					if !contains(dp[i], v) {
						dp[i] = append(dp[i], v)
					}
				} else if stones[j]+v-1 == stones[i] {
					if !contains(dp[i], v-1) {
						dp[i] = append(dp[i], v-1)
					}
				}
			}
		}
	}
	return len(dp[size-1]) > 0
}

func contains(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}
