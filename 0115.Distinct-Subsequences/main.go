package main

import (
	"fmt"
)

func main() {

	println(numDistinct("rabbbit", "rabbit"))
	println(numDistinct("babgbag", "bag"))

	println("=================================================>")

	println(numDistinctDp("rabbbit", "rabbit"))
	println(numDistinctDp("babgbag", "bag"))
}

// # Ø r a b b b i t
// Ø 1 1 1 1 1 1 1 1
// r 0 1 1 1 1 1 1 1
// a 0 0 1 1 1 1 1 1
// b 0 0 0 1 2 3 3 3
// b 0 0 0 0 1 3 3 3
// i 0 0 0 0 0 0 3 3
// t 0 0 0 0 0 0 0 3
func numDistinctDp(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return 0
	}
	x := len(s)
	y := len(t)
	dp := make([][]int, x+1)
	for i := 0; i <= x; i++ {
		dp[i] = make([]int, y+1)
		dp[i][0] = 1
	}
	// 循环以j为外循环，因为我们是要去匹配t的元素
	// s[i-1]是可以不与t[j-1]匹配上，继续往下匹配
	for j := 1; j <= y; j++ {
		for i := j; i <= x; i++ {
			if s[i-1] == t[j-1] {
				// 如果当前两个字符串相当，那么s这个字符可以选择也可以不选择
				// 选择就是从dp[i-1][j-1]转移而来，我们选择了这个几点，那么方案就继承了dp[i-1][j-1]
				// 不选择就是dp[i-1][j]
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				// 不相等i就不能选择这个节点，
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[x][y]
}

// 递归+记忆化搜索 但是还是超时无法ac
func numDistinct(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return 0
	}
	cache := make(map[string]int)
	return helper(s, 0, t, 0, cache)
}

func helper(s string, i int, t string, j int, cache map[string]int) int {
	if j == len(t) {
		return 1
	}
	key := fmt.Sprintf("%d-%d", i, j)
	if val, found := cache[key]; found {
		return val
	}
	result := 0
	for k := i; k < len(s); k++ {
		if s[k] == t[j] {
			result += helper(s, k+1, t, j+1, cache)
		}
	}
	cache[key] = result
	return result
}
