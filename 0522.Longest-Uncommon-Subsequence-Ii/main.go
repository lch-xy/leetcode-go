package main

import "fmt"

func main() {
	fmt.Println(findLUSlength([]string{"abc", "bcd", "abcd", "efg"}))
	fmt.Println(findLUSlength([]string{"aba", "cdc", "eae"}))
	fmt.Println(findLUSlength([]string{"aaa", "aaa", "aa"}))
	fmt.Println(findLUSlength([]string{"aabbcc", "aabbcc", "c"}))
}

// we need to enumerate all subsequence?
// not necessary , if the shortest subsequence is special subsequence , the longer one alse is .
// so we only need to iterate all string , comparing with others.
func findLUSlength(strs []string) int {
	maxLength := -1
	for i := 0; i < len(strs); i++ {
		isUncommon := true
		for j := 0; j < len(strs); j++ {
			if i != j && isSubquence(strs[i], strs[j]) {
				isUncommon = false
				break
			}
		}
		if isUncommon {
			maxLength = max(maxLength, len(strs[i]))
		}
	}
	return maxLength
}

// a is b's subsequence
func isSubquence(a string, b string) bool {
	j := 0
	for i := 0; i < len(b) && j < len(a); i++ {
		if a[j] == b[i] {
			j++
		}
	}
	return j == len(a)
}
