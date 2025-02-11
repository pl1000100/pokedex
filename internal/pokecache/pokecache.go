package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

func NewCache(timeOut time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}
	go cache.reapLoop(timeOut)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	(*cache).entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	elem, ok := (*cache).entries[key]
	return elem.val, ok
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.clearCache(interval)
	}
}

func (cache *Cache) clearCache(interval time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	for key, entry := range cache.entries {
		if entry.createdAt.Add(interval).Compare(time.Now()) == -1 {
			// delete cache
			delete(cache.entries, key)
		}
	}
}
