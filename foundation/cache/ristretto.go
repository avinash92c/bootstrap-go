package cache

//NEEDS FIXING. WORK IN PROGRESS

import (
	"context"
	"time"

	"github.com/dgraph-io/ristretto/v2"
	ekocache "github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/store"
)

type ristrettocache struct {
	cacheManager *ekocache.Cache[any]
}

func newRistrettoCache() CacheStore {
	return &ristrettocache{cacheManager: configurerisrettocache()}
}

func configurerisrettocache() *ekocache.Cache[any] {
	_, err := ristretto.NewCache(&ristretto.Config[string, any]{ //TBD INTERFACE NEEDS TO BE ADJUSTED HERE
		NumCounters: 1000,
		MaxCost:     100,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}
	// ristrettoStore := store.NewRistretto(ristrettoCache, nil) //TBD FIX THIS
	ristrettoStore := store.NewRistretto(nil, nil)
	cacheManager := ekocache.New[any](ristrettoStore)
	return cacheManager
}

func (rcache *ristrettocache) Get(ctx context.Context, key string) (interface{}, error) {
	return rcache.cacheManager.Get(ctx, key)
}

func (rcache *ristrettocache) Set(ctx context.Context, key string, value interface{}, options *CacheOptions) error {
	// coptions := &store.Options{Cost: 2} //TEMP DUMMY //CHANGE POST DEFINING OPTIONS RANGE
	// err := rcache.cacheManager.Set(ctx, key, value, coptions)
	err := rcache.cacheManager.Set(ctx, key, value, store.WithCost(2))
	time.Sleep(time.Nanosecond * 100) //Pause to wait `ristretto.processItems` consume from channel
	return err
}

func (rcache *ristrettocache) Delete(ctx context.Context, key string) error {
	err := rcache.cacheManager.Delete(ctx, key)
	time.Sleep(time.Nanosecond * 1) //Pause to wait `ristretto.processItems` consume from channel
	return err
}

func (rcache *ristrettocache) Invalidate(ctx context.Context, options *CacheOptions) error {
	// coptions := store.InvalidateOptions{}
	// return rcache.cacheManager.Invalidate(coptions)
	coptions := store.InvalidateOption(store.WithInvalidateTags([]string{}))
	return rcache.cacheManager.Invalidate(ctx, coptions)
}

func (rcache *ristrettocache) Clear(ctx context.Context) error {
	return rcache.cacheManager.Clear(ctx)
}
