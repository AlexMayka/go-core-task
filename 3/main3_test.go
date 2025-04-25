package main

import (
	"fmt"
	"testing"
)

func TestStringIntMap_Add(t *testing.T) {
	t.Run("adding multiple elements", func(t *testing.T) {
		m := NewStringIntMap()
		tests := []struct {
			name     string
			key      string
			value    int
			expected int
		}{
			{"add apple", "apple", 10, 10},
			{"add banana", "banana", 20, 20},
			{"add cherry", "cherry", 30, 30},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("Add %v", tt.name), func(t *testing.T) {
				m.Add(tt.key, tt.value)
				got, ok := m.Get(tt.key)
				if !ok {
					t.Errorf("Expected key %q to exist", tt.key)
				}
				if got != tt.expected {
					t.Errorf("For key %q: got %d, want %d", tt.key, got, tt.expected)
				}
			})
		}
	})
}

func TestStringIntMap_Remove(t *testing.T) {
	testMap := NewStringIntMap()
	testMap.Add("apple", 23)
	testMap.Add("banana", 44)
	testMap.Add("tomato", 22)
	testMap.Add("lime", 10)

	tests := []struct {
		key       string
		wantExist bool
	}{
		{"apple", false},
		{"banana", false},
		{"tomato", false},
		{"avocado", false},
		{"orange", false},
		{"lime", true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Remove %s", tt.key), func(t *testing.T) {
			testMap.Remove(tt.key)
			got := testMap.Exist(tt.key)
			if got != tt.wantExist {
				t.Errorf("Exist(%q) = %v; want %v", tt.key, got, tt.wantExist)
			}
		})
	}
}

func TestStringIntMap_Copy(t *testing.T) {
	testMap := NewStringIntMap()

	type TestCase struct {
		key       string
		value     int
	}

	tests := []struct {
		name string
		input []TestCase
	}
}
