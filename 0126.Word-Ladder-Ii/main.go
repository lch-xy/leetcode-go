package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		name           string
		beginWord      string
		endWord        string
		wordList       []string
		expectedResult [][]string
	}{
		{
			name:      "Example 1",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expectedResult: [][]string{
				{"hit", "hot", "dot", "dog", "cog"},
				{"hit", "hot", "lot", "log", "cog"},
			},
		},
		// {
		// 	name:           "Example 2",
		// 	beginWord:      "hit",
		// 	endWord:        "cog",
		// 	wordList:       []string{"hot", "dot", "dog", "lot", "log"},
		// 	expectedResult: [][]string{},
		// },
	}

	for _, tc := range testCases {
		result := findLadders(tc.beginWord, tc.endWord, tc.wordList)
		// if !reflect.DeepEqual(result, tc.expectedResult) {
		fmt.Printf("findLadders(%q, %q, %v) = %v, expected %v", tc.beginWord, tc.endWord, tc.wordList, result, tc.expectedResult)
		// }
	}
}

// step one : bfs to trace the shortest route
// step tow : dfs to find all route
// we need two collection :
// dist: store the distance from beginWord to each word encountered
// pres: a mapping from each word to its possible precursors in the sequence
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	allPath := make([][]string, 0)
	// to create a wordSet
	wordSet := make(map[string]struct{})
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}
	delete(wordSet, beginWord)
	// create dists and pres
	dists := make(map[string]int)
	pres := make(map[string]map[string]struct{})
	// create queue
	queue := []string{beginWord}
	isFindRoute := false
	step := 0
	for len(queue) > 0 && !isFindRoute {
		step++
		for j := len(queue); j > 0; j-- {
			curNode := queue[0]
			queue = queue[1:]
			for i := 0; i < len(curNode); i++ {
				for c := 'a'; c <= 'z'; c++ {
					newWord := replaceAtIndex(curNode, c, i)
					// when we encounter the newWord for the second time
					// wordSet don't contains newWord, but we still need to add it to pres collection
					// why value == step ? only words at the current level are considered.
					if value, ok := dists[newWord]; ok && value == step {
						pres[newWord][curNode] = struct{}{}
					}
					if !contains(wordSet, newWord) {
						continue
					}
					if _, ok := pres[newWord]; !ok {
						pres[newWord] = make(map[string]struct{})
					}
					pres[newWord][curNode] = struct{}{}
					// first loop we will to delete the newWord
					delete(wordSet, newWord)
					dists[newWord] = step
					queue = append(queue, newWord)
					if endWord == newWord {
						isFindRoute = true
					}
				}
			}
		}
	}

	dfs(&allPath, pres, []string{}, beginWord, endWord)

	return allPath
}

func dfs(allPath *[][]string, pres map[string]map[string]struct{}, curPath []string, beginWord, curWord string) {
	if beginWord == curWord {
		newPath := append([]string{curWord}, curPath...)
		*allPath = append(*allPath, newPath)
		return
	}

	for element := range pres[curWord] {
		dfs(allPath, pres, append([]string{curWord}, curPath...), beginWord, element)
	}
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	if i < 0 || i > len(in) {
		return in
	}
	out[i] = r
	return string(out)
}

func getAtIndex(in string, i int) string {
	out := []rune(in)
	if i < 0 || i >= len(out) {
		return in // out of index return original value
	}
	return string(out[i])
}

func contains(wordSet map[string]struct{}, word string) bool {
	if _, ok := wordSet[word]; ok {
		return true
	}
	return false
}
