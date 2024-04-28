package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("baseNeg2(10): %v\n", baseNeg2(10))
}

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	for n != 0 {
		site := n & 1
		res = strconv.Itoa(site) + res
		n = -(n >> 1)
	}
	return res
}
