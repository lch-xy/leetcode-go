package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 和116一模一样的解法
// 层序遍历不需要管是不是属于同一个父节点
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{}
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		var pre *Node
		for i := 0; i < size; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
			if i > 0 {
				pre.Next = curNode
			}
			pre = curNode
		}
	}
	return root
}
