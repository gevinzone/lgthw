package pattern

import (
	"context"
	"github.com/gevinzone/lgthw/advanced/cache"
	"sync"
	"time"
)

// WriteThroughCache 写缓存操作，会同时更新数据库
type WriteThroughCache struct {
	cache.Cache
	StoreFunc func(ctx context.Context, key string, val any) error
	mutex     sync.RWMutex
}

// Set 写缓存，为缓解缓存和数据库的数据不一致问题，通常套路为先写数据库，再写缓存
func (w *WriteThroughCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	err := w.StoreFunc(ctx, key, val)
	if err != nil {
		return err
	}
	return w.Cache.Set(ctx, key, val, expiration)
}
