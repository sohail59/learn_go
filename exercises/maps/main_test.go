package maps

import (
	"reflect"
	"testing"
)

// table for testcases
var TestCases = []struct {
	str      string
	expected map[string]int
}{
	{
		str:      "My Favorite programming language is Go",
		expected: map[string]int{"My": 1, "Favorite": 1, "programming": 1, "language": 1, "is": 1, "Go": 1},
	},
	{
		str:      "tenet is read tenet backward",
		expected: map[string]int{"tenet": 2, "is": 1, "read": 1, "backward": 1},
	},
	{
		str:      "cat is lion unless said cat",
		expected: map[string]int{"cat": 2, "unless": 1, "is": 1, "lion": 1},
	},
}

func TestWordCount(t *testing.T) {

	for _, tt := range TestCases {

		t.Run("Test values", func(t *testing.T) {
			actual := WordCount(tt.str)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("actual %v, expected %v", actual, tt.expected)
			}
		})

	}
}
