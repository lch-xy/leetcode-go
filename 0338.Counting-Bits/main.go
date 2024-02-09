package main

import (
	"fmt"
	"reflect"
)

func main() {
	tests := []struct {
		n      int
		expect []int
	}{
		{
			n:      2,
			expect: []int{0, 1, 1},
		},
		{
			n:      5,
			expect: []int{0, 1, 1, 2, 1, 2},
		},
		{
			n:      10,
			expect: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2},
		},
	}

	for _, test := range tests {
		result := countBits(test.n)
		if !reflect.DeepEqual(result, test.expect) {
			fmt.Printf("countBits(%d) = %v; expect %v", test.n, result, test.expect)
		}
	}
}

// 奇数：dp[i]=dp[i-1] +1
// 偶数：dp[i]=dp[i/2]
func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	res := make([]int, n+1)
	res[0] = 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			res[i] = res[i-1] + 1
		} else {
			res[i] = res[i/2]
		}
	}
	return res
}
