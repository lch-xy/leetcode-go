package main

import "fmt"

func main() {
	fmt.Printf("longestEqualSubarray([]int{1, 3, 2, 3, 1, 3}, 3): %v\n", longestEqualSubarray([]int{1, 3, 2, 3, 1, 3}, 3))
	fmt.Printf("longestEqualSubarray([]int{1,1,2,2,1,1}, 2): %v\n", longestEqualSubarray([]int{1, 1, 2, 2, 1, 1}, 2))

}

// use hashmap to store index of the number appeared
// iterate each number's index list
// values[j]-values[i]+1 is the total length of key
// (j-i+1) is the key length
// for example :
// 1 1 2 2 3 1 1 2
// 1_list : [0,1,5,6]
// 2_list : [2,3,7]
// 3_list : [4]
// iterate 1_list ,if i=0 j= 2
// 1_list[2] = 5  1_list[0] = 0
// 5-0+1 is the total length
// 2-0+1 is the 1 length
// so 5-0+1 -(2-0+1) should less or equal then k
func longestEqualSubarray(nums []int, k int) int {
	cache := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		cache[nums[i]] = append(cache[nums[i]], i)
	}
	res := 0
	for _, values := range cache {
		j := 0
		for i := 0; i < len(values); i++ {
			for j < len(values) && values[j]-values[i]-(j-i) <= k {
				res = max(res, j-i+1)
				j++
			}
		}
	}
	return res
}

// timeout
// func longestEqualSubarray(nums []int, k int) int {
// 	n := len(nums)
// 	res := 0
// 	for i := 0; i < n; i++ {
// 		start := i + 1
// 		skip := 0
// 		for start < n {
// 			if nums[i] != nums[start] {
// 				skip++
// 			}
// 			if skip > k {
// 				break
// 			}
// 			start++
// 		}
// 		res = max(res, start-i-min(k, skip))
// 	}
// 	return res
// }
