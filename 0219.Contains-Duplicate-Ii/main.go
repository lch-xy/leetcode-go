package main

func containsNearbyDuplicate(nums []int, k int) bool {
	cache := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if value, ok := cache[nums[i]]; ok {
			if abs(i, value) <= k {
				return true

			}
		}
		cache[nums[i]] = i
	}
	return false
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
