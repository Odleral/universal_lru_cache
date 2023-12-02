package app

import (
	"sync"
	"time"
)

type Cache struct {
	capacity     int
	list         *List
	withTTL      bool
	withCapacity bool
	rwm          *sync.RWMutex
	hash         map[string]*Node
}

type Duration time.Duration
type Capacity int

func WithTTL(ttl time.Duration) Duration {
	return Duration(ttl)
}

func WithCapacity(capacity int) Capacity {
	return Capacity(capacity)
}

// NewLRUCache return LRU cache with TTL and max size
func NewLRUCache(params ...interface{}) *Cache {
	cache := Cache{hash: make(map[string]*Node), rwm: &sync.RWMutex{}}

	for _, param := range params {
		switch param.(type) {
		case Duration:
			cache.withTTL = true
			cache.list = NewList(time.Duration(param.(Duration)))
		case Capacity:
			cache.withCapacity = true
			cache.hash = make(map[string]*Node, param.(Capacity))
			cache.capacity = int(param.(Capacity))
		default:
			continue
		}
	}

	return &cache
}

func (c *Cache) Add(k string, v any) {
	c.rwm.RLock()
	if _, ok := c.hash[k]; ok {
		c.list.Remove(c.hash[k])
	}
	c.rwm.RUnlock()

	c.rwm.Lock()
	node := c.list.Append(k, v)
	c.hash[k] = node
	defer c.rwm.Unlock()

	if c.withCapacity && (c.list.len > c.capacity) {
		d := c.list.Pop()
		delete(c.hash, d.Key)
	}
}

func (c *Cache) Get(k string) (any, bool) {
	if _, ok := c.hash[k]; ok {
		node := c.hash[k]

		if c.withTTL && !node.TTL.After(time.Now()) {
			c.rwm.Lock()
			defer c.rwm.Unlock()

			c.list.Remove(c.hash[k])
			delete(c.hash, k)

			return nil, false
		}

		c.rwm.Lock()
		defer c.rwm.Unlock()

		c.list.Remove(node)
		c.hash[k] = c.list.Append(k, node.Value)

		return node.Value, true
	}

	return nil, false
}
