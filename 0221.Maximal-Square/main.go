package main

func main() {
	matrix := [][]byte{
		[]byte("10100"),
		[]byte("10111"),
		[]byte("11111"),
		[]byte("10101"),
	}
	println(maximalSquare(matrix))

	matrix = [][]byte{
		[]byte("11111"),
		[]byte("11111"),
		[]byte("00001"),
		[]byte("11111"),
		[]byte("11111"),
	}
	println(maximalSquare(matrix))

	//println(largestSquareAreaStack([]int{2, 1, 5, 6, 2, 3}))
}

// 思路和84题最大长方形一样
// 我们在算面积的时候 面积 = 长 * 宽
// 正方形两边相等取最小值即可
func maximalSquare(matrix [][]byte) int {
	x := len(matrix)
	y := len(matrix[0])
	if x == 0 || y == 0 {
		return 0
	}
	res := 0
	heights := make([]int, y)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if (matrix[i][j] - '0') == 1 {
				heights[j] = heights[j] + 1
			} else {
				heights[j] = 0
			}
		}
		res = max(res, largestSquareAreaStack(heights))
	}
	return res
}

func largestSquareAreaStack(heights []int) int {
	size := len(heights)
	stack := []int{}
	heights = append(heights, 0)

	res := 0
	for i := 0; i <= size; i++ {
		if len(stack) == 0 || heights[stack[len(stack)-1]] < heights[i] {
			stack = append(stack, i)
		} else {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			length := func() int {
				if len(stack) == 0 {
					return i
				} else {
					return i - stack[len(stack)-1] - 1
				}
			}
			res = max(res, min(length(), heights[index])*min(length(), heights[index]))
			i--
		}
	}
	return res
}
