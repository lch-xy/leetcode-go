package main

import "fmt"

func main() {
	tests := []struct {
		n    int
		want int
	}{
		{
			n:    1,
			want: 1,
		},
		{
			n:    2,
			want: 2,
		},
		{
			n:    10,
			want: 12,
		},
		{
			n:    11,
			want: 15,
		},
	}

	for _, tt := range tests {
		got := nthUglyNumber(tt.n)
		if got != tt.want {
			fmt.Printf("nthUglyNumber(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

// 丑数定义：1是丑数，每个丑数的2、3、5倍也是丑数
// 1x2 , 1x3 , 1x5 都是丑数
// 我们将1放入队列，定义3个index分别代表2、3、5倍数的下标
// 每次从以求得的丑数中拿出最小的，放入队列中，以便后面使用
func nthUglyNumber(n int) int {
	res := []int{1}
	if n == 1 {
		return 1
	}
	idx2 := 0
	idx3 := 0
	idx5 := 0

	for len(res) < n {
		mm := min(res[idx2]*2, res[idx3]*3, res[idx5]*5)

		// 这里是3个if 而不是if else
		// 因为2x3 = 3x2 其实是一样的
		// 所以idx2和idx3都需要+1
		if mm == res[idx2]*2 {
			idx2++
		}
		if mm == res[idx3]*3 {
			idx3++
		}
		if mm == res[idx5]*5 {
			idx5++
		}
		res = append(res, mm)
	}

	return res[len(res)-1]
}
