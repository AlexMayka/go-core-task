package main

import (
	"reflect"
	"testing"
)

func TestPipeline(t *testing.T) {
	t.Run("Test Pipeline", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []uint8
			expected []float64
		}{
			{
				name:     "Multiple values",
				input:    []uint8{1, 2, 3, 4, 5},
				expected: []float64{1, 8, 27, 64, 125},
			},
			{
				name:     "Empty input",
				input:    []uint8{},
				expected: nil,
			},
			{
				name:     "Single value",
				input:    []uint8{10},
				expected: []float64{1000},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				in := make(chan uint8)
				out := make(chan float64)

				go Pipeline(in, out)

				go func() {
					for _, v := range tt.input {
						in <- v
					}
					close(in)
				}()

				var result []float64
				for v := range out {
					result = append(result, v)
				}

				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("Expected %v, got %v", tt.expected, result)
				}
			})
		}
	})
}
