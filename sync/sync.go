package sync

import "sync"

// Mutex allows us to add locks to our data
// WaitGroup is a means of waiting for goroutines to finish jobs
// Use channels when passing ownership of data
// Use mutexes for managing state

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

func NewCounter() *Counter {
	return &Counter{}
}
