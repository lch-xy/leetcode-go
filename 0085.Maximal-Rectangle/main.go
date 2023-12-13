package main

func main() {
	matrix := [][]byte{
		[]byte("10100010"),
		[]byte("10111100"),
		[]byte("10111101"),
	}
	println(maximalRectangle(matrix))
}

// 利用84题的解法
// 之前是将matrix重新组装成一个新的数组，存的以当前的节点为底的柱状图的面积，然后逐行拼成一个一维数组进行计算
// 其实不用这么麻烦 直接拼成一维数组即可
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	x := len(matrix)
	y := len(matrix[0])

	res := 0
	heights := make([]int, y)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			height := func() int {
				if int(matrix[i][j]-'0') == 1 {
					return heights[j] + 1
				}
				return 0
			}
			heights[j] = height()
		}
		res = max(res, largestRectangleAreaStack(heights))
	}
	return res
}

// 拼了一个新的二维数组 其实没必要
//func maximalRectangle(matrix [][]byte) int {
//	if len(matrix) == 0 {
//		return 0
//	}
//	x := len(matrix)
//	y := len(matrix[0])
//	// 初始化dp数组
//	newMatrix := make([][]int, x)
//	for i := 0; i < len(matrix); i++ {
//		newMatrix[i] = make([]int, y)
//	}
//
//	// 初始化第一行数据
//	for i := 0; i < y; i++ {
//		newMatrix[0][i] = int(matrix[0][i] - '0')
//	}
//
//	for i := 1; i < x; i++ {
//		for j := 0; j < y; j++ {
//			height := func() int {
//				if int(matrix[i][j]-'0') == 1 {
//					return newMatrix[i-1][j] + 1
//				}
//				return 0
//			}
//			newMatrix[i][j] = height()
//		}
//	}
//	res := 0
//	for i := 0; i < x; i++ {
//		heights := []int{}
//		for j := 0; j < y; j++ {
//			heights = append(heights, newMatrix[i][j])
//		}
//		res = max(res, largestRectangleAreaStack(heights))
//	}
//	return res
//}

func largestRectangleAreaStack(heights []int) int {
	res := 0
	stack := []int{}
	heights = append(heights, 0)

	size := len(heights)
	for i := 0; i < size; i++ {
		if len(stack) == 0 || heights[stack[len(stack)-1]] < heights[i] {
			stack = append(stack, i)
		} else {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			height := heights[index]
			length := func() int {
				if len(stack) == 0 {
					return i
				}
				return i - stack[len(stack)-1] - 1
			}
			res = max(res, height*length())
			i--
		}
	}

	return res
}
