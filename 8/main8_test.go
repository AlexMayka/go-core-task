package main

import (
	"go-core-task/waitgroup"
	"testing"
	"time"
)

func TestAtomicWaitGroup(t *testing.T) {
	t.Run("Wait without timeout", func(t *testing.T) {
		var awg waitgroup.AtomicWaitGroup

		for i := 0; i < 5; i++ {
			awg.Add(1)
			go func() {
				defer awg.Done()
				time.Sleep(50 * time.Millisecond)
			}()
		}

		awg.Wait()

		if count := awg.Count(); count != 0 {
			t.Errorf("Expected count 0 after Wait, got %d", count)
		}
	})

	t.Run("WaitTimeout succeeds before timeout", func(t *testing.T) {
		var awg waitgroup.AtomicWaitGroup

		for i := 0; i < 2; i++ {
			awg.Add(1)
			go func() {
				defer awg.Done()
				time.Sleep(50 * time.Millisecond)
			}()
		}

		ok := awg.WaitTimeout(200 * time.Millisecond)
		if !ok {
			t.Error("Expected WaitTimeout to succeed, but timeout occurred")
		}
	})

	t.Run("WaitTimeout fails due to timeout", func(t *testing.T) {
		var awg waitgroup.AtomicWaitGroup

		awg.Add(1)
		go func() {
			defer awg.Done()
			time.Sleep(2 * time.Second)
		}()

		ok := awg.WaitTimeout(100 * time.Millisecond)
		if ok {
			t.Error("Expected WaitTimeout to fail due to timeout, but it succeeded")
		}
	})
}

func TestWaitGroup(t *testing.T) {
	t.Run("Wait without timeout", func(t *testing.T) {
		var wg waitgroup.WaitGroup

		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				time.Sleep(50 * time.Millisecond)
			}()
		}

		wg.Wait()

		if length := wg.Len(); length != 0 {
			t.Errorf("Expected count 0 after Wait, got %d", length)
		}
	})
}
