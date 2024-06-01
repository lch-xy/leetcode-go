package main

import "fmt"

func main() {
	fmt.Printf("distributeCandies(12, 24): %v\n", distributeCandies(12, 24))
}
func distributeCandies(n int, limit int) int {
	cnt := 0
	for i := 0; i <= min(limit, n); i++ {
		if n-i > 2*limit {
			continue
		}
		cnt += min(n-i, limit) - max(0, n-i-limit) + 1
	}
	return cnt
}
