package redis_lock

import (
	"context"
	"time"
)

type Lock struct {
}

func (l *Lock) AutoRefresh(interval time.Duration, timeout time.Duration) error {
	panic("implement me")
}

func (l *Lock) Refresh(ctx context.Context) error {
	panic("implement me")
}

func (l *Lock) Unlock(ctx context.Context) error {
	panic("implement me")
}
