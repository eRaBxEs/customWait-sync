package waitsync

import (
	"sync"
	"time"
)

// WaitGroup is a custom implementation of a synchronisation primitive
type WaitGroup struct {
	counter int
	mutex   sync.Mutex
}

// NewWaitGroup initializes and returns a new WaitGroup
func NewWaitGroup() *WaitGroup {
	wg := &WaitGroup{}
	return wg
}

// Now to implement some of the needed methods by a Waitgroup
// Add increments the waitGroup counter by a specified delta
func (wg *WaitGroup) Add(delta int) {
	wg.mutex.Lock()
	wg.counter += delta
	wg.mutex.Unlock()
}

// Done decrements the WaitGroup counter by 1
func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

// Wait blocks till the WaitGroup counter is zero
func (wg *WaitGroup) Wait() {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()

	for wg.counter > 0 { // this happens for each goroutine
		wg.mutex.Unlock()
		time.Sleep(10 * time.Millisecond) // Busy wait with sleep
		wg.mutex.Lock()
	}
}
