package cache

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

var (
	errOutOfCapacity = errors.New("cache：超过容量限制")
)

func WrapErrOverCapacity() error {
	return fmt.Errorf("error: %w", errOutOfCapacity)
}

type MaxCntCache struct {
	*BuildInCache
	maxCnt int64
	cnt    int64
}

func NewMaxCntCache(c *BuildInCache, maxCnt int64) *MaxCntCache {
	res := &MaxCntCache{
		BuildInCache: c,
		maxCnt:       maxCnt,
	}

	origin := c.onEvicted
	res.onEvicted = func(key string, val any) {
		atomic.AddInt64(&res.cnt, -1)
		if origin != nil {
			origin(key, val)
		}
	}

	return res
}

func (m *MaxCntCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, ok := m.data[key]
	if !ok {
		if m.cnt+1 > m.maxCnt {
			return WrapErrOverCapacity()
		}
		m.cnt++
	}

	return m.set(key, val, expiration)
}


