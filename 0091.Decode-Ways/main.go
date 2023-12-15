package main

import (
	"strconv"
	"strings"
)

func main() {
	println(numDecodings("226"))
	println(numDecodings("06"))
	println(numDecodings("10"))
	println(numDecodings("27"))
	println(numDecodings("2101"))
	println(numDecodings("1111111111111111111111111111111"))
	println(numDecodings("27"))

	println("=================================================>")

	println(numDecodingsDp("226"))
	println(numDecodingsDp("06"))
	println(numDecodingsDp("10"))
	println(numDecodingsDp("27"))
	println(numDecodingsDp("2101"))
	println(numDecodingsDp("1111111111111111111111111111111"))
	println(numDecodingsDp("27"))

}

// dp[i] 表示s中前i个字符组成的子串的解码方法的个数
// dp的i对应的是s[i-1]
// 大体思路：
// 如果当前这个数不是0，说明至少有dp[i-1]种解法,如果当前的数还能和前面的数拼成一个[10,26]的数，那就多出dp[i-2]种解法。dp[i] = dp[i-1] + dp[i-2]
// 如果当前是0，那么首先自己单独肯定无解，如果能和前面的数拼成一个[10,26]的数，那就说明还有救，有dp[i-2]种解法。dp[i] = dp[i-2]
func numDecodingsDp(s string) int {
	if strings.HasPrefix(s, "0") {
		return 0
	}
	toInt := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}
	size := len(s)
	dp := make([]int, size+1)
	dp[0] = 1
	for i := 1; i <= size; i++ {
		// 这里算的是算当前位置有多少种解法
		init := func() int {
			// 当出现“0”的时候，这条路一定行不通，只能和前面连起来看
			if toInt(s[i-1:i]) == 0 {
				return 0
			}
			// 当不是“0”的时候，说明至少有dp[i-1]种解法
			return dp[i-1]
		}
		dp[i] = init()
		// 这里算的是算当前位置有多少种解法
		// 如果我们将当前位置和前面的位置合并，并且这个数 10<=x<=26，说明是可以合并的
		// 如果我们合并这个数，那么至少有dp[i-2]种解法
		if i > 1 && (toInt(s[i-2:i]) <= 26 && toInt(s[i-2:i]) >= 10) {
			// 这里是在dp[i-1]的基础上加上dp[i-2]
			dp[i] += dp[i-2]
		}
	}

	return dp[size]
}

// 超时无法ac
// 思路就是将手动拆分 调用递归求解 res里存的是所有的解
func numDecodings(s string) int {
	res := [][]string{}
	helper(s, &res, []string{}, 0)
	return len(res)
}

func helper(s string, res *[][]string, curStrs []string, index int) {
	// 因为index是从0开始的，当index = len的时候 说明已经超了
	if len(s) == index {
		*res = append(*res, curStrs)
		return
	}
	toInt := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}
	s1 := s[index]
	s1Str := []string{}
	copy(s1Str, curStrs)
	// 如果当前的“0”，直接return。
	if string(s1) == "0" {
		return
	}
	s1Str = append(s1Str, string(s1))
	helper(s, res, s1Str, index+1)

	// 当前值和后面值合起来 10 <= x <= 26，说明是有效的，可以这么拆
	if index+1 < len(s) && toInt(s[index:index+2]) <= 26 && toInt(s[index:index+2]) >= 10 {
		s2 := s[index+1]
		s2Str := []string{}
		copy(s2Str, curStrs)
		s2Str = append(s2Str, string(s1)+string(s2))
		helper(s, res, s2Str, index+2)
	}
}
