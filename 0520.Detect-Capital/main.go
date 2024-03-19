package main

import "strings"

func main() {

	print(detectCapitalUse("String Apple"))

}

func detectCapitalUse(word string) bool {
	sp := strings.Split(word, " ")
	for i := range sp {
		cnt := 0
		index := -1
		for j := range sp[i] {
			if sp[i][j]-'a' > 26 {
				cnt++
				index = j
			}
		}
		if cnt != 0 && cnt != len(sp[i]) && index == 0 {
			return false
		}
	}
	return true
}
