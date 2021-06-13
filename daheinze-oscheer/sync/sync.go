package sync

import "sync"

// Counter defines the class
type Counter struct {
	value int
	mu    sync.Mutex
}

// Inc increses the value of counter by one
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns current value of the counter
func (c *Counter) Value() int {
	return c.value
}

// NewCounter initialize the object
func NewCounter() *Counter {
	return &Counter{}
}
