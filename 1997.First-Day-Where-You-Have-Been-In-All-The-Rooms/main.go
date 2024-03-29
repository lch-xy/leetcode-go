package main

import "fmt"

func main() {
	// Define test cases slice
	testCases := []struct {
		nextVisit []int // 房间序列
		expected  int   // 预期的首次进入所有房间的天数
	}{
		// {nextVisit: []int{0, 0, 2}, expected: 6},
		{nextVisit: []int{0, 1, 2, 0}, expected: 6},
		// You can add more test cases here
	}

	// Iterate through test cases and run tests
	for i, tc := range testCases {
		result := firstDayBeenInAllRoomsDp(tc.nextVisit)
		if result != tc.expected {
			fmt.Printf("Test case %d failed: Expected output %d, got %d\n", i+1, tc.expected, result)
		}
	}
}

// dp[i]: the first time arrive this room
func firstDayBeenInAllRoomsDp(nextVisit []int) int {
	mod := 1000000007
	dp := make([]int, len(nextVisit))
	for i := 1; i < len(nextVisit); i++ {
		// when we first visit dp[i], we must come from dp[i-1], and all rooms before i-1 have been visited even times
		// when we first visit dp[i-1], we will return to nextVisit[i-1], and then the second time to dp[i-1], we will go to dp[i] next step
		// so we use dp[i-1] - dp[nextVisit[i-1]] + 1 to caculate the time spend of nextVisit[i-1] to dp[i-1]
		// because the first time we visit dp[i-1] , we will go to nextVisit[i-1] next, then go to dp[i-1] again
		// "+ mod" is to  prevent negative number
		dp[i] = (dp[i-1] + 1 + dp[i-1] - dp[nextVisit[i-1]] + 1 + mod) % mod
	}
	return dp[len(nextVisit)-1]
}

func firstDayBeenInAllRooms(nextVisit []int) int {
	mod := 1000000007
	total := len(nextVisit)
	set := make(map[int]struct{})
	visit := make([]int, total)
	cnt, days, index := 0, 0, 0
	for cnt != total {
		visit[index]++
		if _, ok := set[index]; !ok {
			set[index] = struct{}{}
			cnt++
		}
		if visit[index]%2 == 0 {
			index = (index + 1) % total
		} else {
			index = nextVisit[index]
		}
		days++
	}

	return (days - 1) % mod
}
