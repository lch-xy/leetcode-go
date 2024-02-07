package main

import (
	"fmt"
)

func main() {
	tests := []struct {
		n      int
		expect int
	}{
		{n: 0, expect: 1},
		{n: 1, expect: 10},
		{n: 2, expect: 91},
		{n: 3, expect: 739},
		{n: 4, expect: 5275},
		{n: 5, expect: 32491},
		{n: 6, expect: 168571},
		{n: 7, expect: 712891},
		{n: 8, expect: 2345851},
		{n: 9, expect: 5611771},
		{n: 10, expect: 8877691},
	}

	for _, test := range tests {
		result := countNumbersWithUniqueDigits(test.n)
		if result != test.expect {
			fmt.Printf("countNumbersWithUniqueDigits(%d) = %d; expect %d", test.n, result, test.expect)
		}
	}
}

// 思路是基于排列组合的原理
// 当 i=2 时，首位有 9 种选择（1-9），第二位有 9 种选择（0 和剩下的 8 个数字），所以有 9*9 种选择。
// 当 i=3 时，首位有 9 种选择，第二位有 9 种选择，第三位有 8 种选择，所以有 998 种选择。
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	res := 10
	cnt := 9
	for i := 2; i <= n; i++ {
		cnt *= (11 - i)
		res += cnt
	}
	return res
}
