package main

import (
	"fmt"
	"sort"
)

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// [3,1,4,0,2,2]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 7}
	// [0,8,1,null,null,3,2,null,4,5,null,null,7,6]
	//root := &TreeNode{Val: 0}
	//root.Left = &TreeNode{Val: 8}
	//root.Right = &TreeNode{Val: 1}
	//root.Right.Left = &TreeNode{Val: 3}
	//root.Right.Right = &TreeNode{Val: 2}
	//root.Right.Left.Right = &TreeNode{Val: 4}
	//root.Right.Right.Left = &TreeNode{Val: 5}
	//root.Right.Left.Right.Right = &TreeNode{Val: 7}
	//root.Right.Right.Left.Left = &TreeNode{Val: 6}

	verticalTraversal(root)
}

// 虽然ac 但是代码可读性不强
// 递归函数很简单，构建一个map，key为x的坐标，value为一个List，里面也是个map，存放了所有 y 坐标和对应的节点值
// 麻烦的地方就是需要根据 x 坐标排序，然后根据 y 坐标排序，把对应的值组织到 [][]int 中
func verticalTraversal(root *TreeNode) [][]int {
	var collect = map[int][]map[int]int{}
	helper(root, 0, 0, collect)

	// 获取 map 的 key，并排序
	keys := make([]int, 0, len(collect))
	for key := range collect {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// 根据排序后的 key 顺序组织数据到 [][]int 中
	result := make([][]int, 0, len(collect))
	for _, key := range keys {
		subCollect := collect[key]
		subKeys := make([]int, 0, len(subCollect))
		for _, v := range subCollect {
			for k := range v {
				if !contain(subKeys, k) {
					subKeys = append(subKeys, k)
				}
			}
		}
		sort.Ints(subKeys)
		aa := make([]int, 0, len(subCollect))
		for _, subKey := range subKeys {
			ss := []int{}
			for _, v := range subCollect {
				value, exists := v[subKey]
				if exists {
					ss = append(ss, value)
				}
			}
			sort.Ints(ss)
			for _, vv := range ss {
				aa = append(aa, vv)
			}
		}

		result = append(result, aa)
	}
	fmt.Println(result)
	return result
}

func helper(root *TreeNode, x int, y int, collect map[int][]map[int]int) {
	if root == nil {
		return
	}
	list := collect[x]
	list = append(list, map[int]int{y: root.Val})
	collect[x] = list
	helper(root.Left, x-1, y+1, collect)
	helper(root.Right, x+1, y+1, collect)
}

func contain(collect []int, target int) bool {
	for _, v := range collect {
		if v == target {
			return true
		}
	}
	return false
}
