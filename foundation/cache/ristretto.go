package cache

import (
	"time"

	"github.com/dgraph-io/ristretto"
	ekocache "github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
)

type ristrettocache struct {
	cacheManager *ekocache.Cache
}

func newRistrettoCache() CacheStore {
	return &ristrettocache{cacheManager: configurerisrettocache()}
}

func configurerisrettocache() *ekocache.Cache {
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1000,
		MaxCost:     100,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}
	ristrettoStore := store.NewRistretto(ristrettoCache, nil)
	cacheManager := ekocache.New(ristrettoStore)
	return cacheManager
}

func (rcache *ristrettocache) Get(key string) (interface{}, error) {
	return rcache.cacheManager.Get(key)
}

func (rcache *ristrettocache) Set(key string, value interface{}, options *CacheOptions) error {
	coptions := &store.Options{Cost: 2} //TEMP DUMMY //CHANGE POST DEFINING OPTIONS RANGE
	err := rcache.cacheManager.Set(key, value, coptions)
	time.Sleep(time.Nanosecond * 100) //Pause to wait `ristretto.processItems` consume from channel
	return err
}

func (rcache *ristrettocache) Delete(key string) error {
	err := rcache.cacheManager.Delete(key)
	time.Sleep(time.Nanosecond * 1) //Pause to wait `ristretto.processItems` consume from channel
	return err
}

func (rcache *ristrettocache) Invalidate(options *CacheOptions) error {
	coptions := store.InvalidateOptions{}
	return rcache.cacheManager.Invalidate(coptions)
}

func (rcache *ristrettocache) Clear() error {
	return rcache.cacheManager.Clear()
}
