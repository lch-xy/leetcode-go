package main

func main() {

}

// 使用dfs
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	// 去掉了done数组 只使用一个visited数组，0 表示未访问 ，1 表示正在访问 ，-1 表示已访问
	visited := make([]int, numCourses)
	// 我们需要的是空数组
	// 所以这里要使用[]int{}，不能用make([]int, numCourses)
	res := []int{}

	for _, v := range prerequisites {
		graph[v[0]] = append(graph[v[0]], v[1])
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 && !dfs(graph, visited, &res, i) {
			return []int{}
		}
	}

	return res
}

func dfs(graph [][]int, visited []int, res *[]int, i int) bool {
	if visited[i] == 1 {
		return false
	}
	if visited[i] == -1 {
		return true
	}

	visited[i] = 1
	for _, j := range graph[i] {
		if !dfs(graph, visited, res, j) {
			return false
		}
	}
	visited[i] = -1
	// 为什么只需要再这里加一行就可以了？
	// 因为dfs的特性，会一直往里进行查找，最先走到这的肯定是最先要学的课程
	*res = append(*res, i)
	return true
}
