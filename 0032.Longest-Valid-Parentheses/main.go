package main

func main() {
	println(longestValidParentheses(")()()()()())())((()(()()()"))
	println(longestValidParentheses("(()"))
	println(longestValidParentheses("()(()"))
	println(longestValidParentheses("()(()())"))

	println("==============================================")

	println(longestValidParenthesesDp(")()()()()())())((()(()()()"))
	println(longestValidParenthesesDp("(()"))
	println(longestValidParenthesesDp("()(()"))
	println(longestValidParenthesesDp("()(()())"))

	println("==============================================")

	println(longestValidParenthesesDoubleTraversal(")()()()()())())((()(()()()"))
	println(longestValidParenthesesDoubleTraversal("(()"))
	println(longestValidParenthesesDoubleTraversal("()(()"))
	println(longestValidParenthesesDoubleTraversal("()(()())"))
}

// 为什么用stack比较合适？而不是一次for循环直接找出来？
// 因为可能存在n中可能，如果一次for循环就需要记住n个index
// 但是使用stack，我们直接将index入栈即可，长度就是两个index的差值，会方便很多
func longestValidParentheses(s string) int {
	stack := []int{}
	res := 0
	start := 0
	for i := range s {
		if s[i] == '(' {
			// 当是"("的时候，我们做任何处理
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				// 这里为什么是i+1?
				// 因为走到这里肯定是")"，而且又没匹配上，所以需要把start移到下一位，也就是下一个可能开始的地方
				// 那为什么要在这里写逻辑呢？
				// 因为当是"("的时候，我们不能确定这个start是否需要移动的，所以放到")"的时候去更新比较合适
				start = i + 1
			} else {
				// 说明匹配上了，将匹配上的括号弹出去
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					// 说明已经到start的位置了
					res = max(res, i-start+1)
				} else {
					// 说明只是匹配了中间一小段
					// 中间有多段也没关系，因为是用当前index和stack里最后一个元素的index相减
					res = max(res, i-stack[len(stack)-1])
				}
			}
		}
	}
	return res
}

// s:			)	(	)	(	)	(
// index:		0	1	2	4	5	5
// =========================================>
// dp:		0	0	0	2	0	4	0
// index:	0	1	2	3	4	5	6
// =========================================>
// dp[i]：以s[i-1]结尾的字符串最长括号的长度
// 为什么要这么定义？
// 因为我们计算当前的长度的时候需要以上一步的最大结果做比较
// 例如:我们要计算dp[6]，我们需要心啊知道dp[5]的长度，因为dp[5]的长度是4
// 那么我们可再向前推4为，通过公式6-dp[5]-2可以得到index为0的节点
// 为什么要-2？
// 因为当前节点6肯定不能算进去，而且我们定位是s的位置，相对于dp的index还需要减一
// 所以我们在dp数组上多加了一个长度，dp[i]从1开始，对应的是s[0]
func longestValidParenthesesDp(s string) int {
	dp := make([]int, len(s)+1)
	res := 0
	// 因为dp[i]对应是以s[0]结尾的字符串，所以index从1开始
	for i := 1; i <= len(s); i++ {
		// 这个j是s对应的index
		j := i - dp[i-1] - 2
		// 如果当前位置为'('，或者向前推n个位置后是')'，以及计算的j<0，都是无法匹配的
		if j < 0 || s[j] == ')' || s[i-1] == '(' {
			dp[i] = 0
		} else {
			// 当期的最大值 = 以上个字符结尾的最大值 + 当前匹配的括号 + 更早的连续的值
			// 因为j是s的index，dp[j]表示的是以s[j-1]结尾的字符串最长括号的长度，还要再往前一位
			dp[i] = dp[i-1] + 2 + dp[j]
			res = max(res, dp[i])
		}
	}
	return res
}

// 这个方法很难想到，最开始我也是想到用遍历的方式去做，但是没过ac,因为我只定义了一个cnt记录长度
// 这里定义左右指针，遇到后进行++操作，遇到非法情况，将left和right重置为0
// 这样做会对"(()"这种没办法，因为是匹配到"(())"后才会觉得没问题
// 所以可以从后往前再遍历一次，这样就可以找来了
func longestValidParenthesesDoubleTraversal(s string) int {
	left := 0
	right := 0
	res := 0
	size := len(s)
	for i := 0; i < size; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			res = max(res, 2*right)
		}
		if left < right {
			left = 0
			right = 0
		}
	}
	left = 0
	right = 0
	for i := size - 1; i > 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			res = max(res, 2*left)
		}
		if left > right {
			left = 0
			right = 0
		}
	}
	return res
}
