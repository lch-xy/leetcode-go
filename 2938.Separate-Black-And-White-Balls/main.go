package main

import "fmt"

func main() {
	fmt.Println(minimumSteps("00"))
	fmt.Println(minimumSteps("100"))
	fmt.Println(minimumSteps("01110011101"))
	fmt.Println(minimumSteps("00000000011111111111"))
}

// we want to move all of the '0' to left
// first step : find the most left number '1' and find the most left number '0' after '1'
// second step : caculate the difference of two index
// third step : oneIndex++,because the number '1' has beed moved to the next position, we should find next number '0' to exchange position
// loop this operation until oneIndex or zeroIndex out of bound
func minimumSteps(s string) int64 {
	zeroIndex, oneIndex := -1, -1
	cnt := 0
	for i := 0; i < len(s); i++ {
		cur := s[i] - '0'
		if cur == 1 {
			oneIndex = i
			break
		}
	}
	if oneIndex < 0 {
		return int64(0)
	}
	for oneIndex < len(s) && zeroIndex < len(s) {
		for i := max(zeroIndex, oneIndex); i < len(s); i++ {
			cur := s[i] - '0'
			if cur == 0 {
				zeroIndex = i
				break
			}
		}
		if zeroIndex > 0 && s[zeroIndex]-'0' == 0 && zeroIndex > oneIndex {
			cnt += zeroIndex - oneIndex
			oneIndex++
		}
		zeroIndex++
	}
	return int64(cnt)
}
