package main

import "fmt"

func main() {
	tests := []struct {
		letters []byte
		target  byte
		want    byte
	}{
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
		{[]byte{'c', 'f', 'j'}, 'c', 'f'},
		{[]byte{'c', 'f', 'j'}, 'd', 'f'},
		{[]byte{'c', 'f', 'j'}, 'g', 'j'},
		{[]byte{'c', 'f', 'j'}, 'j', 'c'},
		{[]byte{'c', 'f', 'j'}, 'k', 'c'},
	}

	for _, tt := range tests {
		got := nextGreatestLetter(tt.letters, tt.target)
		if got != tt.want {
			fmt.Printf("nextGreatestLetter(%v, %v) = %v; want %v", tt.letters, tt.target, got, tt.want)
		}
	}
}

func nextGreatestLetter(letters []byte, target byte) byte {
	size := len(letters)
	if target > letters[size-1] {
		return letters[0]
	}
	left, right := 0, size
	for left < right {
		mid := left + (right-left)/2
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if right == size {
		return letters[0]
	}
	return letters[left]
}
