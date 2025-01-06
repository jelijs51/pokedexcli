package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createAt time.Time
	val []byte
}

type Cache struct {
	entries map[string]CacheEntry
	mutex sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = CacheEntry{
		createAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if entry, ok := c.entries[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		c.mutex.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mutex.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]CacheEntry),
	}
	go cache.ReapLoop(interval)
	return cache
}