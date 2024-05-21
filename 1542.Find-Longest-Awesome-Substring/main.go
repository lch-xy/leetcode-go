package main

import "fmt"

func main() {
	fmt.Printf("longestAwesome(\"3242415\"): %v\n", longestAwesome("3242415"))
	fmt.Printf("longestAwesome(\"12345678\"): %v\n", longestAwesome("12345678"))
	fmt.Printf("longestAwesome(\"213123\"): %v\n", longestAwesome("213123"))
}

// palindromes have two situations , all composed by even number or include one odd number
// so we should remark the number occurrences , only one number can appear odd times
// use bitmap to save the number , use 0 to present even number and 1 to present odd number
// traversal string to count the number occurrences , here we use 1 << n to left shift n bits
// if the bitmap not in hashmap means it's the first appearence , we should store in hashmap,key is the number and value is the index
// if we meet the number again , means the number in the middle of the current position and the previous position appears even times
// how to caculate palindromes composed by including only one odd number？
// flip the number on each digit in sequence
// if this number in the hashmap means the number in the middle of the current position and the previous position only one number appears odd times
func longestAwesome(s string) int {
	n := len(s)
	prefix := 0
	cache := make(map[int]int)
	cache[0] = -1
	res := 0
	for i := 0; i < n; i++ {
		curNumber := s[i] - '0'
		prefix ^= 1 << curNumber
		if _, ok := cache[prefix]; ok {
			res = max(res, i-cache[prefix])
		} else {
			cache[prefix] = i
		}

		for j := 0; j <= 9; j++ {
			prefix ^= 1 << j
			if _, ok := cache[prefix]; ok {
				res = max(res, i-cache[prefix])
			}
			prefix ^= 1 << j
		}
	}
	return res
}

// time out
// func longestAwesome(s string) int {
// 	arr := make([]int, 10)
// 	res := 1
// 	for i := 0; i < len(s); i++ {
// 		clear(arr)
// 		for j := i; j < len(s); j++ {
// 			arr[s[j]-'0']++
// 			if !checkArrays(arr) {
// 				continue
// 			}
// 			res = max(res, j-i+1)
// 		}
// 	}
// 	return res
// }

// func checkArrays(arr []int) bool {
// 	oddCnt := 0
// 	for _, v := range arr {
// 		if v&1 == 1 {
// 			oddCnt++
// 		}
// 		if oddCnt > 1 {
// 			return false
// 		}
// 	}
// 	return true
// }
