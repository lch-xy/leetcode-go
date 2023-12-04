package main

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// [[[-1,null,-1]],[1],[2]]
	root := &TreeNode{Val: -1}
	root.Right = &TreeNode{Val: -1}
	Constructor(root)
}

// 将节点序列化成一个集合
type FindElements struct {
	list []int
}

// 通过递归 序列化
func Constructor(root *TreeNode) FindElements {
	res := []int{}
	helper(root, &res, 0)
	return FindElements{list: res}
}

func helper(root *TreeNode, res *[]int, curVal int) {
	if root == nil {
		return
	}
	*res = append(*res, curVal)
	helper(root.Left, res, curVal*2+1)
	helper(root.Right, res, curVal*2+2)
}

func (this *FindElements) Find(target int) bool {
	for _, v := range this.list {
		if v == target {
			return true
		}
	}
	return false
}
