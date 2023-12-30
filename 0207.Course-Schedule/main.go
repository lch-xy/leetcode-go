package main

import "fmt"

func main() {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		// {2, [][]int{{1, 0}}, true},
		// {2, [][]int{{1, 0}, {0, 1}}, false},
		// {3, [][]int{{0, 1}, {0, 2}, {1, 2}}, true},
		// {3, [][]int{{1, 0}, {2, 1}, {0, 2}}, false},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, true},
	}

	for _, test := range tests {
		result := canFinish(test.numCourses, test.prerequisites)
		if result != test.expected {
			fmt.Printf("Expected result to be %v, but got %v", test.expected, result)
		}
	}
}

// 思路就是看是否成环，如果成环就表示存在相互依赖的情况
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	visited := make([]int, numCourses)
	// 用来记录已经检查过的课程，防止重复检查
	// 如果拿掉会无法ac，时间超时
	done := make([]bool, numCourses)
	// 这里构造的是一个数组，相当于graph[0]里存放的就是课程0依赖的所有程
	// 也可以构造一个有向图，那么就需要将默认值设置为-1，因为0开始就是课程号了，并且dfs的时候碰到-1要返回true
	for _, v := range prerequisites {
		graph[v[0]] = append(graph[v[0]], v[1])
	}
	// 逐个课程使用dfs检查
	for i := 0; i < numCourses; i++ {
		if !done[i] && !dfs(graph, visited, done, i) {
			return false
		}
	}
	return true
}

func dfs(graph [][]int, visited []int, done []bool, i int) bool {
	if visited[i] == 1 {
		return false
	}
	visited[i] = 1
	for _, j := range graph[i] {
		if !dfs(graph, visited, done, j) {
			return false
		}
	}
	visited[i] = 0
	done[i] = true
	return true
}
