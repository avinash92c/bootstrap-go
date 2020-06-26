package cache

import (
	"time"

	"github.com/allegro/bigcache"
	ekocache "github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
)

type bigcachestore struct {
	cacheManager *ekocache.Cache
}

func newBigCache() CacheStore {
	return &bigcachestore{cacheManager: configurebigcache()}
}

func configurebigcache() *ekocache.Cache {
	bigcacheClient, _ := bigcache.NewBigCache(bigcache.DefaultConfig(5 * time.Minute))
	bigcacheStore := store.NewBigcache(bigcacheClient, nil) //TBD PASS OPTIONS
	cacheManager := ekocache.New(bigcacheStore)
	return cacheManager
}

func (bcache *bigcachestore) Get(key string) (interface{}, error) {
	return bcache.cacheManager.Get(key)
}

func (bcache *bigcachestore) Set(key string, value interface{}, options *CacheOptions) error {
	coptions := &store.Options{Cost: 2} //TEMP DUMMY //CHANGE POST DEFINING OPTIONS RANGE
	err := bcache.cacheManager.Set(key, value, coptions)
	return err
}

func (bcache *bigcachestore) Delete(key string) error {
	err := bcache.cacheManager.Delete(key)
	return err
}

func (bcache *bigcachestore) Invalidate(options *CacheOptions) error {
	coptions := store.InvalidateOptions{}
	return bcache.cacheManager.Invalidate(coptions)
}

func (bcache *bigcachestore) Clear() error {
	return bcache.cacheManager.Clear()
}
