package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type cache struct {
	cache map[string]cacheEntry
	m     *sync.RWMutex
}

func (c *cache) Add(k string, v []byte) {
	c.m.Lock()
	c.cache[k] = cacheEntry{
		createdAt: time.Now(),
		val:       v,
	}
	c.m.Unlock()
}

func (c *cache) Get(k string) ([]byte, bool) {
	c.m.RLock()
	ce, ok := c.cache[k]
	c.m.RUnlock()
	return ce.val, ok
}

func (c *cache) reapLoop(ttl time.Duration) {
	ticker := time.NewTicker(ttl)
	for tick := range ticker.C {
		c.m.Lock()
		for k, v := range c.cache {
			if v.createdAt.Before(tick.Add(-ttl)) {
				delete(c.cache, k)
			}
		}
		c.m.Unlock()
	}
}

func NewCache(ttl time.Duration) cache {
	c := cache{
		cache: map[string]cacheEntry{},
		m:     &sync.RWMutex{},
	}
	go c.reapLoop(ttl)
	return c
}
