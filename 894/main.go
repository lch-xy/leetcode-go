package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fbt := allPossibleFBT(9)
	println(fbt)
}

// 以空间换时间
var cache = make(map[int][]*TreeNode)

func allPossibleFBT(n int) []*TreeNode {
	if cache[n] != nil {
		return cache[n]
	}
	res := []*TreeNode{}
	// 如果是满二叉树 肯定是奇数
	if n%2 == 0 {
		return nil
	}
	if n == 1 {
		return append(res, &TreeNode{Val: 0})
	}

	// 将节点分配给左右两个子树 进行递归操作
	for i := 1; i < n; i = i + 2 {
		var left = allPossibleFBT(i)
		var right = allPossibleFBT(n - 1 - i)
		// 拿到左右节点存在的所有可能结果进行组合
		// n = 1时，直接返回，1种情况
		// n = 3时，左右各分配1，两个1进行组装也是1种情况
		// n = 5时，左右分配有2种情况（1，3）（3，1），每种可能有1种情况，共2种情况
		// n = 7时，左右分配有3种情况（1，5）（3，3）（5，1），（1，5）是1种和2种组装后是2种，（3，3）是1种和1种组装后是1种，所一共是2+1+2种
		// n = 9时，左右分配有5种情况（1，7）（3，5）（5，3）（7，1），一次类推为5+2+2+5种可能
		for _, l := range left {
			for _, r := range right {
				// 每次都要创建一个新的node，因为要求出所有的情况
				node := &TreeNode{Val: 0}
				node.Left = l
				node.Right = r
				res = append(res, node)
			}
		}
	}
	cache[n] = res
	return res
}
