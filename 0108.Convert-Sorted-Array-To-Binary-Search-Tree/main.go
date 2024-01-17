package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 这题的结果会有很多种，测试用例可能不准确 但是可以ac
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	// 每次取中间的当跟节点一定能保证左右子树相差不超过1
	mid := start + (end-start)/2
	left := helper(nums, start, mid-1)
	right := helper(nums, mid+1, end)
	return &TreeNode{nums[mid], left, right}
}
