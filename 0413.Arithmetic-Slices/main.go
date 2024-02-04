package main

import "fmt"

func main() {
	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Case 1: No arithmetic slices",
			nums: []int{1, 2, 4},
			want: 0,
		},
		{
			name: "Case 2: One arithmetic slice",
			nums: []int{1, 2, 3},
			want: 1,
		},
		{
			name: "Case 3: Multiple arithmetic slices",
			nums: []int{1, 2, 3, 4},
			want: 3,
		},
		{
			name: "Case 4",
			nums: []int{1, 2, 3, 8, 9, 10},
			want: 2,
		},
		{
			name: "Case 5",
			nums: []int{2, 1, 3, 4, 2, 3},
			want: 0,
		},
	}

	for _, tc := range testCases {
		got := numberOfArithmeticSlices(tc.nums)
		if got != tc.want {
			fmt.Printf("numberOfArithmeticSlices(%v) = %v; want %v。。。。。", tc.nums, got, tc.want)
		}
	}
}

// nums:[1,2,3,4,5]
// len = 3: [1,2,3], [2,3,4], [3,4,5]
// len = 4: [1,2,3,4], [2,3,4,5]
// len = 5: [1,2,3,4,5]
// 可以推出 res = (n-1)(n-2)/2
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	res := 0
	size := 2
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			size++
		} else {
			if size > 2 {
				res += (size - 1) * (size - 2) / 2
			}
			// 重置成2，不是0
			size = 2
		}
	}
	// 如果最后一个数能构成等差队列，就需要在这一步兜底计算
	if size > 2 {
		res += (size - 1) * (size - 2) / 2
	}
	return res
}
