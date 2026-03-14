package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu *sync.Mutex
	data map[string]CacheEntry
}

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		mu: &sync.Mutex{},
		data: make(map[string]CacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.data[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.data {
			if time.Since(entry.createdAt) > interval {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}