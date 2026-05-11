package cache

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
	ekocache "github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/store"
)

type bigcachestore struct {
	cacheManager *ekocache.Cache[any]
}

func newBigCache() CacheStore {
	return &bigcachestore{cacheManager: configurebigcache()}
}

func configurebigcache() *ekocache.Cache[any] {
	// bigcacheClient, _ := bigcache.NewBigCache(bigcache.DefaultConfig(5 * time.Minute))
	bigcacheClient, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(5*time.Minute))
	bigcacheStore := store.NewBigcache(bigcacheClient) //TBD PASS OPTIONS
	// cacheManager := ekocache.New(bigcacheStore)
	cacheManager := ekocache.New[any](bigcacheStore)
	return cacheManager
}

func (bcache *bigcachestore) Get(ctx context.Context, key string) (interface{}, error) {
	return bcache.cacheManager.Get(ctx, key)
}

func (bcache *bigcachestore) Set(ctx context.Context, key string, value interface{}, options *CacheOptions) error {
	// coptions := &store.Options{} //TEMP DUMMY //CHANGE POST DEFINING OPTIONS RANGE
	err := bcache.cacheManager.Set(ctx, key, value, store.WithCost(2))
	return err
}

func (bcache *bigcachestore) Delete(ctx context.Context, key string) error {
	err := bcache.cacheManager.Delete(ctx, key)
	return err
}

func (bcache *bigcachestore) Invalidate(ctx context.Context, options *CacheOptions) error {
	// coptions := store.InvalidateOptions{}
	coptions := store.InvalidateOption(store.WithInvalidateTags([]string{}))
	return bcache.cacheManager.Invalidate(ctx, coptions)
}

func (bcache *bigcachestore) Clear(ctx context.Context) error {
	return bcache.cacheManager.Clear(ctx)
}

// func (bcache *bigcachestore) GetCodec() codec.CodecInterface {
// 	return bcache.cacheManager.GetCodec()
// }

func (bcache *bigcachestore) GetType() string {
	return bcache.cacheManager.GetType()
}
