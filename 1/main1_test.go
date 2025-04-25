package main

import (
	"reflect"
	"testing"
)

func TestCombiningValuesToStr(t *testing.T) {
	t.Run("Check Combining Values To String", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []any
			expected string
		}{
			{"The same types", []any{1, 2, 3, 4}, "1234"},
			{"Different types", []any{"Alex", 'M', 24, 2000.00}, "Alex77242000"},
			{"Types from the Task", []any{42, 052, 0x2A, 3.14, "Golang", true, 1 + 2i}, "4242423.14Golangtrue(1+2i)"},
		}

		for _, tt := range tests {
			result := combiningValuesToStr(tt.input...)
			if result != tt.expected {
				t.Errorf("Combining Values To String failed. Got %v, expected %v", result, tt.expected)
			}
		}
	})
}

func TestConvertToRunes(t *testing.T) {
	t.Run("Check Convert To Runes", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			expected []rune
		}{
			{"Test one", "AlexMayka", []rune("AlexMayka")},
			{"Test two", "4242423.14Golangtrue(1+2i)", []rune("4242423.14Golangtrue(1+2i)")},
			{"Test three", "", []rune("")},
		}

		for _, tt := range tests {
			result := convertToRunes(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Converting To Runes failed. Got %v, expected %v", result, tt.expected)
			}
		}
	})
}

func TestHashRunesSHA256(t *testing.T) {
	t.Run("Check Hash Runes SHA256", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []rune
			expected string
		}{
			{"Test one", []rune("AlexMayka"), "c4e116a5482d5e70e9293fc40ae2fe5ac12a4ce439690ce4a452a6dd0d4747b4"},
			{"Test two", []rune("4242423.14Golangtrue(1+2i)"), "53f2f60ac6c41389d3ed3d84d88d8c2860bf8981c677be18243a6f35a6b6a1b3"},
			{"Test three", []rune(""), "66802df107aace17871a5b610ff9eb11706e13477bb24e93966ca80671c0fac6"},
		}

		for _, tt := range tests {
			result := hashRunesSHA256(tt.input)
			if result != tt.expected {
				t.Errorf("Hash Runes SHA256 failed. Got %v, expected %v", result, tt.expected)
			}
		}
	})
}
