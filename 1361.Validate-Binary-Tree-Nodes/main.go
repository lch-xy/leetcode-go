package main

// Definition for a binary nodeList node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//root := &TreeNode{Val: 1}
	//root.Left = &TreeNode{Val: 2}
	//root.Right = &TreeNode{Val: 3}
	//root.Left.Left = &TreeNode{Val: 4}
	//root.Left.Right = &TreeNode{Val: 5}
	//root.Right.Left = &TreeNode{Val: 6}
	//root.Right.Right = &TreeNode{Val: 7}
	//root.Left.Left.Left = &TreeNode{Val: 8}
	//root.Left.Left.Right = &TreeNode{Val: 9}
	//root.Left.Right.Left = &TreeNode{Val: 10}
	//root.Left.Right.Right = &TreeNode{Val: 11}

	// n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]
	println(validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}))
	// n = 6, leftChild = [1,-1,-1,4,-1,-1], rightChild = [2,-1,-1,5,-1,-1] false
	println(validateBinaryTreeNodes(6, []int{1, -1, -1, 4, -1, -1}, []int{2, -1, -1, 5, -1, -1}))
}

// 先算出每个节点的入度
// 遍历数组找出入度为0的节点，为根节点
// 如果存在入度大于1的节点，或者多个入度为0的节点，直接return false
// 找出根节点，递归算出整个节点数，看是否等于n
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	inDegree := make([]int, n)
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			inDegree[leftChild[i]]++
		}
		if rightChild[i] != -1 {
			inDegree[rightChild[i]]++
		}
	}

	rootIndex := -1
	for i := 0; i < n; i++ {
		if inDegree[i] > 1 {
			return false
		}
		if inDegree[i] != 0 {
			continue
		}
		// 如果两次过来 说明存在两个root节点
		if rootIndex != -1 {
			return false
		}
		rootIndex = i
	}

	cnt := caculate(rootIndex, leftChild, rightChild)
	if cnt != n {
		return false
	}

	return true
}

func caculate(rootIndex int, leftChild []int, rightChild []int) int {
	if rootIndex == -1 {
		return 0
	}
	left := caculate(leftChild[rootIndex], leftChild, rightChild)
	right := caculate(rightChild[rootIndex], leftChild, rightChild)
	return left + right + 1
}
