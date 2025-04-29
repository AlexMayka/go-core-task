package main

import (
	"reflect"
	"testing"
)

func TestSliceUnification(t *testing.T) {
	t.Run("Test Slice Unification", func(t *testing.T) {
		type InputData struct {
			first  []int
			second []int
		}

		type OutputData struct {
			first  bool
			second []int
		}

		tests := []struct {
			name     string
			input    InputData
			expected OutputData
		}{
			{"Test 1", InputData{[]int{1, 2, 3, 4, 5}, []int{2, 3}}, OutputData{true, []int{2, 3}}},
			{"Test 2", InputData{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 6}}, OutputData{true, []int{2, 3, 4}}},
			{"Test 3", InputData{[]int{5}, []int{2, 3, 4, 6, 5}}, OutputData{true, []int{5}}},
			{"Test 4", InputData{[]int{5}, []int{6, 7, 8}}, OutputData{false, []int{}}},
			{"First empty", InputData{[]int{}, []int{1, 2, 3, 4}}, OutputData{false, []int{}}},
			{"Second empty", InputData{[]int{1, 2, 3, 4}, []int{}}, OutputData{false, []int{}}},
			{"Both empty", InputData{[]int{}, []int{}}, OutputData{false, []int{}}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ok, result := SliceUnification(tt.input.first, tt.input.second)
				if ok != tt.expected.first || !reflect.DeepEqual(result, tt.expected.second) {
					t.Errorf("got %v %v, expected %v", ok, result, tt.expected)
				}
			})
		}
	})
}
