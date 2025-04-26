package main

import (
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {
	t.Run("Test Generator", func(t *testing.T) {
		tests := []int{10, 20, 30, 1, 2, 0, -1}

		for _, test := range tests {
			t.Run(fmt.Sprintf("%d", test), func(t *testing.T) {
				ch := make(chan int)

				if test <= 0 {
					go Generator(ch, test)
					if _, ok := <-ch; ok {
						t.Errorf("Generator was not supposed to send values for count=%d", test)
					}
					return
				}

				result := make([]int, 0, test)
				go Generator(ch, test)

				for num := range ch {
					result = append(result, num)
				}

				if len(result) != test {
					t.Errorf("Got %d results, expected %d", len(result), test)
				}

				for _, num := range result {
					if num < 0 || num >= 10 {
						t.Errorf("Number out of range: got %d", num)
					}
				}
			})
		}
	})
}
