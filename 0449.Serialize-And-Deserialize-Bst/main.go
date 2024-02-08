package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

type Codec struct {
	data []string
}

func Constructor() Codec {
	return *&Codec{[]string{}}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	serializeHelper(this, root)
	return strings.Join(this.data, " ")
}

func serializeHelper(this *Codec, root *TreeNode) {
	if root == nil {
		this.data = append(this.data, "#")
		return
	}
	this.data = append(this.data, strconv.Itoa(root.Val))
	serializeHelper(this, root.Left)
	serializeHelper(this, root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.data = strings.Split(data, " ")
	return this.deserializeHelper()
}

func (this *Codec) deserializeHelper() *TreeNode {
	if len(this.data) == 0 {
		return nil
	}
	curVal := this.data[0]
	this.data = this.data[1:]
	if curVal == "#" {
		return nil
	}
	val, _ := strconv.Atoi(curVal)
	left := this.deserializeHelper()
	right := this.deserializeHelper()
	return &TreeNode{val, left, right}
}
