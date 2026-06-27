package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.Mutex
	interval time.Duration
	cache    map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		mu:       sync.Mutex{},
		interval: interval,
		cache:    make(map[string]cacheEntry),
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		val:       value,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.cache {
			if time.Since(v.createdAt) > c.interval {
				delete(c.cache, k)
			}
		}
		c.mu.Unlock()
	}
}
