package main

func main() {

	//println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	//println(largestRectangleArea([]int{2, 4}))
	println(largestRectangleArea([]int{1}))

}

// 最后几个case超时了
// 遍历找到一个峰值，然后就以这个峰值为起点，向前遍历，计算最大
// 这个方法相对暴力搜索而言，进行了大量的剪枝
// 因为是一个递增序列的话，那么加上后面的的面积肯定是比前面要打的
func largestRectangleArea(heights []int) int {
	res := 0
	size := len(heights)
	for i := 0; i < size; i++ {
		if i+1 < size && heights[i] <= heights[i+1] {
			continue
		}
		minHeight := heights[i]
		for j := 0; j <= i; j++ {
			// 从后往前遍历
			minHeight = min(minHeight, heights[i-j])
			res = max(res, minHeight*(j+1))
		}
	}
	return res
}
