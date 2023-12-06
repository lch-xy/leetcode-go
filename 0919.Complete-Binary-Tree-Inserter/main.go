package main

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{4, nil, nil}
	root.Left.Right = &TreeNode{5, nil, nil}
	root.Left.Left.Left = &TreeNode{6, nil, nil}
	constructor := Constructor(root)

	v := constructor.Insert(7)
	println(v)
	v = constructor.Insert(8)
	println(v)

	println(constructor.Get_root())
}

/**
 * 解法2，基于数组去做
 */
type CBTInserter struct {
	nodeList []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	nodeList := []*TreeNode{root}
	for i := 0; i < len(nodeList); i++ {
		if nodeList[i].Left != nil {
			nodeList = append(nodeList, nodeList[i].Left)
		}
		if nodeList[i].Right != nil {
			nodeList = append(nodeList, nodeList[i].Right)
		}
	}
	return CBTInserter{nodeList}
}
func (this *CBTInserter) Insert(val int) int {
	node := &TreeNode{Val: val}
	list := this.nodeList
	n := len(list)
	list = append(list, node)
	// 判断最后一个是左子树还是右子树
	if n%2 == 1 {
		list[(n-1)/2].Left = node
	} else {
		list[(n-1)/2].Right = node
	}
	this.nodeList = list
	return list[(n-1)/2].Val
}
func (this *CBTInserter) Get_root() *TreeNode {
	return this.nodeList[0]
}

/**
 * 解法1，基于TreeNode，基于层序遍历去做
 */
//type CBTInserter struct {
//	root_node *TreeNode
//	nodeList     []*TreeNode
//}
//func Constructor(root *TreeNode) CBTInserter {
//	nodeList := []*TreeNode{root}
//	for len(nodeList) > 0 {
//		topNode := nodeList[0]
//		if topNode.Left == nil || topNode.Right == nil {
//			break
//		}
//		nodeList = append(nodeList, topNode.Left)
//		nodeList = append(nodeList, topNode.Right)
//		nodeList = nodeList[1:]
//	}
//
//	return CBTInserter{root, nodeList}
//}
//
//func (this *CBTInserter) Insert(val int) int {
//	node := &TreeNode{Val: val}
//	curQueue := this.nodeList
//
//	curNode := curQueue[0]
//	if curNode.Left == nil {
//		curNode.Left = node
//	} else {
//		curNode.Right = node
//		curQueue = append(curQueue, curNode.Left)
//		curQueue = append(curQueue, curNode.Right)
//		curQueue = curQueue[1:]
//	}
//	this.nodeList = curQueue
//	return curNode.Val
//}
//
//func (this *CBTInserter) Get_root() *TreeNode {
//	return this.root_node
//}
