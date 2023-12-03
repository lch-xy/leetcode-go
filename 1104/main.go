package main

func main() {
	pathInZigZagTree(14)
}

// 父节点.val = (该层的最大结点值 + 最小结点值 - 当前结点值并 ) / 2
func pathInZigZagTree(label int) []int {
	level := 0
	// 计算在节点在第几层
	for 1<<level <= label {
		level++
	}
	res := []int{}
	for label >= 1 {
		res = append([]int{label}, res...)
		label = (1<<level - 1 + 1<<(level-1) - 1 + 1 - label) / 2
		level--
	}
	return res
}
