package main

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// [1,2,null,3]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	//root.Right.Left = &TreeNode{Val: 4}
	//root.Right.Right = &TreeNode{Val: 5}

	//preOrder := []int{1}
	preOrder := []int{1, 3, 2}
	//preOrder := []int{1, 2, 3, 4, 5}
	for _, v := range flipMatchVoyage(root, preOrder) {
		println(v)
	}
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	var res = []int{}
	var i = 0
	var rpt = &i
	b := helper(root, voyage, rpt, &res)
	if !b {
		return []int{-1}
	}
	return res
}

// 不要去判断左右子树是否为空，可以直接放到递归最前面判断 如果为空就返回true
// 不要尝试去吧左右子树划分出来，然后递归，这样很难，因为一棵树的形状靠前序遍历是不够的，可能存在的情况太多太多了
// 全局定义一个i，类似与index，正好 根-左-右 就是前序遍历的结果
// 只需要一个id，我们直接递归去和voyage数组去对比
func helper(root *TreeNode, voyage []int, i *int, res *[]int) bool {
	if root == nil {
		return true
	}
	if root.Val != voyage[*i] {
		return false
	}
	*i++

	if root.Left != nil && root.Left.Val != voyage[*i] {
		*res = append(*res, root.Val)
		return helper(root.Right, voyage, i, res) &&
			helper(root.Left, voyage, i, res)
	}

	return helper(root.Left, voyage, i, res) &&
		helper(root.Right, voyage, i, res)
}

// not ac，想着把左右子树分出来，这思路是错的
//func helper(root *TreeNode, voyage []int, start int, end int, res *[]int) bool {
//	if root == nil || start == end {
//		return true
//	}
//	if start < 0 || end < 0 || start > end {
//		return false
//	}
//	val := root.Val
//	if voyage[start] != val {
//		return false
//	}
//	start++
//	if root.Left != nil && root.Left.Val != voyage[start] {
//		*res = append(*res, val)
//		index := findIndex(voyage, root.Left)
//		return helper(root.Right, voyage, start, index-1, res) &&
//			helper(root.Left, voyage, index, end, res)
//	}
//	if root.Right != nil && root.Right.Val != voyage[start] {
//		index := findIndex(voyage, root.Right)
//		return helper(root.Left, voyage, start, index-1, res) &&
//			helper(root.Right, voyage, index, end, res)
//	}
//	return true
//}
//func findIndex(voyage []int, root *TreeNode) int {
//	target := root.Val
//	for i := range voyage {
//		if voyage[i] == target {
//			return i
//		}
//	}
//	return -1
//}
