package main

import "fmt"

func main() {
	testCases := []struct {
		s        string
		expected byte
	}{
		{s: "leetcode", expected: 'e'},
		{s: "loveleetcode", expected: 'l'},
		{s: "hello", expected: 'l'},
	}

	for _, tc := range testCases {
		output := repeatedCharacter(tc.s)
		if output != tc.expected {
			fmt.Printf("Test case failed: expected %c, got %c\n", tc.expected, output)
			return
		}
	}

	fmt.Println("All test cases passed!")
}

func repeatedCharacter(s string) byte {
	cache := make([]int, 26)
	for _, char := range s {
		if cache[char-'a'] == 1 {
			return byte(char)
		}
		cache[char-'a']++
	}
	return byte(0)
}
