package main

import (
	"fmt"
	"math"
)

func main() {
	tests := []struct {
		ring string
		key  string
		want int
	}{
		{"godding", "gd", 4},
		// {"abcde", "edcba", 10},
		// {"abcde", "abcde", 9},
		// {"abcde", "eabcd", 10},
	}

	for _, tt := range tests {
		got := findRotateSteps(tt.ring, tt.key)
		if got != tt.want {
			fmt.Printf("findRotateSteps(%v, %v) = %v, want %v。", tt.ring, tt.key, got, tt.want)
		}
	}
}

// dp[i][j] 表示从 ring 的第 j 个字符开始，拼写出 key 的前 i 个字符所需要的最少步数。
func findRotateSteps(ring string, key string) int {
	n := len(ring)
	m := len(key)

	pos := make([][]int, 26)

	// 将环转化为数组
	// 为什么是二维数组？
	// 会存在相同字符，每个字符的步长其实不一样的，这些都需要存下来
	// 后续我们从 a -> b ，只需要去遍历pos[a-‘a’]就可以了
	for i := range ring {
		pos[ring[i]-'a'] = append(pos[ring[i]-'a'], i)
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	// key[0]代表第一个字符
	// key[0]-'a' 代表ring中这个字符所在的index是多少
	// n-p相当于是反着走
	for _, p := range pos[key[0]-'a'] {
		dp[0][p] = min(p, n-p) + 1
	}

	for i := 1; i < m; i++ {
		for _, j := range pos[key[i]-'a'] {
			for _, k := range pos[key[i-1]-'a'] {
				// j代表当前字符的，k代表上一个字符
				// 我们知道了当前字符在哪些位置，上一个字符在哪个位置
				// 两层for循环就能遍历出所有组合了
				// 例如 j:[3,6] k:[5,9]
				// 那么就有 [3,5],[3,9],[6,5],[6,9]四种组合
				// dp[i-1][k]代表匹配到key[i-1]的最少步数
				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), n-abs(j-k))+1)
			}
		}
	}

	return min(dp[m-1]...)
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
