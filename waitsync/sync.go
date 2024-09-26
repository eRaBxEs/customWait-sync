package waitsync

import (
	"sync"
)

// WaitGroup is a custom implementation of a synchronisation primitive
type WaitGroup struct {
	counter int
	mutex   sync.Mutex
}
