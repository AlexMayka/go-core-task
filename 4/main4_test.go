package main

import (
	"reflect"
	"testing"
)

func TestSliceDifference(t *testing.T) {
	t.Run("Test Slice Difference", func(t *testing.T) {
		type InputData struct {
			first  []string
			second []string
		}

		tests := []struct {
			name     string
			input    InputData
			expected []string
		}{
			{"Test 1", InputData{[]string{"a", "b", "c", "d", "f"}, []string{"a", "b", "c"}}, []string{"d", "f"}},
			{"Test 2", InputData{[]string{"c", "b", "f"}, []string{"a", "b", "e"}}, []string{"c", "f"}},
			{"Test 3", InputData{[]string{"a", "b", "c"}, []string{"d", "f", "g"}}, []string{"a", "b", "c"}},
			{"Test 4", InputData{[]string{"a", "b", "c"}, []string{"a", "b", "c"}}, []string{}},
			{"First empty", InputData{[]string{}, []string{"a", "b", "c"}}, []string{}},
			{"Second empty", InputData{[]string{"a", "b", "c"}, []string{}}, []string{"a", "b", "c"}},
			{"Both empty", InputData{[]string{}, []string{}}, []string{}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := SliceDifference(tt.input.first, tt.input.second)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("got %v, expected %v", result, tt.expected)
				}
			})
		}
	})
}
