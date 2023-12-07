package main

func main() {

	println(isMatchByDp("aaaaaaa", "a*"))
	println(isMatch("a", "a*"))
	println(isMatch("ab", "ab*"))
	println(isMatchByDp("", "a*"))
	println(isMatchByDp("a", "a*"))
}

func isMatchByDp(s string, p string) bool {
	x := len(s)
	y := len(p)
	dp := make([][]bool, x+1)
	for i := range dp {
		dp[i] = make([]bool, y+1)
	}
	// 初始化数组
	dp[0][0] = true
	for i := 1; i <= x; i++ {
		dp[i][0] = false
	}
	for i := 1; i <= y; i++ {
		// s="" p="a*"，当前是*，但是可以出现0次，就是想"a*"去掉，等于对比两个空字符串
		if p[i-1] == '*' && dp[0][i-2] {
			dp[0][i] = true
		} else {
			dp[0][i] = false
		}
	}

	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			// 当前对比的两个字符串相等 或者 p是"."，因为"."可以匹配任何字符
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// 因为p是"*"可以出现一次或者多次，所以我们要和上一位作比较
				// 如果当前s位置和p前一个位置不相等 或者 p的前一个位置不是"."
				if s[i-1] != p[j-2] && p[j-2] != '.' {
					// 可以当做"a*"出现了0次，直接拿掉
					dp[i][j] = dp[i][j-2]
				} else {
					// 到这里肯定是s当前位置和p前一个位置是可以匹配上的
					// 那么这里会出现3种："a*"可能出现 1次 2次 或者 n次
					// 出现0次，其实就是上面一样直接拿掉"a*"就好了
					// 出现1次，其实就是把"*"拿掉就好了
					// 如果是出现多次的情况，我们已经知道s[i-1] == p[j-2]，而且当前是"*"
					// s="aaaaa" p="a*"，最后a出现多少次其实是没关系，我们只需要知道第一个a在哪，所以要往前找
					dp[i][j] = dp[i][j-2] || dp[i][j-1] || dp[i-1][j]
				}
			}
		}
	}

	return dp[x][y]
}

// 如果是整个字符串进行比较复杂度会很高很高，那我们可以将已经匹配上的字符都丢掉，将剩下的继续匹配
// 所以递归来解这类题目会方便很多，我们只需要管当前那么一小段的状态即可
// 我们一般先判断p再判断s，因为p的情况比较复杂。如果两个一起判断也会增加复杂度
func isMatch(s string, p string) bool {
	// 第一种情况 都为空
	if len(p) == 0 {
		return len(s) == 0
	}

	// 第二种情况 长度为1的情况
	if len(p) == 1 {
		if len(s) == 1 && (s == p || p == ".") {
			return true
		} else {
			return false
		}
	}

	// 第三种情况 长度大于等于2，但是p的第二的字符不是*
	if p[1] != '*' {
		if len(s) == 0 {
			return false
		}
		return (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[1:])
	}

	// 第四种情况 长度大于等于2，但是p的第二的字符是*，
	for len(s) != 0 && (s[0] == p[0] || p[0] == '.') {
		if isMatch(s, p[2:]) {
			return true
		}
		// 为什么这里只需要对s进行处理？
		// 因为p的前两个字符会在下一次循环的自动被处理
		s = s[1:]
	}
	// 处理 s="",p="a*b" 这种情况，因为上面所有的条件都不满足
	return isMatch(s, p[2:])
}
