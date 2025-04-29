package waitgroup

import (
	"sync"
	"sync/atomic"
	"time"
)

type AtomicWaitGroup struct {
	count    int32
	sem      chan struct{}
	doneOnce sync.Once
}

func (wg *AtomicWaitGroup) Add(delta int32) {
	if atomic.LoadInt32(&wg.count) == 0 && delta > 0 {
		wg.sem = make(chan struct{})
		wg.doneOnce = sync.Once{}
	}
	atomic.AddInt32(&wg.count, delta)
}

func (wg *AtomicWaitGroup) Done() {
	newCount := atomic.AddInt32(&wg.count, -1)
	if newCount == 0 {
		wg.doneOnce.Do(func() {
			close(wg.sem)
		})
	}
}

func (wg *AtomicWaitGroup) Wait() {
	<-wg.sem
}

func (wg *AtomicWaitGroup) WaitTimeout(timeout time.Duration) bool {
	select {
	case <-wg.sem:
		return true
	case <-time.After(timeout):
		wg.doneOnce.Do(func() {
			atomic.StoreInt32(&wg.count, 0)
			close(wg.sem)
		})
		return false
	}
}

func (wg *AtomicWaitGroup) Count() int32 {
	return atomic.LoadInt32(&wg.count)
}
