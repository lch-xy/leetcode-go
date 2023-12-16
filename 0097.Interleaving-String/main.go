package main

func main() {

	println(isInterleave("qwertyuiop", "asdfghjkl", "qawsedrftyughjikolp"))
	println(isInterleave("a", "", "a"))
	
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
