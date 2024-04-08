package main

func main() {
	tree := Constructor(6, []int{1, 2, 3, 4, 5, 0})
	print(tree.GetKthAncestor(1, 4))
}

type TreeAncestor struct {
	dp [][]int
}

// we not need to store all result , only use arrays dp[i] to save 2^i-th ancestor
// jump up in powers of two
// reducde time complexity from  O(n) to log(n)
func Constructor(n int, parent []int) TreeAncestor {
	dp := make([][]int, n)
	for i := range dp {
		// 2^18 >  5 * 10^4`
		dp[i] = make([]int, 18)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = parent[i]
	}

	for j := 1; j < 18; j++ {
		for i := 0; i < n; i++ {
			if dp[i][j-1] != -1 {
				dp[i][j] = dp[dp[i][j-1]][j-1]
			} else {
				dp[i][j] = -1
			}
		}
	}

	tree := TreeAncestor{
		dp: dp,
	}
	return tree
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	for i := 17; i >= 0; i-- {
		// check the i-th position is 1
		if (k >> i & 1) == 1 {
			node = this.dp[node][i]
			if node == -1 {
				break
			}
		}
	}
	return node
}
