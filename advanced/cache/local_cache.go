package cache

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	errKeyNotFound = errors.New("cache：键不存在")
	errKeyExpired  = errors.New("cache：键过期")
)

func WrapErrKeyNotFound(key string) error {
	return fmt.Errorf("%w, key is %s", errKeyNotFound, key)
}

func WrapErrKeyExpired(key string) error {
	return fmt.Errorf("%w, key is %s", errKeyExpired, key)
}

type BuildInCacheOption func(cache *BuildInCache)

// BuildInCache 基于内存的本地缓存
// 1. 并发安全
// 2. 过期删除机制：（1）延时删除，即Get()数据时，发现数据过期，则删除 （2）轮询定时删除一定数量的数据
type BuildInCache struct {
	mutex       sync.RWMutex
	data        map[string]*item
	onEvicted   func(key string, val any)
	closeSignal chan struct{}
}

func NewBuildInCache(interval time.Duration, opts ...BuildInCacheOption) *BuildInCache {
	cache := &BuildInCache{
		data: make(map[string]*item, 100),
		onEvicted: func(key string, val any) {

		},
		closeSignal: make(chan struct{}, 1),
	}

	for _, opt := range opts {
		opt(cache)
	}

	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-cache.closeSignal:
				return
			case t := <-ticker.C:
				count, max := 0, 1000
				cache.mutex.Lock()
				for key, itm := range cache.data {
					if count >= max {
						break
					}
					if !itm.isBeforeDeadline(t) {
						cache.delete(key)
					}
					count++
				}
				cache.mutex.Unlock()
			}
		}
	}()

	return cache
}

func (b *BuildInCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	var dl time.Time
	if expiration > 0 {
		dl = time.Now().Add(expiration)
	}
	b.data[key] = &item{
		val:      val,
		deadline: dl,
	}
	return nil
}

func (b *BuildInCache) Get(ctx context.Context, key string) (any, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	b.mutex.RLock()
	itm, ok := b.data[key]
	b.mutex.RUnlock()
	if !ok {
		return nil, WrapErrKeyNotFound(key)
	}
	if itm.isExpired() {
		b.mutex.Lock()
		defer b.mutex.Unlock()
		itm, ok := b.data[key]
		if !ok {
			return nil, WrapErrKeyNotFound(key)
		}
		if !itm.isExpired() {
			return itm, nil
		}
		b.delete(key)
		return nil, WrapErrKeyExpired(key)
	}
	return itm.val, nil
}

func (b *BuildInCache) Delete(ctx context.Context, key string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.delete(key)
	return nil
}

func (b *BuildInCache) delete(key string) {
	itm, ok := b.data[key]
	if !ok {
		return
	}
	delete(b.data, key)
	if b.onEvicted != nil {
		b.onEvicted(key, itm.val)
	}
}

func (b *BuildInCache) LoadAndDelete(ctx context.Context, key string) (any, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	itm, ok := b.data[key]
	if !ok {
		return nil, WrapErrKeyNotFound(key)
	}
	b.delete(key)
	return itm.val, nil
}

type item struct {
	val      any
	deadline time.Time
}

func (i *item) isExpired() bool {
	return i.isAfterDeadline(time.Now())
}

func (i *item) isBeforeDeadline(t time.Time) bool {
	return i.deadline.IsZero() || t.Before(i.deadline)
}

func (i *item) isAfterDeadline(t time.Time) bool {
	return !i.deadline.IsZero() && t.After(i.deadline)
}
