package creational

import (
	"sync"
)

// Pool represents the pool of objects to use.
type Pool struct {
	sync.Mutex
	inuse     []interface{}
	available []interface{}
	new       func() interface{}
}

// NewPool creates a new pool.
func NewPool(new func() interface{}) *Pool { _ = "STUB: not implemented"; return nil }

// Acquire acquires a new PoolObject to use from the pool.
// Here acquire creates a new instance of a PoolObject if none available.
func (p *Pool) Acquire() interface{} { _ = "STUB: not implemented"; return nil }

// Release releases a PoolObject back to the Pool.
func (p *Pool) Release(object interface{}) { _ = "STUB: not implemented"; return }
