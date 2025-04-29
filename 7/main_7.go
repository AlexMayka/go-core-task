package main

import (
	"fmt"
	"sync"
)

func Merch(channels ...chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		var wg sync.WaitGroup
		wg.Add(len(channels))

		for _, ch := range channels {
			go func(c <-chan int) {
				defer wg.Done()
				for v := range c {
					out <- v
				}
			}(ch)
		}
		wg.Wait()
	}()

	return out
}

func main() {
	for i := range 5 {
		fmt.Println(i)
	}
}
