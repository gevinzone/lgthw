package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

var errFailedToSetCache = errors.New("cache: 写入 redis 失败")

type RedisCache struct {
	client redis.Cmdable
}

func NewRedisCache(c redis.Cmdable) *RedisCache {
	return &RedisCache{client: c}
}

func (r *RedisCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	res, err := r.client.Set(ctx, key, val, expiration).Result()
	if err != nil {
		return err
	}
	if res != "OK" {
		return fmt.Errorf("%w, 返回信息 %s", errFailedToSetCache, res)
	}
	return nil
}

func (r *RedisCache) Get(ctx context.Context, key string) (any, error) {
	res := r.client.Get(ctx, key)
	return res.Result()
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	res := r.client.Del(ctx, key)
	return res.Err()
}

func (r *RedisCache) LoadAndDelete(ctx context.Context, key string) (any, error) {
	res := r.client.GetDel(ctx, key)
	return res.Result()
}
