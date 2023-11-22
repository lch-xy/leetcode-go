package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorder -> [1] [2,4,5] [3,6,7]
// postorder -> [4,5,2] [6,7,3] [root]
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	return helper(preorder, postorder, 0, len(preorder)-1, 0, len(postorder)-1)
}

func helper(preorder []int, postorder []int, preL int, preR int, postL int, postR int) *TreeNode {
	if preL > preR || postL > postR {
		return nil
	}
	root := &TreeNode{preorder[preL], nil, nil}
	if preL == preR {
		return root
	}
	idx := -1
	for idx = postL; idx < postR; idx++ {
		if postorder[preL+1] == postorder[idx] {
			break
		}
	}

	root.Left = helper(preorder, postorder, preL+1, postL+1+idx-postL, postL, idx)
	root.Right = helper(preorder, postorder, postR-1-idx-1, preR, idx+1, postR-1)

	return root
}
