package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {
	tests := []struct {
		name string
		root *Node
		want *Node
	}{
		{
			name: "Test Case 1",
			root: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
					},
					Right: &Node{
						Val: 5,
					},
				},
				Right: &Node{
					Val: 3,
					Left: &Node{
						Val: 6,
					},
					Right: &Node{
						Val: 7,
					},
				},
			},
			want: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
						Next: &Node{
							Val: 5,
						},
					},
					Right: &Node{
						Val: 5,
						Next: &Node{
							Val: 6,
						},
					},
					Next: &Node{
						Val: 3,
					},
				},
				Right: &Node{
					Val: 3,
					Left: &Node{
						Val: 6,
						Next: &Node{
							Val: 7,
						},
					},
					Right: &Node{
						Val: 7,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		if got := connect(tt.root); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("connect() = %v, want %v", got, tt.want)
		}
	}
}

// 层序遍历
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
