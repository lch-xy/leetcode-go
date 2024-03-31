package main

import (
	"fmt"
	"strings"
)

func main() {
	testCases := []struct {
		preorder string
		expected bool
	}{
		{preorder: "9,3,4,#,#,1,#,#,2,#,6,#,#", expected: true},
		{preorder: "1,#", expected: false},
		{preorder: "9,#,#,1", expected: false},
		{preorder: "9,#,#,1,#,#", expected: false},
		{preorder: "1,#,#,#,#", expected: false},
	}

	for i, tc := range testCases {
		result := isValidSerialization(tc.preorder)
		if result != tc.expected {
			fmt.Printf("Test case %d failed: Expected %t, got %t\n", i+1, tc.expected, result)
		}
	}
}

// use stack to simulate preorder traversal
// key point: regard "#" as a leaf node
// so we meet "number,#,#", we can use "#" to replace it
// if there is only "#" left in the stack, we return true, otherwise we return false
func isValidSerialization(preorder string) bool {
	stack := make([]string, 0)

	nodes := strings.Split(preorder, ",")
	for _, node := range nodes {
		stack = append(stack, node)
		for len(stack) >= 3 && stack[len(stack)-1] == "#" && stack[len(stack)-2] == "#" && stack[len(stack)-3] != "#" {
			stack = stack[:len(stack)-3]
			stack = append(stack, "#")
		}
	}
	if len(stack) == 1 && stack[0] == "#" {
		return true
	}
	return false
}
