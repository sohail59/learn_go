package maps

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {

	t.Run("test empty string input", func(t *testing.T) {
		input := ""
		actual := WordCount(input)
		expected := make(map[string]int, 0)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected an empty map: %v", actual)
		}
	})

	t.Run("test large space between words", func(t *testing.T) {
		input := `Hello						world`
		actual := WordCount(input)
		expected := map[string]int{
			"Hello": 1,
			"world": 1,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v does not match with actual: %v", expected, actual)
		}

	})

	t.Run("test case sensitiveness", func(t *testing.T) {
		input := "Tenet is read tenet"
		actual := WordCount(input)
		expected := map[string]int{
			"Tenet": 1,
			"is":    1,
			"read":  1,
			"tenet": 1,
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v does not match with actual: %v", expected, actual)
		}
	})
}
