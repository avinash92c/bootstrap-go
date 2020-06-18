package foundation

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
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, options *CacheOptions) error
	Delete(key string) error
	Invalidate(options *CacheOptions) error
	Clear() error
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

func (cache *cachestore) Get(key string) (interface{}, error) {
	return cache.cachemap[key], nil
}

func (cache *cachestore) Set(key string, value interface{}, options *CacheOptions) error {
	cache.cachemap[key] = value
	return nil
}

func (cache *cachestore) Delete(key string) error {
	delete(cache.cachemap, key)
	return nil
}

func (cache *cachestore) Invalidate(options *CacheOptions) error {
	cache.cachemap = make(map[string]interface{})
	return nil
}

func (cache *cachestore) Clear() error {
	cache.cachemap = make(map[string]interface{})
	return nil
}
