package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := []struct {
		envelopes [][]int
		expect    int
	}{
		{
			envelopes: [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
			expect:    3,
		},
		{
			envelopes: [][]int{{1, 1}, {1, 1}, {1, 1}},
			expect:    1,
		},
		{
			envelopes: [][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}},
			expect:    4,
		},
	}

	for _, test := range tests {
		result := maxEnvelopes(test.envelopes)
		if result != test.expect {
			fmt.Printf("maxEnvelopes(%v) = %d; expect %d", test.envelopes, result, test.expect)
		}
	}
}

// 解法类似于Longest-Increasing-Subsequence
// 原方法无法通过ac，这里主要用二分查找对时间复杂度进行了优化
// 我们通过算法会得到一个tail数组，这和tail数组的长度就是我们要的答案
// 思想：如果当前数大于之后一个数，直接追加。如果小于这个数，那么找到比这个数大的第一个数进行替换。
// 如果遍历到3时，会将[2,5] -> [2,3]，因为3比5小，找打比3大的数拼成递增序列肯定比5要多
// ===============================================================================>
// nums = [10, 9, 2, 5, 3, 7, 101, 18]
// 对于 nums[0] = 10，tails 是空的，所以我们将 10 添加到 tails 的末尾。现在，tails = [10]。
// 对于 nums[1] = 9，我们在 tails 中找到 10 的位置，然后用 9 替换 10。现在，tails = [9]。
// 对于 nums[2] = 2，我们在 tails 中找到 9 的位置，然后用 2 替换 9。现在，tails = [2]。
// 对于 nums[3] = 5，5 大于 tails 的最后一个元素 2，所以我们将 5 添加到 tails 的末尾。现在，tails = [2, 5]。
// 对于 nums[4] = 3，我们在 tails 中找到 5 的位置，然后用 3 替换 5。现在，tails = [2, 3]。
// 对于 nums[5] = 7，7 大于 tails 的最后一个元素 3，所以我们将 7 添加到 tails 的末尾。现在，tails = [2, 3, 7]。
// 对于 nums[6] = 101，101 大于 tails 的最后一个元素 7，所以我们将 101 添加到 tails 的末尾。现在，tails = [2, 3, 7, 101]。
// 对于 nums[7] = 18，我们在 tails 中找到 101 的位置，然后用 18 替换 101。现在，tails = [2, 3, 7, 18]。
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	tail := []int{}
	for _, curNumber := range envelopes {
		if len(tail) == 0 {
			tail = append(tail, curNumber[1])
		} else {
			if curNumber[1] > tail[len(tail)-1] {
				tail = append(tail, curNumber[1])
			} else {
				left := 0
				right := len(tail)
				for left < right {
					mid := left + (right-left)/2
					if tail[mid] >= curNumber[1] {
						right = mid
					} else if tail[mid] < curNumber[1] {
						left = mid + 1
					}
				}
				tail[right] = curNumber[1]
			}
		}
	}
	return len(tail)
}
