package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	tests := []struct {
		tickets [][]string
		want    []string
	}{
		{
			tickets: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
			want:    []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			tickets: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			want:    []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
		{
			tickets: [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}},
			want:    []string{"JFK", "NRT", "JFK", "KUL"},
		},
	}

	for _, tt := range tests {
		got := findItineraryIteration(tt.tickets)
		if !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("findItinerary(%v) = %v, want %v", tt.tickets, got, tt.want)
		}
	}
}

// 模拟一下stack的情况
// 输入：["JFK", "SFO"], ["JFK", "ATL"], ["SFO", "ATL"], ["ATL", "JFK"], ["ATL", "SFO"]
// 我们期望的输出是 ["JFK", "ATL", "JFK", "SFO", "ATL", "SFO"]。
// 初始状态，栈 stack 为 ["JFK"]。
// 从 "JFK" 出发，我们有两个选择："SFO" 和 "ATL"。因为 "ATL" 在字典序上更小，所以我们选择 "ATL"。此时，栈 stack 为 ["JFK", "ATL"]。
// 从 "ATL" 出发，我们有两个选择："JFK" 和 "SFO"。因为 "JFK" 在字典序上更小，所以我们选择 "JFK"。此时，栈 stack 为 ["JFK", "ATL", "JFK"]。
// 从 "JFK" 出发，我们只有一个选择："SFO"。所以我们选择 "SFO"。此时，栈 stack 为 ["JFK", "ATL", "JFK", "SFO"]。
// 从 "SFO" 出发，我们只有一个选择："ATL"。所以我们选择 "ATL"。此时，栈 stack 为 ["JFK", "ATL", "JFK", "SFO", "ATL"]。
// 从 "ATL" 出发，我们只有一个选择："SFO"。所以我们选择 "SFO"。此时，栈 stack 为 ["JFK", "ATL", "JFK", "SFO", "ATL", "SFO"]。
// 所有的机票都被使用过了，我们开始从栈中取出机场并添加到行程中。因为我们是从行程的最后一个机场开始添加的，所以我们需要将机场添加到行程的开头。
// 最后，我们得到的行程就是 ["JFK", "ATL", "JFK", "SFO", "ATL", "SFO"]。
func findItineraryIteration(tickets [][]string) []string {
	m := make(map[string][]string)
	for _, ticket := range tickets {
		m[ticket[0]] = append(m[ticket[0]], ticket[1])
	}
	for key := range m {
		sort.Strings(m[key])
	}

	res := []string{}
	stack := []string{"JFK"}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		// 找一下机场列表里有没有数据
		if len(m[cur]) == 0 {
			res = append([]string{cur}, res...)
			stack = stack[:len(stack)-1]
		} else {
			next := m[cur][0]
			m[cur] = m[cur][1:]
			stack = append(stack, next)
		}
	}
	return res
}

// 构建一个有向图，使用dfs，从JFK进行搜索
// 需要对每个机场的目的地列表进行排序，每次选择完就删除他（所有的机票必须都用一次，且只能用一次）
// 其实我们只考虑了从JFK开始进行dfs，并没有考虑怎么回到JFK
// 因为题目说了假定所有机票至少存在一种合理的行程，所以最后肯定是可以回来的，我们不需要去考虑这种情况
func findItinerary(tickets [][]string) []string {
	m := make(map[string][]string)
	for _, ticket := range tickets {
		m[ticket[0]] = append(m[ticket[0]], ticket[1])
	}
	for key := range m {
		sort.Strings(m[key])
	}

	res := []string{}
	dfs(m, "JFK", &res)

	return res
}

func dfs(m map[string][]string, cur string, res *[]string) {
	for len(m[cur]) > 0 {
		next := m[cur][0]
		m[cur] = m[cur][1:]
		dfs(m, next, res)
	}
	// 这里要倒着添加
	// 因为dfs的特性，会搜索到最后一个机场
	*res = append([]string{cur}, *res...)
}
