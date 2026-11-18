package proto

import "sync/atomic"

type Counter struct {
	counter int32
}

func (c *Counter) Increment() {
	atomic.AddInt32(&c.counter, 1)
}

func (c *Counter) Decrement() {
	atomic.AddInt32(&c.counter, -1)
}

func (c *Counter) Load() int32 {
	return atomic.LoadInt32(&c.counter)
}
