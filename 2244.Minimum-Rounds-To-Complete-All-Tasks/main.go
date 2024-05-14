package main

// one step : find only appeared once number
// two step : use 3 as denominator, if the remainder is 1 , we can divide data into (2,2)
// if the remainder is 2 ,we can divide data into (2,3)
// so if remainder is not 0 , result should add 1
func minimumRounds(tasks []int) int {
	onceAppear := 0
	cache := make(map[int]int)
	for i := 0; i < len(tasks); i++ {
		cache[tasks[i]]++
		if cache[tasks[i]] == 1 {
			onceAppear++
		} else if cache[tasks[i]] == 2 {
			onceAppear--
		}
	}
	if onceAppear > 0 {
		return -1
	}
	res := 0
	for _, value := range cache {
		res += value / 3
		if value%3 != 0 {
			res++
		}
	}
	return res
}
