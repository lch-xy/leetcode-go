package main

// Base on Moore Voting Algorithm
// if we want to get the number greater than 1/3 , we need two candidates
// because there can be at most two numbers
func majorityElementMoore(nums []int) []int {
	candidate1, candidate2 := 0, 0
	count1, count2 := 0, 0
	for _, num := range nums {
		if candidate1 == num {
			count1++
		} else if candidate2 == num {
			count2++
		} else if count1 == 0 {
			candidate1 = num
			count1 = 1
		} else if count2 == 0 {
			candidate2 = num
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, num := range nums {
		if candidate1 == num {
			count1++
		} else if candidate2 == num {
			count2++
		}
	}

	res := make([]int, 0)
	if count1 > len(nums)/3 {
		res = append(res, candidate1)
	}
	if count2 > len(nums)/3 {
		res = append(res, candidate2)
	}

	return res
}

func majorityElement(nums []int) []int {
	limit := len(nums) / 3
	cache := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cache[nums[i]]++
	}
	res := []int{}
	for key, value := range cache {
		if value > limit {
			res = append(res, key)
		}
	}
	return res
}
