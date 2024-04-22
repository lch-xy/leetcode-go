package main

func sumImbalanceNumbers(nums []int) int {
	size := len(nums)
	res := 0

	for i := 0; i < size; i++ {
		cnt := -1
		visited := make([]int, 1002)

		for j := i; j < size; j++ {
			if visited[nums[j]] == 0 {
				cnt++
				if visited[nums[j]+1] > 0 {
					cnt--
				}
				if visited[nums[j]-1] > 0 {
					cnt--
				}
			}
			visited[nums[j]] = 1
			res += cnt
		}
	}
	return res
}
