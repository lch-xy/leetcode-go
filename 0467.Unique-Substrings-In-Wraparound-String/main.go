package main

import "fmt"

func main() {
	tests := []struct {
		s    string
		want int
	}{
		{"a", 1},
		{"cac", 2},
		{"zab", 6},
		{"abcdefghijklmnopqrstuvwxyz", 351},
		{"abczabc", 18},
	}

	for _, tt := range tests {
		if got := findSubstringInWraproundString(tt.s); got != tt.want {
			fmt.Printf("findSubstringInWraproundString(%v) = %v, want %v。", tt.s, got, tt.want)
		}
	}
}

// dp[i]:以26个字母结尾的字符串，最长连续子串长度
func findSubstringInWraproundString(s string) int {
	dp := make([]int, 27)
	count := 0
	for i := 0; i < len(s); i++ {
		if i >= 1 && (s[i]-s[i-1] == 1 || s[i-1]-s[i] == 25) {
			count++
		} else {
			count = 1
		}
		// 为什么这里是取最大值？
		// 因为一某个字母结尾的字符串，如果是连续的话，那么肯定是有重复的组合，所以取最大的即可，
		dp[s[i]-'a'] = max(dp[s[i]-'a'], count)
	}

	sum := 0
	for i := range dp {
		sum += dp[i]
	}
	return sum
}
