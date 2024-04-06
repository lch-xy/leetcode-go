package main

import (
	"fmt"
	"strings"
)

func main() {
	testCases := []struct {
		s        string
		expected string
	}{
		{"abc", "abc"},
		{"string", "rtsng"},
		{"poiinter", "ponter"},
	}

	for _, tc := range testCases {
		output := finalString(tc.s)
		if output != tc.expected {
			fmt.Printf("Test case failed: expected %s, got %s\n", tc.expected, output)
			return
		}
	}
}

func finalString(s string) string {
	slice := make([]string, 0)
	flag := 1
	for _, v := range s {
		if string(v) == "i" {
			flag *= -1
			continue
		}
		if flag == 1 {
			slice = append(slice, string(v))
		} else {
			slice = append([]string{string(v)}, slice...)
		}
	}
	result := strings.Join(slice, "")
	if flag == -1 {
		return ReverseString(result)
	}
	return result
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
