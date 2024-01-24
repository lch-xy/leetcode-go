package main

type NumMatrix struct {
	dp [][]int
}

// dp[i][j]: 从(0,0)到(i-1,j-1)围成的面积
func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{}
	}
	rows := len(matrix)
	cols := len(matrix[0])
	dp := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		dp[i] = make([]int, cols+1)
	}
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j] - dp[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{dp}
}

// 采用面积相减的方法，只需要O(1)的时间复杂度
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.dp[row2+1][col2+1] + this.dp[row1][col1+1] - this.dp[row1][col2+1] - this.dp[row2+1][col1]
}

//
// func Constructor(matrix [][]int) NumMatrix {
// 	return NumMatrix{matrix}
// }

// func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
// 	sumNumber := 0
// 	for i := row1; i <= row2; i++ {
// 		for j := col1; j <= col2; j++ {
// 			sumNumber += this.matrix[i][j]
// 		}
// 	}
// 	return sumNumber
// }

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
