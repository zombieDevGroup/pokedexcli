package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache  map[string]cacheEntry
	mutex  sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache creates a new Cache that will cleanup entries older than the given interval
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache:    make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

// Add adds a value to the cache with the given key
func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get retrieves a value from the cache for the given key
// Returns the value and true if found, nil and false if not found
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

// reapLoop continuously cleans up old entries based on the interval
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.reap()
	}
}

// reap removes entries that are older than the interval
func (c *Cache) reap() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	now := time.Now()
	for k, v := range c.cache {
		if now.Sub(v.createdAt) > c.interval {
			delete(c.cache, k)
		}
	}
}
