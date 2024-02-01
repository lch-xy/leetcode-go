package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Test 1",
			nums: []int{1, 2, 3},
			want: []int{1, 2},
		},
		{
			name: "Test 2",
			nums: []int{1, 2, 4, 8},
			want: []int{1, 2, 4, 8},
		},
		{
			name: "Test 3",
			nums: []int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720},
			want: []int{9, 18, 90, 180, 360, 720},
		},
	}

	for _, tt := range tests {
		if got := largestDivisibleSubset(tt.nums); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("largestDivisibleSubset() = %v, want %v", got, tt.want)
		}
	}
}

// dp[i]:[0,i]区间内，最大整除子集数量
// prev[i]:前一位的index下标
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	size := len(nums)
	dp := make([]int, size)
	prev := make([]int, size)
	for i := range dp {
		dp[i] = 1
		prev[i] = -1
	}

	maxId := 0
	// 两层循环其实只是为了更新dp[i]
	// 更新dp[i]需要遍历[0,i]所有的数，算出能被nums[i]整除的数
	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
			if dp[i] > dp[maxId] {
				maxId = i
			}
		}
	}
	res := []int{}
	// 通过maxId 和 prev数组，就能从后往前推出这串数字了
	for i := maxId; i != -1; {
		res = append([]int{nums[i]}, res...)
		i = prev[i]
	}

	return res
}
