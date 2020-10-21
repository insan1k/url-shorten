package cache

import (
	"github.com/allegro/bigcache"
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
)

var C Cache

type Cache struct {
	cache *bigcache.BigCache
}

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

func (c Cache) Set(key string, value []byte) (err error) {
	return c.cache.Set(key, value)
}

func (c Cache) Get(key string) (value []byte, err error) {
	return c.cache.Get(key)
}

func (c Cache) Delete(key string) (err error) {
	return c.cache.Delete(key)
}
