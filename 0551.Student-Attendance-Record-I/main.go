package main

import "strings"

func checkRecord(s string) bool {
	cnt := 0
	for i := range s {
		if s[i] == 'A' {
			cnt++
		}
	}
	if cnt >= 2 {
		return false
	}

	newStrOne := strings.Replace(s, "LLL", "", -1)
	if len(newStrOne) != len(s) {
		return false
	}

	return true
}
