package main

import "fmt"

func main() {
	tests := []struct {
		s1   string
		n1   int
		s2   string
		n2   int
		want int
	}{
		// {"acb", 4, "ab", 2, 2},
		// {"abc", 3, "abc", 3, 1},
		// {"abc", 1, "abc", 1, 1},
		// {"aaa", 3, "aa", 1, 4},
		{"abaacdbac", 4, "adcbd", 1, 2},
	}

	for _, tt := range tests {
		got := getMaxRepetitions(tt.s1, tt.n1, tt.s2, tt.n2)
		if got != tt.want {
			fmt.Printf("getMaxRepetitions(%v, %v, %v, %v) = %v; want %v。", tt.s1, tt.n1, tt.s2, tt.n2, got, tt.want)
		}
	}
}

// s1:abaacdbac n1:4  s2:adcbd n2:2
//                   |<-            这里是循环节             ->|
// a b a a c d b a c | a b a a c d b a c | a b a a c d b a c | a b a a c d b a c
// a         d     c     b       d | a               d     c     b       d | a
// sum[3,6,8]
// mod[3,1,3]
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	len2 := len(s2)
	sum, mod := []int{}, []int{}
	cnt := 0
	loop := true
	for loop {
		for i := range s1 {
			if s1[i] == s2[cnt%len2] {
				cnt++
			}
		}
		for _, v := range mod {
			if v == cnt%len2 {
				loop = false
				break
			}
		}
		sum = append(sum, cnt)
		mod = append(mod, cnt%len2)
	}

	// 循环节占几个s1
	s1_repeat := len(sum) - 1
	// 为什么是n1-1？因为从第二开始才叫循环，如上图所示，中间是才是循环
	coeff, leftover := (n1-1)/s1_repeat, (n1-1)%s1_repeat
	// 为什么要加上sum[leftover]？
	// 因为我们将中间部分剔除后，头和尾会进行一个组合
	// 上面的例子会得到leftover=1，再加上头，其实sum[leftover]就是算前两个能匹配多少个s2
	// 这时候leftover正好就是sum的坐标，如果leftover=0，那么sum[0]正好也就只算了头部，很巧妙
	final := (sum[len(sum)-1]-sum[0])*coeff + sum[leftover]

	return final / len2 / n2
}
