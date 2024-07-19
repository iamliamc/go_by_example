package sync

import "sync"

type Counter struct {
	// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
