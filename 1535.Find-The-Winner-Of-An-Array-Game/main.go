package main

import "fmt"

func main() {
	fmt.Printf("getWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2): %v\n", getWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2))
}

func getWinner(arr []int, k int) int {
	n := len(arr)
	if k >= n {
		return getMaxNumber(arr)
	}
	winner := -1
	winCnt := 0
	one, two := 0, 1
	for winCnt < k {
		if arr[one%n] > arr[two%n] {
			two = max(one, two) + 1
			if winner == one%n {
				winCnt++
			} else {
				winner = one % n
				winCnt = 1
			}
		} else {
			one = max(one, two) + 1
			if winner == two%n {
				winCnt++
			} else {
				winner = two % n
				winCnt = 1
			}
		}
	}

	return arr[winner]
}

func getMaxNumber(arr []int) int {
	maxNumber := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxNumber {
			maxNumber = arr[i]
		}
	}
	return maxNumber
}
