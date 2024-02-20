package main

import "fmt"

func main() {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 2, 3, 1}, 2},
		{[]int{2, 4, 3, 5, 1}, 3},
		{[]int{1, 1, 1, 1, 1}, 0},
		{[]int{5, 4, 3, 2, 1}, 4},
	}

	for _, tt := range tests {
		got := reversePairsForce(tt.nums)
		if got != tt.want {
			fmt.Printf("reversePairsForce(%v) = %v, want %v。", tt.nums, got, tt.want)
		}
	}

	for _, tt := range tests {
		got := reversePairs(tt.nums)
		if got != tt.want {
			fmt.Printf("reversePairs(%v) = %v, want %v。", tt.nums, got, tt.want)
		}
	}
}

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := left + (right-left)/2
	leftCnt := mergeSort(nums, left, mid)
	rightCnt := mergeSort(nums, mid+1, right)

	// 归并排序后，左右两边应该是有序的了
	// 我们直接使用双指针遍历两个数组，找到满足nums[i] > 2*nums[j]的数了
	res := 0
	for i, j := left, mid+1; i <= mid; i++ {
		for j <= right && nums[i] > 2*nums[j] {
			j++
		}
		res += j - mid - 1
	}

	// 计算完毕后，将数组排好序，提供给上一层使用
	// 因为左右都已经排序好了，所以这里使用归并排序也是很快的
	merge(nums, left, mid, right)

	return leftCnt + rightCnt + res
}

// 归并排序
func merge(nums []int, left, mid, right int) {
	i, j := left, mid+1
	temp := make([]int, right-left+1)
	k := 0

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = nums[i]
		k++
		i++
	}

	for j <= right {
		temp[k] = nums[j]
		k++
		j++
	}
	copy(nums[left:right+1], temp)
}

func reversePairsForce(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]*2 < nums[i] {
				res++
			}
		}
	}
	return res
}
