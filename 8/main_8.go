package main

import (
	"go-core-task/waitgroup"
	"time"
)

func main() {
	var awg waitgroup.AtomicWaitGroup

	for i := 0; i < 3; i++ {
		awg.Add(1)
		go func() {
			defer awg.Done()
			time.Sleep(2 * time.Second)
		}()
	}

	awg.WaitTimeout(3 * time.Second)

	var wg waitgroup.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
		}()
	}

	wg.Wait()
}
