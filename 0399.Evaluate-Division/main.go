package main

import (
	"fmt"
	"reflect"
)

func main() {
	testCases := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		expected  []float64
	}{
		{
			equations: [][]string{{"a", "b"}, {"b", "c"}},
			values:    []float64{2.0, 3.0},
			queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			expected:  []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		// 在这里添加更多的测试用例
	}

	for _, tc := range testCases {
		result := calcEquation(tc.equations, tc.values, tc.queries)
		if !reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

// 根据已知条件先构造一个有向连通图和一个set集合
// set集合主要解决没有出现的字符串
// 然后dfs去找target，返回值设为0主要是怕出现负数的情况
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	set := make(map[string]bool)
	for i, equation := range equations {
		a := equation[0]
		b := equation[1]
		val := values[i]

		if _, ok := graph[a]; !ok {
			graph[a] = make(map[string]float64)
		}
		if _, ok := graph[b]; !ok {
			graph[b] = make(map[string]float64)
		}
		set[a] = true
		set[b] = true
		graph[a][b] = val
		graph[b][a] = 1 / val
	}
	res := []float64{}
	for _, querie := range queries {
		a := querie[0]
		b := querie[1]
		value := calculate(graph, make(map[string]bool), a, b)
		getValue := func() float64 {
			if value == 0 {
				return -1
			}
			if !contain(set, a) || !contain(set, b) {
				return -1
			}
			return value
		}
		res = append(res, getValue())
	}
	return res
}

func contain(set map[string]bool, key string) bool {
	for k := range set {
		if key == k {
			return true
		}
	}
	return false
}

func calculate(graph map[string]map[string]float64, visited map[string]bool, cur, target string) float64 {
	if cur == target {
		return 1.0
	}
	visited[cur] = true
	for key, val := range graph[cur] {
		if visited[key] {
			continue
		}
		v := calculate(graph, visited, key, target)
		if v != 0 {
			return val * v
		}
	}
	return 0
}
