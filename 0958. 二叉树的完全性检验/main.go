package main

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{1, nil, nil}
	//root.Left = &TreeNode{2, nil, nil}
	//root.Right = &TreeNode{3, nil, nil}
	//root.Left.Left = &TreeNode{4, nil, nil}
	//root.Left.Right = &TreeNode{5, nil, nil}
	//root.Right.Right = &TreeNode{6, nil, nil}

	println(isCompleteTree(root))
}

// 优秀解法：层序遍历，添加左右节点，直接遇到一个空节点，说明后面的所以节点都是空节点
func isCompleteTree(root *TreeNode) bool {
	list := []*TreeNode{root}
	for list[0] != nil {
		node := list[0]
		list = append(list, node.Left)
		list = append(list, node.Right)
		list = list[1:]
	}
	for len(list) > 0 && list[0] == nil {
		list = list[1:]
	}
	return len(list) == 0
}

// 优化后的解法
//func isCompleteTree(root *TreeNode) bool {
//	list := []*TreeNode{root}
//	isEnd := false
//
//	for len(list) > 0 {
//		node := list[0]
//		if node == nil {
//			isEnd = true
//		} else {
//			if isEnd {
//				return false
//			}
//			list = append(list, node.Left)
//			list = append(list, node.Right)
//		}
//		list = list[1:]
//	}
//	return true
//}

// 待优化解法
//func isCompleteTree(root *TreeNode) bool {
//	list := []*TreeNode{root}
//	isEnd := false
//
//	for len(list) > 0 {
//		size := len(list)
//		for i := 0; i < size; i++ {
//			node := list[0]
//			if node.Left != nil {
//				if !isEnd {
//					list = append(list, node.Left)
//				} else {
//					return false
//				}
//			} else {
//				if !isEnd {
//					isEnd = true
//				}
//			}
//			if node.Right != nil {
//				if !isEnd {
//					list = append(list, node.Right)
//				} else {
//					return false
//				}
//			} else {
//				if !isEnd {
//					isEnd = true
//				}
//			}
//			list = list[1:]
//		}
//	}
//	return true
//}
