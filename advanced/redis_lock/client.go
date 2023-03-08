package redis_lock

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"time"
)

type Client struct {
	client       redis.Cmdable
	valGenerator func() string
}

func (c *Client) Lock(ctx context.Context, key string, expiration, timeout time.Duration, retry RetryStrategy) (*Lock, error) {
	panic("implement me")
}

func (c *Client) TryLock(ctx context.Context, key string, expiration time.Duration) (*Lock, error) {
	panic("implement me")
}

func (c *Client) createVal() string {
	if c.valGenerator == nil {
		return uuid.New().String()
	}
	return c.valGenerator()
}
