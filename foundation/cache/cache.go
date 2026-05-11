package cache

import "context"

//BELOW IS A FILLER CACHE IMPLEMENTATION, INTERFACE IS PRIORITY
//TODO ADD LRU,LFU OPTIONS,REMOVE PLACEHOLDER STATIC CACHEMAP
//TODO ENSURE CONCURRENCY SAFE

//STATIC CACHEMAP IN-MEMORY
var (
	staticcachemap = make(map[string]interface{})
	cache          *cachestore
)

// CacheStore interface
type CacheStore interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}, options *CacheOptions) error
	Delete(ctx context.Context, key string) error
	Invalidate(ctx context.Context, options *CacheOptions) error
	Clear(ctx context.Context) error
	// GetType() string
}

// CacheOptions contains options for configuring cache
type CacheOptions struct {
}

func initcache() {
	cache = &cachestore{
		cachemap: staticcachemap,
	}
}

type cachestore struct {
	cachemap map[string]interface{}
}

// GetCache function
func GetCache() CacheStore {
	return cache
}

func (cache *cachestore) Get(ctx context.Context, key string) (interface{}, error) {
	return cache.cachemap[key], nil
}

func (cache *cachestore) Set(ctx context.Context, key string, value interface{}, options *CacheOptions) error {
	cache.cachemap[key] = value
	return nil
}

func (cache *cachestore) Delete(ctx context.Context, key string) error {
	delete(cache.cachemap, key)
	return nil
}

func (cache *cachestore) Invalidate(ctx context.Context, options *CacheOptions) error {
	cache.cachemap = make(map[string]interface{})
	return nil
}

func (cache *cachestore) Clear(ctx context.Context) error {
	cache.cachemap = make(map[string]interface{})
	return nil
}
