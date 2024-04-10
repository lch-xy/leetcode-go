package main

import (
	"fmt"
	"strings"
)

func main() {
	testCases := []struct {
		binary   string
		expected string
	}{
		{binary: "000110", expected: "111011"},
		{binary: "01", expected: "01"},
		{binary: "0", expected: "0"},
		{binary: "111", expected: "111"},
		{binary: "1100", expected: "1110"},
	}

	for i, tc := range testCases {
		maxBinary := maximumBinaryString(tc.binary)
		if maxBinary != tc.expected {
			fmt.Printf("Test case %d failed: expected %s, got %s\n", i+1, tc.expected, maxBinary)
			return
		}
		fmt.Printf("Test case %d passed\n", i+1)
	}
	fmt.Println("=====================================>")
	for i, tc := range testCases {
		maxBinary := maximumBinaryStringTwo(tc.binary)
		if maxBinary != tc.expected {
			fmt.Printf("Test case %d failed: expected %s, got %s\n", i+1, tc.expected, maxBinary)
			return
		}
		fmt.Printf("Test case %d passed\n", i+1)
	}
}

// count the number of '0' after the first '0'
func maximumBinaryStringTwo(binary string) string {
	length := len(binary)
	firstIndex := strings.Index(binary, "0")
	if firstIndex == -1 {
		return binary
	}
	count := firstIndex
	for i := firstIndex + 1; i < length; i++ {
		if binary[i] == '0' {
			count++
		}
	}

	// we should use []byte{} to instead of string
	// because stirngs int go area immutable, each operation creates a new string
	// so ti can lead to inefficient
	res := make([]byte, length)
	index := 0
	for index < count {
		res[index] = '1'
		index++
	}
	res[count] = '0'
	index = count + 1
	for index < length {
		res[index] = '1'
		index++
	}
	return string(res)
}

// timeout
// we can find the patterns from example
// 01110 -> 01101 -> 01011 -> 00111 -> 10111
// if start with '0' and next position is '0',we can replace '00' to '10'
// if next position is '1', skip it , and find the next '0' position ,
// then we replace the first position to '1' ,the next postion to '0', the last position to '1',like '01110' to '10111'
func maximumBinaryString(binary string) string {
	// 00 -> 10 -> 01
	charArray := []rune(binary)
	i := 0
	for i < len(charArray) {
		c1 := charArray[i]
		end := i + 1
		if c1 == '0' {
			if i+1 < len(charArray) {
				if charArray[i+1] == '0' {
					charArray[i] = '1'
					i++
					continue
				}
			}
			for j := end; j < len(charArray); j++ {
				if charArray[j] == '1' {
					continue
				}
				end = j
				break
			}
			if end == i+1 {
				i++
				continue
			}
			for j := i; j <= end; j++ {
				charArray[j] = '1'
			}
			charArray[i+1] = '0'
		}
		i++
	}
	return string(charArray)
}
