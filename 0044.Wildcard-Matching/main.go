package main

import "strings"

func main() {

	println(isMatchDoublePoint("aa", "*"))
	println(isMatchDoublePoint("aa", "a?"))
	println(isMatchDoublePoint("avjskljfdsl", "a*jfdsl"))
	println(isMatchDoublePoint("aabbvvss", "a*bv?ss"))
	println(isMatchDoublePoint("avjskljfdsl", "a*jfdssl"))
	println(isMatchDoublePoint("abcdefabdef", "a*bd*"))
	println(isMatchDoublePoint("adceb", "*a*b"))
	println(isMatchDoublePoint("aaabababaaabaababbbaaaabbbbbbabbbbabbbabbaabbababab", "*ab***ba**b*b*aaab*b"))
	println(isMatchDoublePoint("aaaabaaaabbbbaabbbaabbaababbabbaaaababaaabbbbbbaabbbabababbaaabaabaaaaaabbaabbbbaababbababaabbbaababbbba", "*****b*aba***babaa*bbaba***a*aaba*b*aa**a*b**ba***a*a*"))
	println(isMatchDoublePoint("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))

	println("===============================================>")

	println(isMatchDp("aa", "*"))
	println(isMatchDp("aa", "a?"))
	println(isMatchDp("avjskljfdsl", "a*jfdsl"))
	println(isMatchDp("aabbvvss", "a*bv?ss"))
	println(isMatchDp("avjskljfdsl", "a*jfdssl"))
	println(isMatchDp("abcdefabdef", "a*bd*"))
	println(isMatchDp("adceb", "*a*b"))
	println(isMatchDp("aaabababaaabaababbbaaaabbbbbbabbbbabbbabbaabbababab", "*ab***ba**b*b*aaab*b"))
	println(isMatchDp("aaaabaaaabbbbaabbbaabbaababbabbaaaababaaabbbbbbaabbbabababbaaabaabaaaaaabbaabbbbaababbababaabbbaababbbba", "*****b*aba***babaa*bbaba***a*aaba*b*aa**a*b**ba***a*a*"))
	println(isMatchDp("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))

	println("===============================================>")

	println(isMatchRecursion("aa", "*"))
	println(isMatchRecursion("aa", "a?"))
	println(isMatchRecursion("avjskljfdsl", "a*jfdsl"))
	println(isMatchRecursion("aabbvvss", "a*bv?ss"))
	println(isMatchRecursion("avjskljfdsl", "a*jfdssl"))
	println(isMatchRecursion("abcdefabdef", "a*bd*"))
	println(isMatchRecursion("adceb", "*a*b"))
	println(isMatchRecursion("aaabababaaabaababbbaaaabbbbbbabbbbabbbabbaabbababab", "*ab***ba**b*b*aaab*b"))
	println(isMatchRecursion("aaaabaaaabbbbaabbbaabbaababbabbaaaababaaabbbbbbaabbbabababbaaabaabaaaaaabbaabbbbaababbababaabbbaababbbba", "*****b*aba***babaa*bbaba***a*aaba*b*aa**a*b**ba***a*a*"))
	//println(isMatchRecursion("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))

	println("===============================================>")

	println(isMatchRecursionNew("aa", "*"))
	println(isMatchRecursionNew("aa", "a?"))
	println(isMatchRecursionNew("avjskljfdsl", "a*jfdsl"))
	println(isMatchRecursionNew("aabbvvss", "a*bv?ss"))
	println(isMatchRecursionNew("avjskljfdsl", "a*jfdssl"))
	println(isMatchRecursionNew("abcdefabdef", "a*bd*"))
	println(isMatchRecursionNew("adceb", "*a*b"))
	println(isMatchRecursionNew("aaabababaaabaababbbaaaabbbbbbabbbbabbbabbaabbababab", "*ab***ba**b*b*aaab*b"))
	println(isMatchRecursionNew("aaaabaaaabbbbaabbbaabbaababbabbaaaababaaabbbbbbaabbbabababbaaabaabaaaaaabbaabbbbaababbababaabbbaababbbba", "*****b*aba***babaa*bbaba***a*aaba*b*aa**a*b**ba***a*a*"))
	println(isMatchRecursionNew("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))

}

// 大佬能过ac的代码
// 如何剪枝？我们定义三种状态：
// 1.返回0表示匹配到了s串的末尾，但是未匹配成功；
// 2.返回1表示未匹配到s串的末尾就失败了；
// 3.返回2表示成功匹配。那么只有返回值大于1，才表示成功匹配。
// 及时的对状态为0或者2的进行剪枝
func isMatchRecursionNew(s string, p string) bool {
	return helper(s, p, 0, 0) > 1
}

func helper(s string, p string, i int, j int) int {
	if i == len(s) && j == len(p) {
		return 2
	}
	if i == len(s) && p[j] != '*' {
		return 0
	}
	if j == len(p) {
		return 1
	}

	if s[i] == p[j] || p[j] == '?' {
		return helper(s, p, i+1, j+1)
	}

	if p[j] == '*' {
		if j+1 < len(p) && p[j+1] == '*' {
			return helper(s, p, i, j+1)
		}

		for k := 0; k <= len(s)-i; k++ {
			res := helper(s, p, i+k, j+1)
			if res == 0 || res == 2 {
				return res
			}
		}
	}

	return 1
}

// 优化后的递归代码 依然无法通过ac，超时
// 即使加了缓存，也过不了ac
func isMatchRecursion(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(s) == 0 {
		p = strings.ReplaceAll(p, "*", "")
		return len(p) == 0
	}

	if p[0] == '*' {
		// 跳过多个*
		for len(p) > 1 && p[0] == p[1] {
			p = p[1:]
		}
		// 其实就是慢在这里，两个递归是或的关系，直接爆炸
		return (len(s) > 0 && isMatchRecursion(s[1:], p)) || isMatchRecursion(s, p[1:])
	} else {
		return len(s) > 0 && (s[0] == p[0] || p[0] == '?') && isMatchRecursion(s[1:], p[1:])
	}
}

// 第一版递归代码
//func isMatchRecursion(s string, p string) bool {
//	if len(p) == 0 {
//		if len(s) == 0 {
//			return true
//		} else {
//			return false
//		}
//	}
//
//	if len(p) == 1 {
//		if (len(s) == 1 && (s[0] == p[0] || p[0] == '?')) || p[0] == '*' {
//			return true
//		} else {
//			return false
//		}
//	}
//
//	if p[0] == '*' {
//		if len(s) == 0 {
//			return isMatch(s, p[1:])
//		} else {
//			return isMatch(s[1:], p) || isMatch(s[1:], p[1:])  || isMatch(s, p[1:])
//		}
//	} else {
//		if len(s) > 0 && (s[0] == p[0] || p[0] == '?') {
//			return isMatch(s[1:], p[1:])
//		}
//	}
//	return false
//}

// dp[i][j] 以s[i-1] 和 p[j-1] 结尾的字符串是否匹配
func isMatchDp(s string, p string) bool {
	n := len(s)
	m := len(p)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true // 空串肯定是匹配的

	// 初始化第一行
	// 当前是“*”，并且前一位是true
	for j := 1; j <= m; j++ {
		if p[j-1] == '*' && dp[0][j-1] {
			dp[0][j] = true
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j-1] == '*' {
				// 因为“*”可以匹配任何或者空串
				// 所以只要 p中的前j个字符跟s中的前 i-1 个字符匹配成功 or  j-1 个字符跟s中前i个字符匹配成功
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else {
				// 如果不是“*”，则必须要当前字符匹配上，并且之前的也要匹配上
				dp[i][j] = (s[i-1] == p[j-1] || p[j-1] == '?') && dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}

// 使用双指针，如果匹配规则是尽量向后匹配，如果一旦匹配不上了，那就会回到最后面的“*”的位置，进行通配符的匹配
// 因为iStar记录到最后一次匹配的s的位置，所以只需要将iStar++并复制给i，就可以回到位置
// p的位置尽量向后匹配，只要s[i] == p[j] 就往后移动，如果无法匹配且p也不为“*”的话，就让s往前移动，p保持不变，因为p前面的*可以匹配任意的字符串
// iStart：当前s中准备和*匹配的位置
// jStart：当前p中匹配到*的位置
func isMatchDoublePoint(s string, p string) bool {
	i := 0
	j := 0
	iStar := -1
	jStar := -1
	m := len(s)
	n := len(p)

	for i < m {
		if j < n && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < n && p[j] == '*' {
			// 当找到“*”后，我们要更新iStar和jStar信息
			iStar = i
			jStar = j
			j++
		} else if iStar >= 0 {
			// 如果没有匹配到，但是之前有*号
			iStar++
			i = iStar
			j = jStar + 1
		} else {
			return false
		}
	}
	// s匹配完后如果p后面都是*的话，也是可以匹配上
	for j < n && p[j] == '*' {
		j++
	}
	return j == n
}
