package cache

import (
	"github.com/allegro/bigcache"
	"time"
)

var C Cache

type Cache struct{
	cache *bigcache.BigCache
}

func LoadCache()(c Cache,err error){
	var cacheConfig = bigcache.Config{
		// todo: get number of Shards from config
		Shards:               1028,
		// todo: get LifeWindow of from config
		LifeWindow:           30 * time.Minute,
		// todo: get CleanWindow of from config
		CleanWindow:          5 * time.Minute,
		// todo: get MaxEntriesInWindow of from config
		MaxEntriesInWindow:   1000 * 10 * 60,
		// todo: get MaxEntrySize from constant after we discover the size of redirect struct in memory
		MaxEntrySize:         32,
		// todo: get HardMaxCacheSize of from config
		HardMaxCacheSize:     512,
		// todo: get Logger of from logging
		Logger:               nil,
	}
	c.cache, err = bigcache.NewBigCache(cacheConfig)
	return
}


func (c Cache)Set(key string, value  []byte)(err error){
	return c.cache.Set(key,value)
}

func (c Cache)Get(key string)(value []byte, err error){
	return c.cache.Get(key)
}

func (c Cache)Delete(key string)(err error){
	return c.cache.Delete(key)
}
