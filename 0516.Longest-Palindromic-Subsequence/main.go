package main

import "fmt"

func main() {
	tests := []struct {
		input string
		want  int
	}{
		{"bbbab", 4},
		{"cbbd", 2},
		{"a", 1},
		{"abcde", 1},
		{"aaaaa", 5},
	}

	for _, test := range tests {
		got := longestPalindromeSubseq(test.input)
		if got != test.want {
			fmt.Printf("longestPalindromeSubseq(%q) = %v; want %v。", test.input, got, test.want)
		}
	}
}

func longestPalindromeSubseq(s string) int {
	size := len(s)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
		dp[i][i] = 1
	}
	for lenght := 2; lenght <= size; lenght++ {
		for i := 0; i <= size-lenght; i++ {
			j := i + lenght - 1
			// 如果i和j相邻也没事
			// 例如i=0，j=1 ---> i+1=1，j-1=0
			// dp[1][0]其实是在正方形的左下部分，也就是默认值0
			if s[i] == s[j] {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2)
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][size-1]
}
