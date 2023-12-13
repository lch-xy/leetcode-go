package main

func main() {

	//println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	//println(largestRectangleArea([]int{2, 4}))
	//println(largestRectangleArea([]int{1}))

	println(largestRectangleAreaStack([]int{2, 1, 5, 6, 2, 3}))
	println(largestRectangleArea([]int{2, 4}))
	println(largestRectangleArea([]int{1}))

}

// 使用单调递增栈，碰到递减的元素就停下来计算面积
// 意思就是说：只要是单调递增，就能以任意一边为高度，一直向右滑到底，来计算面积
// 因为右边的值比你大，肯定能包括你
func largestRectangleAreaStack(heights []int) int {
	res := 0
	stack := []int{}
	// 为了让最后一个元素入栈，所以多添加了一个为面积为0的柱子
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
				// 为什么是需要-1？
				// 我们以[2, 1, 5, 6, 2, 3]为例
				// 当i=4时，此时stack为[2,5]，因为6在上面已经出栈了，而i代表的位置又在6的后面，所以要-1
				// 此时算出来的面积是柱子6的单个面积
				return i - stack[len(stack)-1] - 1
			}
			res = max(res, height*length())
			// 因为我们只是在stack里面自己玩，并没有处理heights[i]
			// 所以要回去重新处理heights[i]
			i--
		}
	}

	return res
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
