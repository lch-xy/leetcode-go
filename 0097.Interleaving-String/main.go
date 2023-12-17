package main

import "strconv"

func main() {

	println(isInterleave("qwertyuiop", "asdfghjkl", "qawsedrftyughjikolp"))
	println(isInterleave("a", "", "a"))

	println("================================================================>")

	println(isInterleaveRecursion("qwertyuiop", "asdfghjkl", "qawsedrftyughjikolp"))
	println(isInterleaveRecursion("a", "", "a"))

}

func isInterleaveRecursion(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	cache := map[string]string{}
	return helper(s1, 0, s2, 0, s3, 0, cache)
}

func helper(s1 string, i int, s2 string, j int, s3 string, k int, cache map[string]string) bool {
	// 将i和j的关系映射到key里就可以，随便什么方法都行
	key := strconv.Itoa(i*len(s3) + j)
	_, ok := cache[key]
	if ok {
		return false
	}
	if i == len(s1) {
		// 如果匹配上就返回true 匹配不是就返回false
		return s2[j:] == s3[k:]
	}
	if j == len(s2) {
		return s1[i:] == s3[k:]
	}
	// 看s3的当前位置能都和s1、s2的任意位置匹配上
	if (s1[i] == s3[k] && helper(s1, i+1, s2, j, s3, k+1, cache)) ||
		(s2[j] == s3[k] && helper(s1, i, s2, j+1, s3, k+1, cache)) {
		return true
	}
	// 如果没匹配上 下次遇到直接返回
	cache[key] = key
	return false
}

// dp[i][j] 以s1[:i-1]和s2[:j-1]能否组成s3[:i+j]
// 每次只对比一个字符，要么从dp数组上面而来，要么从dp数组左边而来
// 从上面而来，j长度不变，只增加了i的值。dp[i-1][j] && s1[i-1] == s3[i-1+j]
// 从左面而来，i长度不变，只增加了j的值。dp[i][j-1] && s2[j-1] == s3[i-1+j]
func isInterleave(s1 string, s2 string, s3 string) bool {
	x := len(s1)
	y := len(s2)
	if x+y != len(s3) {
		return false
	}
	dp := make([][]bool, x+1)
	for i := 0; i <= x; i++ {
		dp[i] = make([]bool, y+1)
	}
	dp[0][0] = true
	for i := 1; i <= x; i++ {
		if dp[i-1][0] && s1[i-1] == s3[i-1] {
			dp[i][0] = true
		}
	}
	for i := 1; i <= y; i++ {
		if dp[0][i-1] && s2[i-1] == s3[i-1] {
			dp[0][i] = true
		}
	}
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if (dp[i-1][j] && s1[i-1] == s3[i-1+j]) || (dp[i][j-1] && s2[j-1] == s3[i-1+j]) {
				dp[i][j] = true
			}
		}
	}
	return dp[x][y]
}
