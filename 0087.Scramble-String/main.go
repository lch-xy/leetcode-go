package main

import (
	"reflect"
	"sort"
	"strings"
)

func main() {
	println(isScrambleCutOff("great", "rgeat"))
	println(isScrambleCutOff("abcdbdacbdac", "bdacabcdbdac"))
	println(isScrambleDp("eebaacbcbcadaaedceaaacadccd", "eadcaacabaddaceacbceaabeccd"))
	println(isScrambleDp("abcdefghijklmfasgasgas", "abcdefghijklmfasgasgas"))
}

// dp[i][j][len]表示的是以i和j为起点的长度为 len 的字符串是不是互为 scramble
// 为什么可以通过dp去计算？
// 我们通过将len放到最外层，那么在进行而k又是小于当前的len的，那么正好可以利用之前算的数据
// 例如len=3时，我们我们判断是否scramble时，那么k就要选择在0-2之间进行切割，而0-2我们之前在最外层已经都算好了
func isScrambleDp(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	size := len(s1)
	dp := make([][][]bool, size)
	for i := 0; i < size; i++ {
		dp[i] = make([][]bool, size)
		for j := 0; j < size; j++ {
			dp[i][j] = make([]bool, size+1)
			for k := 0; k <= size; k++ {
				dp[i][j][k] = false
			}
		}
	}
	// i + len的长度不能大于size
	for len := 1; len <= size; len++ {
		for i := 0; i <= size-len; i++ {
			for j := 0; j <= size-len; j++ {
				if len == 1 {
					res := func() bool {
						// 长度为1时 只用判断两个字符串是否相等即可
						if s1[i:i+1] == s2[j:j+1] {
							return true
						}
						return false
					}
					dp[i][j][len] = res()
				} else {
					// 如果我们要判断dp[i][j][len]是否是scramble？那我们就要在1-len之前找到一个k值，可以将len分割并且是scramble
					// 所以我们要遍历[1,len)，找个看看是否满足条件
					for k := 1; k < len; k++ {
						// 例如：将"qwertyuiop"&"asdfghjklz"进行比较时
						// k=1表示
						// (i=0 j=0 k=1 && i+k=1 j+k=1 len-k=9 ) || (i+k=1 j=0 len-k=9 && i=0 j+len-k=9 k=1)
						// ("q"&"a"和"wertyuiop"&"sdfghjklz") || ("wertyuiop"&"asdfghjkl"和"q"&"z")
						// k=2表示
						// (i=0 j=0 k=2 && i+k=2 j+k=2 len-k=8 ) || (i=0 j+len-k=8 k=2 && i+k=2 j=0 len-k=8 )
						// ("qw"&"as" 和 "ertyuiop"&"dfghjklz") || ("qw"&"lz" 和 "ertyuiop"&"asdfghjk")
						if (dp[i][j][k] && dp[i+k][j+k][len-k]) || (dp[i][j+len-k][k]) && dp[i+k][j][len-k] {
							dp[i][j][len] = true
						}
					}
				}
			}
		}
	}
	return dp[0][0][size]
}

// 对s1和s2进行切分，然后递归调用去判断
// 如果s1和s2包含的数据不一样 直接return false
// 如果s1==s2或者翻转后s1==s2，就return true
func isScrambleCutOff(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return true
	}
	if !areStringsEqual(s1, s2) {
		return false
	}
	if s1 == s2 || reverseString(s1) == s2 {
		return true
	}
	for i := 1; i < len(s1); i++ {
		if isScrambleCutOff(s1[:i], s2[:i]) && isScrambleCutOff(s1[i:], s2[i:]) {
			return true
		}
		if isScrambleCutOff(s1[:i], s2[len(s1)-i:]) && isScrambleCutOff(s1[i:], s2[:len(s1)-i]) {
			return true
		}
	}
	return false
}

func areStringsEqual(s1, s2 string) bool {
	// 将字符串转换为字符切片
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	// 对字符切片进行排序
	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	// 使用 DeepEqual 判断两个字符切片是否相同
	return reflect.DeepEqual(runes1, runes2)
}

func reverseString(s string) string {
	var builder strings.Builder

	for i := len(s) - 1; i >= 0; i-- {
		builder.WriteRune(rune(s[i]))
	}
	reversedStr := builder.String()

	return reversedStr
}
