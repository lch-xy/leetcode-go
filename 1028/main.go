package main

import (
	"strconv"
	"strings"
)

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preorder := recoverFromPreorder2("3-5--10---10----6--4-6")
	println(preorder)
}

// 利用栈的思想
// 出栈的条件是当前栈的数量 > 当前的层数,为什么要这么设计？
// 因为我们要保证当前节点一定是stack[len(stack)-1]的子树
// 先序遍历是按照 中-左-右 的顺序，而且题目说一定保证左子树优先
// 我们可以用countLeadingHyphens算出当前的level，如果有左右子树的话，那么len(stack)肯定大于level的
// 例如:3-5-6：
//
//	1:len(stack) = 0 ，level = 0， len(stack) > level = false ，入栈 [3]
//	2:len(stack) = 1 ，level = 1， len(stack) > level = false ，入栈 [3,5],先将5放到3的下面，再入栈
//	2:len(stack) = 2 ，level = 1， len(stack) > level = true ，出栈 [3]，先将6放到3的下面，再入栈 [3，6]
//
// 我们最后直接返回stack[0]即可
func recoverFromPreorder2(traversal string) *TreeNode {
	stack := []*TreeNode{}

	size := len(traversal)
	for i := 0; i < size; {

		level := countLeadingHyphens(traversal)
		traversal = traversal[level:]
		i += level
		notHyphens := countNotLeadingHyphens(traversal)
		i += notHyphens
		num := traversal[:notHyphens]
		traversal = traversal[notHyphens:]

		atoi, err := strconv.Atoi(num)
		if err != nil {
			return nil
		}
		node := &TreeNode{Val: atoi}

		// 当前栈的数量 > 当前的层数
		for len(stack) > level {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			if stack[len(stack)-1].Left == nil {
				stack[len(stack)-1].Left = node
			} else {
				stack[len(stack)-1].Right = node
			}
		}

		stack = append(stack, node)
	}

	return stack[0]
}

// 思路是将字符串分割成 中-左-右 三份，然后调用递归函数
func recoverFromPreorder(traversal string) *TreeNode {
	return helper(traversal, 0)
}

func helper(traversal string, level int) *TreeNode {
	if traversal == "" {
		return nil
	}
	wordCnt := countNotLeadingHyphens(traversal)

	atoi, err := strconv.Atoi(traversal[:wordCnt])
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: atoi}

	level++
	tmp := traversal[wordCnt:]
	left := ""
	right := ""
	for {
		cnt := countLeadingHyphens(tmp)
		if cnt == 0 {
			break
		}
		if cnt == level {
			if left == "" {
				left = tmp[cnt:]
			} else {
				right = tmp[cnt:]
			}
		}

		tmp = tmp[cnt:]
		cnt = countNotLeadingHyphens(tmp)
		tmp = tmp[cnt:]
	}

	if right != "" {
		lastIndex := strings.LastIndex(left, right)
		left = left[:lastIndex-1]
		root.Right = helper(right, level)
	}
	root.Left = helper(left, level)

	return root
}

func countLeadingHyphens(str string) int {
	count := 0
	for _, char := range str {
		if char == '-' {
			count++
		} else {
			break
		}
	}
	return count
}
func countNotLeadingHyphens(str string) int {
	count := 0
	for _, char := range str {
		if char != '-' {
			count++
		} else {
			break
		}
	}
	return count
}
