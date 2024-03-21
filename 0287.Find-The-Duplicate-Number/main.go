package main

import "fmt"

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{1, 3, 4, 2, 2}, expected: 2},
		{nums: []int{3, 1, 3, 4, 2}, expected: 3},
	}

	for i, tc := range testCases {
		result := findDuplicateTwo(tc.nums)
		if result != tc.expected {
			fmt.Printf("Test case %d failed: Expected output %d, got %d\n", i+1, tc.expected, result)
		}
	}
}

// Use Floyd's Tortoise and Hare
func findDuplicateTwo(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = nums[0]

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

// timeout
func findDuplicate(nums []int) int {
	left, right := 0, len(nums)-1

	// use binary search to find the duoplicate number
	for left < right {
		mid := left + (right-left)/2
		count := 0

		// iterate over the array and count elements less or equal to mid
		for _, num := range nums {
			// include mid
			if num <= mid {
				count++
			}
		}

		// if the count is greater than mid, thw duplicate number is in the left half
		if count > mid {
			right = mid
		} else {
			left = left + 1
		}

	}
	// when the left == right
	return left
}
