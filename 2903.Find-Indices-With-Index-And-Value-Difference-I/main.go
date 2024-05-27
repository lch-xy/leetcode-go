package main

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if abs(i, j) >= indexDifference && abs(nums[i], nums[j]) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
