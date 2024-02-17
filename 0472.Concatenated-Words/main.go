package main

import (
	"fmt"
	"reflect"
)

func main() {
	testCases := []struct {
		words []string
		want  []string
	}{
		{
			words: []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatsdogcat"},
			want:  []string{"catsdogcats", "dogcatsdog", "ratcatsdogcat"},
		},
		{
			words: []string{"cat", "dog", "catdog"},
			want:  []string{"catdog"},
		},
		{
			words: []string{"cat", "dog"},
			want:  []string{},
		},
	}

	for _, tc := range testCases {
		got := findAllConcatenatedWordsInADict(tc.words)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("findAllConcatenatedWordsInADict(%v) = %v; want %v。", tc.words, got, tc.want)
		}
	}
}

// 回溯 + dp
func findAllConcatenatedWordsInADict(words []string) []string {
	sets := make(map[string]struct{})
	for _, word := range words {
		sets[word] = struct{}{}
	}

	var result []string

	for key := range sets {
		if key == "" {
			continue
		}
		// 先把自己排除掉
		delete(sets, key)

		if canBreak(key, sets) {
			result = append(result, key)
		}

		sets[key] = struct{}{}
	}
	return result
}

// dp[i] s[0]到s[i-1]是否是个连接词
func canBreak(key string, sets map[string]struct{}) bool {
	size := len(key)
	dp := make([]bool, size+1)

	dp[0] = true
	for i := 0; i <= size; i++ {
		// 这里j就相当于是从中间切一刀，dp[j]代表前面是否是true。
		// 如果前面是true 后面也可以满足，那么dp[i]就是true了
		// 如果前面不满足，就continue，继续查找
		// 如果满足都满足，那么设置dp[i]=true，就可以直接break了
		for j := 0; j < size; j++ {
			if !dp[j] {
				continue
			}
			if contains(sets, key[j:i]) {
				dp[i] = true
				break
			}
		}
	}
	return dp[size]
}

func contains(sets map[string]struct{}, target string) bool {
	if _, ok := sets[target]; ok {
		return true
	}
	return false
}

// func findAllConcatenatedWordsInADict(words []string) []string {
// 	hmap := make(map[string]int)
// 	for _, word := range words {
// 		hmap[word] = 1
// 	}

// 	var result []string

// 	for _, word := range words {
// 		helper(word, word, &result, hmap)
// 	}

// 	return result
// }

// func helper(source string, target string, result *[]string, hmap map[string]int) {
// 	if len(target) == 0 {
// 		if !contain(result, source) {
// 			*result = append(*result, source)
// 		}
// 		return
// 	}

// 	for key, _ := range hmap {
// 		if key == source {
// 			continue
// 		}
// 		if len(key) > len(target) {
// 			continue
// 		}
// 		newTarget := strings.Replace(target, key, "", -1)
// 		if newTarget == source || newTarget == target {
// 			continue
// 		}
// 		helper(source, newTarget, result, hmap)
// 	}

// }

// func contain(result *[]string, str string) bool {
// 	for _, word := range *result {
// 		if word == str {
// 			return true
// 		}
// 	}
// 	return false
// }
