package main

import (
	"fmt"
	"math"
)

func main() {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			nums1: []int{1, 3, 5, 4},
			nums2: []int{1, 2, 3, 7},
			want:  1,
		},
		{
			nums1: []int{0, 4, 4, 5},
			nums2: []int{0, 1, 6, 8},
			want:  1,
		},
		{
			nums1: []int{3, 3, 8, 9, 10},
			nums2: []int{1, 7, 4, 6, 8},
			want:  1,
		},
	}

	for _, tt := range tests {
		got := minSwap(tt.nums1, tt.nums2)
		if got != tt.want {
			fmt.Printf("minSwap(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
		}
		got = minSwapDp(tt.nums1, tt.nums2)
		if got != tt.want {
			fmt.Printf("minSwapDp(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
		}
	}
}

func minSwapDp(nums1 []int, nums2 []int) int {
	n := len(nums1)
	// 初始化两个数组，keep 用于存储不交换当前位置元素的最小交换次数，swap 用于存储交换当前位置元素的最小交换次数
	keep := make([]int, n)
	swap := make([]int, n)
	keep[0] = 0
	swap[0] = 1
	for i := 1; i < n; i++ {
		keep[i] = n
		swap[i] = n
		// 如果当前位置的元素大于前一个位置的元素，那么我们可以选择不交换当前位置的元素，或者交换当前位置的元素
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			// 保持不变
			keep[i] = keep[i-1]
			// 如果当前位置进行交换，那就要同时交换i-1，使得序列变得合法
			swap[i] = swap[i-1] + 1
		}
		// 如果当前位置的元素大于前一个位置交换后的元素，那么我们可以选择交换当前位置的元素，或者不交换当前位置的元素
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			// 交换当前位置，前一位不变所以和keep[i-1]+1进行比较
			swap[i] = min(swap[i], keep[i-1]+1)
			// 不交换当前位置，交换前一位，所以是跟swap[i-1]进行比较
			keep[i] = min(keep[i], swap[i-1])
		}
	}
	return min(keep[n-1], swap[n-1])
}

func minSwap(nums1 []int, nums2 []int) int {
	res := math.MaxInt32
	helper(nums1, nums2, 1, 0, &res)
	return res
}

// index：当前处理到第几位了
// cnt：交换次数
// res：最小交换次数
func helper(nums1 []int, nums2 []int, index, cnt int, res *int) {
	// 剪枝
	if cnt >= *res {
		return
	}
	// 数组处理完毕
	if len(nums1) == index {
		*res = min(*res, cnt)
		return
	}

	if nums1[index] > nums1[index-1] && nums2[index] > nums2[index-1] {
		// 不做任何处理
		helper(nums1, nums2, index+1, cnt, res)
	}

	// 如果交换后还能保持有序，那就交换一下试试
	if nums2[index] > nums1[index-1] && nums1[index] > nums2[index-1] {
		swap(nums1, nums2, index)
		helper(nums1, nums2, index+1, cnt+1, res)
		swap(nums1, nums2, index)
	}
}

func swap(nums1 []int, nums2 []int, index int) {
	temp := nums1[index]
	nums1[index] = nums2[index]
	nums2[index] = temp
}
