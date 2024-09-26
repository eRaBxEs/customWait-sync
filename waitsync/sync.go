package waitsync

import (
	"sync"
)

// waitGroup is a custom implementation of a synchronisation primitive
type waitGroup struct {
	counter int
	mutex   sync.Mutex
}
