package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu *sync.Mutex
}

type cacheEntry struct {
		createdAt time.Time
		val       []byte
}

func NewCache(interval time.Duration) Cache{
	cache := Cache{
		cache: map[string]cacheEntry{},
		mu: &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c Cache) Add(key string, val []byte){
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[key]
	if !ok { return nil, ok }

	return entry.val, true
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().Add(-interval))
	}

}

func (c Cache) reap(cutoff time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.cache {
		if v.createdAt.Before(cutoff) {
			delete(c.cache, k)
		}
	}

}

