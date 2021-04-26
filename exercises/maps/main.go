package maps

import (
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCountMap := make(map[string]int)

	//count the word in the map
	for _, word := range words {

		if _, ok := wordCountMap[word]; ok {
			wordCountMap[word] += 1
		} else {
			wordCountMap[word] = 1
		}
	}
	return wordCountMap
}
