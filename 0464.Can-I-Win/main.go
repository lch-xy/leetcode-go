package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	tests := []struct {
		maxChoosableInteger int
		desiredTotal        int
		expected            bool
	}{
		{10, 11, false},
		{10, 0, true},
		{10, 1, true},
		{10, 40, false},
		{18, 79, true},
	}

	for _, test := range tests {
		result := canIWin(test.maxChoosableInteger, test.desiredTotal)
		if result != test.expected {
			fmt.Printf("canIWin(%d, %d) = %v; expected %v。", test.maxChoosableInteger, test.desiredTotal, result, test.expected)
		}
	}
}

// 使用回溯算法进行决策，添加memo解决重复的子问题
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	state := make([]int, maxChoosableInteger+1)
	return canWin(state, maxChoosableInteger, desiredTotal, make(map[string]bool))
}

func canWin(state []int, maxChoosableInteger, desiredTotal int, hmap map[string]bool) bool {
	key := arrayToString(state)

	if _, ok := hmap[key]; ok {
		return hmap[key]
	}

	for i := 1; i <= maxChoosableInteger; i++ {
		if state[i] == 0 {
			state[i] = 1
			// 如果选择i后超过desiredTotal 或者 对方不能赢
			if desiredTotal-i <= 0 || !canWin(state, maxChoosableInteger, desiredTotal-i, hmap) {
				// 这里记录是原state是否可以赢，而不是改变后的state。所以key是原来的
				hmap[key] = true
				state[i] = 0
				return true
			}
			state[i] = 0
		}
	}
	hmap[key] = false
	return false
}

func arrayToString(arrys []int) string {
	var str strings.Builder
	for _, nums := range arrys {
		str.WriteString(strconv.Itoa(nums))
	}
	return str.String()
}
