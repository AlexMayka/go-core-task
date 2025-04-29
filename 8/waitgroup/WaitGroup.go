package waitgroup

import (
	"sync"
)

type WaitGroup struct {
	count int
	mutex sync.Mutex
	sem   chan struct{}
}

func (wg *WaitGroup) Add(delta int) {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()

	if wg.count == 0 {
		wg.sem = make(chan struct{})
	}

	wg.count += delta
}

func (wg *WaitGroup) Done() {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()

	wg.count -= 1
	if wg.count == 0 {
		close(wg.sem)
	}
}

func (wg *WaitGroup) Wait() {
	<-wg.sem
}

func (wg *WaitGroup) Len() int {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()

	return wg.count
}
