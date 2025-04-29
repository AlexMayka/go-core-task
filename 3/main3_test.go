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

	removeTests := []struct {
		key string
	}{
		{"apple"},
		{"banana"},
		{"tomato"},
		{"avocado"},
		{"orange"},
	}

	for _, tt := range removeTests {
		testMap.Remove(tt.key)
	}

	checkTests := []struct {
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

	for _, tt := range checkTests {
		t.Run(fmt.Sprintf("Check key %s", tt.key), func(t *testing.T) {
			got := testMap.Exists(tt.key)
			if got != tt.wantExist {
				t.Errorf("Exists(%q) = %v; want %v", tt.key, got, tt.wantExist)
			}
		})
	}
}

func TestStringIntMap_Copy(t *testing.T) {
	t.Run("copying map", func(t *testing.T) {
		original := NewStringIntMap()
		original.Add("one", 1)
		original.Add("two", 2)
		original.Add("three", 3)

		copyMap := original.Copy()

		for key, expected := range map[string]int{"one": 1, "two": 2, "three": 3} {
			got, ok := copyMap[key]
			if !ok || got != expected {
				t.Errorf("Copy()[%q] = %d, want %d", key, got, expected)
			}
		}

		copyMap["four"] = 4
		if original.Exists("four") {
			t.Errorf("Original map unexpectedly contains key 'four'")
		}
	})
}

func TestStringIntMap_Exist(t *testing.T) {
	t.Run("checking key existence", func(t *testing.T) {
		m := NewStringIntMap()
		m.Add("alpha", 100)
		m.Add("beta", 200)

		tests := []struct {
			key      string
			expected bool
		}{
			{"alpha", true},
			{"beta", true},
			{"gamma", false},
		}

		for _, tt := range tests {
			t.Run(tt.key, func(t *testing.T) {
				got := m.Exists(tt.key)
				if got != tt.expected {
					t.Errorf("Exists(%q) = %v, want %v", tt.key, got, tt.expected)
				}
			})
		}
	})
}

func TestStringIntMap_Get(t *testing.T) {
	t.Run("getting values", func(t *testing.T) {
		m := NewStringIntMap()
		m.Add("dog", 5)
		m.Add("cat", 7)

		tests := []struct {
			key       string
			wantValue int
			wantOK    bool
		}{
			{"dog", 5, true},
			{"cat", 7, true},
			{"bird", 0, false},
		}

		for _, tt := range tests {
			t.Run(tt.key, func(t *testing.T) {
				got, ok := m.Get(tt.key)
				if ok != tt.wantOK {
					t.Errorf("Get(%q) ok = %v, want %v", tt.key, ok, tt.wantOK)
				}
				if got != tt.wantValue {
					t.Errorf("Get(%q) = %d, want %d", tt.key, got, tt.wantValue)
				}
			})
		}
	})
}
