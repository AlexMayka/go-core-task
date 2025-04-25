package main

import (
	"errors"
	"reflect"
	"slices"
	"testing"
)

func TestGetRandSlice(t *testing.T) {
	t.Run("Test Rand Slice", func(t *testing.T) {
		var lastSlice []int
		for i := 0; i < 10; i++ {
			testSlice := getRandSlice()
			if len(testSlice) != 10 {
				t.Errorf("getRandSlice() returned wrong number of slices")
			}

			if reflect.TypeOf(testSlice) != reflect.TypeOf([]int{}) {
				t.Errorf("getRandSlice() returned wrong type")
			}

			if reflect.DeepEqual(testSlice, lastSlice) {
				t.Errorf("getRandSlice() returned same slice")
			}

			lastSlice = testSlice
		}
	})
}

func TestSliceExample(t *testing.T) {
	t.Run("Test Filtering Even Numbers", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []int
			expected []int
		}{
			{"Only even elements", []int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}},
			{"Only odd elements", []int{1, 3, 5, 7, 9}, []int{}},
			{"All elements", []int{1, 2, 3, 4, 5, 6, 7, 8, 2}, []int{2, 4, 6, 8, 2}},
			{"Nil elements", []int{}, []int{}},
		}

		for _, tt := range tests {
			result := sliceExample(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sliceExample(%v) returned wrong result. Must: %v", tt.input, tt.expected)
			}

			if len(result) > 0 {
				maxElem := slices.Max(result)
				result[0] = maxElem + 1
				if reflect.DeepEqual(result, tt.expected) {
					t.Errorf("sliceExample(%v) returned same slice. Must: %v", tt.input, tt.expected)
				}
			}
		}
	})
}

func TestAddElement(t *testing.T) {
	t.Run("Test Adding Element", func(t *testing.T) {

		type inputParam struct {
			slice []int
			num   int
		}

		tests := []struct {
			name     string
			input    inputParam
			expected []int
		}{
			{"Add elem", inputParam{[]int{1, 2, 3, 4, 5}, 2}, []int{1, 2, 3, 4, 5, 2}},
			{"Add to empty slice", inputParam{[]int{}, 2}, []int{2}},
			{"Add to nil elice", inputParam{nil, 2}, []int{2}},
		}

		for _, tt := range tests {
			result := addElements(tt.input.slice, tt.input.num)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("addElement(%v) returned wrong result. Must: %v", tt.input, tt.expected)
			}

			maxElem := slices.Max(result)
			result[0] = maxElem + 1
			if reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sliceExample(%v) returned same slice. Must: %v", tt.input, tt.expected)
			}

		}
	})
}

func TestCopySlice(t *testing.T) {
	t.Run("Test Copying Slice", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []int
			expected []int
		}{
			{"Task 1", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			{"Task 2", []int{3, 1, 2, 4, 5}, []int{3, 1, 2, 4, 5}},
			{"Task 3", []int{}, []int{}},
			{"Task 4", nil, nil},
		}

		for _, tt := range tests {
			result := copySlice(tt.input)

			if len(result) > 0 && len(tt.expected) > 0 {
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("copySlice(%v) returned wrong result. Must: %v", tt.input, tt.expected)
				}

				maxElem := slices.Max(result)
				result[0] = maxElem + 1
				if reflect.DeepEqual(result, tt.expected) {
					t.Errorf("copySlice(%v) returned same slice. Must: %v", tt.input, tt.expected)
				}

				continue
			}

			if len(result) == 0 && len(tt.expected) == 0 {
				result = append(result, 2)
				if reflect.DeepEqual(result, tt.expected) {
					t.Errorf("copySlice(%v) returned same slice. Must: %v", tt.input, tt.expected)
				}

				continue
			}

			t.Errorf("copySlice(%v) returned wrong number of results. Must: %v", tt.input, tt.expected)
		}
	})
}

func TestRemoveElement(t *testing.T) {
	t.Run("Test Removing Element", func(t *testing.T) {
		type testInput struct {
			slice []int
			index int
		}

		tests := []struct {
			name     string
			input    testInput
			expected []int
			err      error
		}{
			{"Remove element first", testInput{[]int{1, 2, 3, 4, 5}, 0}, []int{2, 3, 4, 5}, nil},
			{"Remove element second", testInput{[]int{3, 1, 2, 4, 5}, 1}, []int{3, 2, 4, 5}, nil},
			{"Remove element second", testInput{[]int{3, 1, 2, 4, 5}, 4}, []int{3, 1, 2, 4}, nil},
			{"Error of range index", testInput{[]int{1, 2, 3, 4, 5}, 10}, nil, errors.New("index out of range")},
			{"Error of range index", testInput{[]int{1, 2, 3, 4, 5}, -1}, nil, errors.New("index out of range")},
			{"Error of range index", testInput{nil, 0}, nil, errors.New("index out of range")},
		}

		for _, tt := range tests {
			result, err := removeSlice(tt.input.slice, tt.input.index)
			if err == nil {
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("removeElement(%v) returned wrong result. Must: %v", tt.input, tt.expected)
				}

				maxElem := slices.Max(result)
				result[0] = maxElem + 1
				if reflect.DeepEqual(result, tt.expected) {
					t.Errorf("removeElement(%v) returned same slice. Must: %v", tt.input, tt.expected)
				}

				continue
			}

			if err.Error() == tt.err.Error() {
				continue
			}

			t.Errorf("removeElement(%v) returned wrong result. Must: %v", tt.input, tt.err)
		}
	})
}
