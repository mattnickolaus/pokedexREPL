package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newEntries := make(map[string]cacheEntry)
	var mutex sync.Mutex

	cache := Cache{
		entries:  newEntries,
		mu:       &mutex,
		interval: interval,
	}
	cache.reapLoop()

	return cache
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	go func() {
		for {
			t := <-ticker.C // ticker returns the current time at tick
			c.mu.Lock()
			for key, val := range c.entries {
				if c.interval < t.Sub(val.createdAt) { // current tick time - createdAt
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}()

}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.entries[key] = e
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}
