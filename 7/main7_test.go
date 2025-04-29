package main

import (
	"fmt"
	"maps"
	"testing"
)

func TestMerge(t *testing.T) {
	t.Run("Test Merge Channels", func(t *testing.T) {
		tests := []struct {
			name          string
			countChannels int
			countBuffer   int
		}{
			{"Test 1", 3, 1},
			{"Test 2", 5, 2},
			{"Test 3", 0, 3},
			{"Test 4", 1, 5},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("%v. Count channel %v", tt.name, tt.countChannels), func(t *testing.T) {
				channels := make([]chan int, tt.countChannels)

				check := make(map[int]int, tt.countBuffer)
				for i := 0; i < tt.countBuffer; i++ {
					check[i] = 0
				}

				for index := range channels {
					channels[index] = make(chan int, tt.countBuffer)
					for val := range tt.countBuffer {
						channels[index] <- val
						check[val] += 1
					}
					close(channels[index])
				}

				result := Merch(channels...)
				collect := make(map[int]int, tt.countBuffer*tt.countChannels)

				for value := range result {
					if _, ok := check[value]; !ok {
						collect[value] = 1
						continue
					}

					collect[value] += 1
				}

				if tt.countChannels > 0 && !maps.Equal(check, collect) {
					t.Errorf("Expected %v, but got %v", check, collect)
				}
			})
		}
	})
}
