package cache

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// ReadThroughCache 获取数据时，如果缓存里没有，会自动去数据库捞数据，然后更新缓存
type ReadThroughCache struct {
	Cache
	mutex      sync.RWMutex
	LoadFunc   func(ctx context.Context, key string) (any, error)
	Expiration time.Duration
}

func (r *ReadThroughCache) Get(ctx context.Context, key string) (any, error) {
	r.mutex.RLock()
	res, err := r.Cache.Get(ctx, key)
	r.mutex.RUnlock()
	if err == nil {
		return res, nil
	}
	if !errors.Is(err, errKeyNotFound) {
		return nil, err
	}
	r.mutex.Lock()
	defer r.mutex.Unlock()
	res, err = r.Cache.Get(ctx, key)
	if err == nil {
		return res, nil
	}
	val, err := r.LoadFunc(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("cache: 无法加载数据, %w", err)
	}
	err = r.Set(ctx, key, val, r.Expiration)
	return val, err
}
