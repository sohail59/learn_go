package maps

import (
	"strings"
)

// WordCount return the count of each word present in given 
// string separated by one or more space
func WordCount(s string) map[string]int {
	words := strings.Fields(s) //default delimiter is 'space'
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
