package main

import "fmt"

func main() {
	testCases := []struct {
		name  string
		boxes []int
		want  int
	}{
		{
			name:  "Case 1: Single box",
			boxes: []int{1},
			want:  1,
		},
		{
			name:  "Case 2: Two different boxes",
			boxes: []int{1, 2},
			want:  2,
		},
		{
			name:  "Case 3: Two same boxes",
			boxes: []int{1, 1},
			want:  4,
		},
		{
			name:  "Case 4: Three boxes, two are the same",
			boxes: []int{1, 2, 2},
			want:  5,
		},
	}

	for _, tc := range testCases {
		got := removeBoxes(tc.boxes)
		if got != tc.want {
			fmt.Printf("removeBoxes(%v) = %v; want %v", tc.boxes, got, tc.want)
		}
	}
}

func removeBoxes(boxes []int) int {
	n := len(boxes)
	dp := make([][100][100]int, n)

	return calculatePoints(boxes, dp, 0, n-1, 0)
}

// 这里我们需要升维，用到三维数组，盒子的移除不仅仅与左右边界有关，还与边界外同颜色盒子的数量有关。
// dp[l][r][k] 表示在子数组 boxes[l] 到 boxes[r] 中，右侧还有 k 个与 boxes[r] 相同的颜色盒子
func calculatePoints(boxes []int, dp [][100][100]int, l, r, k int) int {
	if l > r {
		return 0
	}

	// boxes数组的处理，遵循的策略是尽可能让区间右边颜色相同的盒子更多
	for r > l && boxes[r] == boxes[r-1] {
		r--
		k++
	}

	// 这一步不能提前 不然会超时
	if dp[l][r][k] != 0 {
		return dp[l][r][k]
	}

	// 第一种选择是移除这些颜色的所有盒子
	// 为什么是k+1？k是后面连续的字符，还需要加上r上的字符，所以是k+1个字符
	// 为什么区间是[l,r-1]？因为在上面的for循环已经将r处理过了，这里只需要减去当前和后面相同的数值即可，所以是r-1
	dp[l][r][k] = calculatePoints(boxes, dp, l, r-1, 0) + (k+1)*(k+1)
	// 第二种选择是保留一部分颜色相同的盒子，试图把它们与后面相同颜色的盒子一并消除，可能获得更高的得分。
	for i := l; i < r; i++ {
		// 我们是找到与boxes[r]相同颜色的的盒子，然后一起消除
		if boxes[i] == boxes[r] {
			// calculatePoints(boxes, l, i, k+1)代表，和后面的k+1个元素一起消除，为什么是k+1？因为boxes[r]也是
			// calculatePoints(boxes, i+1, r-1, 0)代表，将中间的元素先消除掉，因为boxes[i]和boxes[r]都是和后面一样的，所以区间是[i+1,r-1]
			dp[l][r][k] = max(dp[l][r][k], calculatePoints(boxes, dp, l, i, k+1)+calculatePoints(boxes, dp, i+1, r-1, 0))
		}
	}

	return dp[l][r][k]
}

// // 下面是错误的解法
// // dp[i][j]: boxes[i] -> boxes[j]之间你能获得的最大积分和
// func removeBoxes(boxes []int) int {
// 	size := len(boxes)
// 	if size == 0 {
// 		return 0
// 	}
// 	dp := make([][]int, size)
// 	for i := range dp {
// 		dp[i] = make([]int, size)
// 		dp[i][i] = 1
// 	}

// 	for lenght := 2; lenght <= size; lenght++ {
// 		for i := 0; i <= size-lenght; i++ {
// 			j := i + lenght - 1
// 			for k := i; k <= j; {
// 				start := k
// 				end := k
// 				cnt := 1
// 				for end+1 <= j && boxes[end] == boxes[end+1] {
// 					cnt++
// 					end++
// 				}
// 				if start > 0 && end < j {
// 					dp[i][j] = max(dp[i][j], dp[i][start-1]+cnt*cnt+dp[end+1][j])
// 				} else if start > 0 {
// 					dp[i][j] = max(dp[i][j], dp[i][start-1]+cnt*cnt)
// 				} else if end < j {
// 					dp[i][j] = max(dp[i][j], cnt*cnt+dp[end+1][j])
// 				} else {
// 					dp[i][j] = cnt * cnt
// 				}
// 				k = end + 1
// 			}
// 		}
// 	}

// 	return dp[0][size-1]
// }
