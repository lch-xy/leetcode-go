package main

import "fmt"

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 7, 4, 9, 2, 5}, 6},
		{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 2},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 1},
	}

	for _, testCase := range testCases {
		result := wiggleMaxLength(testCase.nums)
		if result != testCase.expected {
			fmt.Printf("Error: expected %d, got %d", testCase.expected, result)
		}
	}
}

// up[i] 表示以第 i 个元素结尾的最长上升摆动序列的长度。
// down[i] 表示以第 i 个元素结尾的最长下降摆动序列的长度。
// 如果当前元素大于前一个元素，那么我们可以将它接在前一个元素的下降摆动序列后面，形成一个更长的上升摆动序列。所以 up[i] = max(up[i], down[j] + 1)。
// 如果当前元素小于前一个元素，那么我们可以将它接在前一个元素的上升摆动序列后面，形成一个更长的下降摆动序列。所以 down[i] = max(down[i], up[j] + 1)。
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up := make([]int, n)
	down := make([]int, n)
	up[0], down[0] = 1, 1
	for i := 1; i < n; i++ {
		// 上升的时候别忘了处理下降的数组
		// 下降的时候别忘了处理上升的数组
		// 不然虽然还是通过ac，但是数组里会存在0的情况，其实是不应该存在的
		if nums[i] > nums[i-1] {
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			down[i] = max(down[i-1], up[i-1]+1)
			up[i] = up[i-1]
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	return max(up[n-1], down[n-1])
}
