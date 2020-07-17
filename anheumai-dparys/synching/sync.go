package synching

import "sync"

type Counter struct {
	mux   sync.Mutex
	count int
}

func (c *Counter) Inc() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}
