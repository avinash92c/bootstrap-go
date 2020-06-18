package foundation

//TODO ADD LRU,LFU OPTIONS,REMOVE PLACEHOLDER STATIC CACHEMAP
//TODO ENSURE CONCURRENCY SAFE

//STATIC CACHEMAP IN-MEMORY
var (
	staticcachemap = make(map[string]interface{})
	cache          *cachestore
)

type cachestore struct {
	cachemap map[string]interface{}
}

// GetCache function
func GetCache() CacheStore {
	return cache
}

func initcache() {
	cache = &cachestore{
		cachemap: staticcachemap,
	}
}

// CacheStore interface
type CacheStore interface {
	Get(key string) interface{}
	Set(key, value string)
}

func (cache *cachestore) Get(key string) interface{} {
	return cache.cachemap[key]
}

func (cache *cachestore) Set(key, value string) {
	cache.cachemap[key] = value
}
