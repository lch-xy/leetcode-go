package main

import "math"

func main() {
	print(minSkips([]int{1, 3, 2}, 4, 2))
}

// dp[i][j]: the minimum time to reach the end of rhe i-th road with j skips
func minSkips(dist []int, speed int, hoursBefore int) int {
	length := len(dist)
	dp := make([][]float64, length+1)

	for i := range dp {
		dp[i] = make([]float64, length+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxFloat64
		}
	}

	dp[0][0] = 0
	// to avoid precision issues
	epsilon := 1e-9
	for i := 1; i <= length; i++ {
		for j := 0; j <= i; j++ {
			// not skip
			if j < i {
				dp[i][j] = min(dp[i][j], math.Ceil(dp[i-1][j]-epsilon)+float64(dist[i-1])/float64(speed))
			}

			// skip
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+float64(dist[i-1])/float64(speed))
			}
		}
	}

	for k := 0; k <= length; k++ {
		if dp[length][k] <= float64(hoursBefore)+float64(epsilon) {
			return k
		}
	}

	return -1
}
