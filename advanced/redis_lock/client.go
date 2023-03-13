package redis_lock

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

var (
	ErrFailedToPreemptLock = errors.New("redis-lock: 抢锁失败")
	//go:embed lua/lock.lua
	luaLock string
)

type Client struct {
	client       redis.Cmdable
	g            singleflight.Group
	valGenerator func() string
}

// SingleFlightLock 用SingleFlight防止缓存击穿
func (c *Client) SingleFlightLock(ctx context.Context, key string, expiration, timeout time.Duration, retry RetryStrategy) (*Lock, error) {
	panic("implement me")
}

func (c *Client) Lock(ctx context.Context, key string, expiration, timeout time.Duration, retry RetryStrategy) (*Lock, error) {
	val := c.createVal()
	var timer *time.Timer
	for {
		lCtx, cancel := context.WithTimeout(ctx, timeout)
		res, err := c.client.Eval(lCtx, luaLock, []string{key}, val, expiration.Seconds()).Result()
		cancel()
		// 超时以外的err，直接返回，否则可以重试
		if err != nil && !errors.Is(err, context.DeadlineExceeded) {
			return nil, err
		}
		if res == "OK" {
			return NewLock(c.client, key, val, expiration), nil
		}

		// 开始重试
		interval, ok := retry.Next()
		if !ok {
			return nil, fmt.Errorf("redis-lock: 超出重试限制, %w", ErrFailedToPreemptLock)
		}
		if timer == nil {
			timer = time.NewTimer(interval)
		} else {
			timer.Reset(interval)
		}
		select {
		case <-timer.C:
		case <-ctx.Done():
			return nil, ErrFailedToPreemptLock
		}
	}

}

func (c *Client) TryLock(ctx context.Context, key string, expiration time.Duration) (*Lock, error) {
	val := c.createVal()
	ok, err := c.client.SetNX(ctx, key, val, expiration).Result()
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrFailedToPreemptLock
	}
	return NewLock(c.client, key, val, expiration), nil
}

func (c *Client) createVal() string {
	if c.valGenerator == nil {
		return uuid.New().String()
	}
	return c.valGenerator()
}

type ClientV2 struct {
	Client
	mu sync.Mutex
}

// LockLock 用上锁方式防止缓存击穿
func (c *ClientV2) LockLock(ctx context.Context, key string, expiration, timeout time.Duration, retry RetryStrategy) (*Lock, error) {
	panic("implement me")
}
