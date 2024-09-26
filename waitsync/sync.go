package waitsync

import (
	"sync"
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
