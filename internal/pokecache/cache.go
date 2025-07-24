package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	var newEntries map[string]cacheEntry
	var mutex sync.Mutex

	cache := Cache{
		entries: newEntries,
		mu:      &mutex,
	}

	return cache
}

func (c *Cache) reapLoop(interval time.Duration) {
	time.NewTicker(interval)
}

func (c *Cache) Add(key string, val []byte) {
	e := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.entries[key] = e
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}
