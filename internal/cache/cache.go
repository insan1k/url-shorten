package cache

import (
	"github.com/allegro/bigcache"
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
)

// C is our cache singleton
var C Cache

// Cache is our cache struct holding a pointer to big cache
type Cache struct {
	cache *bigcache.BigCache
}

// LoadCache is our
func LoadCache() (err error) {
	var cacheConfig = bigcache.Config{
		Shards:             configuration.C.CacheShards,
		LifeWindow:         configuration.C.CacheLifeWindow,
		CleanWindow:        configuration.C.CacheCleanWindow,
		MaxEntriesInWindow: configuration.C.CacheMaxEntriesInWindow,
		MaxEntrySize:       configuration.C.CacheMaxEntrySize,
		HardMaxCacheSize:   configuration.C.CacheHardMaxCacheSize,
	}
	C.cache, err = bigcache.NewBigCache(cacheConfig)
	return
}

// Set some value in cache
func (c Cache) Set(key string, value []byte) (err error) {
	return c.cache.Set(key, value)
}

// Get some value in cache
func (c Cache) Get(key string) (value []byte, err error) {
	return c.cache.Get(key)
}

// Delete some value from cache
func (c Cache) Delete(key string) (err error) {
	return c.cache.Delete(key)
}
