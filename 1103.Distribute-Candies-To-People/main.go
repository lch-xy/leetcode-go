package main

import "fmt"

func main() {
	fmt.Printf("distributeCandies(7, 4): %v\n", distributeCandies(7, 4))
	fmt.Printf("distributeCandies(10, 3): %v\n", distributeCandies(10, 3))
}

func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)
	start := 1
	index := 0
	for candies > 0 {
		if candies < start {
			start = candies
		}
		res[index%num_people] += start
		candies -= start
		start++
		index++
	}
	return res
}
