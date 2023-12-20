package main

func main() {
	println(minCut("aab"))
}

// dp[i]：s[0,i]的子串，拆分成回文串的最小切割数
func minCut(s string) int {
	len := len(s)
	if len == 0 {
		return 0
	}
	// s[j,i]是否是回文串
	isVaild := make([][]bool, len)
	for i := 0; i < len; i++ {
		isVaild[i] = make([]bool, len)
	}
	dp := make([]int, len)
	for i := 0; i < len; i++ {
		// 最多就是切i次
		dp[i] = i
		// 开始从[0,i]这个区间内 一刀刀的切
		for j := 0; j <= i; j++ {
			// i-j<2：说明1个字符和2个字符的情况
			// isVaild[j+1][i-1]：主要是判断3个以上字符的情况
			if s[i] == s[j] && (i-j < 2 || isVaild[j+1][i-1]) {
				// 进来说明s[j,i]已经是回文了
				isVaild[j][i] = true
				// 当j==0时，说明[0,i]之间是回文，那么一刀都不用切
				if j == 0 {
					dp[i] = 0
				} else {
					// dp[j-1]表示从位置0到位置j-1的最小切割次数，+1表示与前面分割的那一刀
					dp[i] = min(dp[i], dp[j-1]+1)
				}
			}
		}
	}
	return dp[len-1]
}
