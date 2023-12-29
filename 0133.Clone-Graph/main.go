package main

import "reflect"

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func main() {
	// 创建原始图
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	// 克隆图
	clonedNode := cloneGraph(node1)

	// 检查克隆的图是否与原图相同
	if !reflect.DeepEqual(node1, clonedNode) {
		println("error")
	}
}

// 创建的map类型应该是指针类型
var cache = map[*Node]*Node{}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	// 防止成环 遍历过的节点不会再遍历了
	if value, ok := cache[node]; ok {
		return value
	}
	cur := &Node{Val: node.Val}
	cache[node] = cur

	for _, v := range node.Neighbors {
		// 这里要添加返回的克隆值cloneGraph(v)，而不是v
		cur.Neighbors = append(cur.Neighbors, cloneGraph(v))
	}
	return cur
}
