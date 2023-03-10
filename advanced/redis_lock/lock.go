package redis_lock

import (
	"context"
	_ "embed"
	"errors"
	"github.com/go-redis/redis/v9"
	"time"
)

var (
	ErrLockNotHold = errors.New("redis-lock: 你没有持有锁")

	//
	////go:embed lua/refresh.lua
	//luaRefresh string
	//
	////go:embed lua/lock.lua
	//luaLock string

	//go:embed lua/unlock.lua
	luaUnlock string
)

type Lock struct {
	client redis.Cmdable
	key    string
	value  string
}

func (l *Lock) AutoRefresh(interval time.Duration, timeout time.Duration) error {
	panic("implement me")
}

func (l *Lock) Refresh(ctx context.Context) error {
	panic("implement me")
}

func (l *Lock) Unlock(ctx context.Context) error {
	res, err := l.client.Eval(ctx, luaUnlock, []string{l.key}, l.value).Int64()
	if err != nil {
		return err
	}
	if res != 1 {
		return ErrLockNotHold
	}
	return nil
}

// Unlock 解锁
// 该代码解锁逻辑不是原子级的：检查自己是否持有锁，和删除自己的锁，分了2步
// 潜在问题
// 所以，解锁逻辑，要通过lua脚本改为原子级的，才能保证正确
//func (l *Lock) Unlock(ctx context.Context) error {
//	val, err := l.client.Get(ctx, l.key).Result()
//	if err != nil {
//		return err
//	}
//	if val != l.value {
//		return errors.New("not your lock")
//	}
//	// 这里如果有其他redis操作，把这个锁对应的redis 键值对(key, value)删掉了
//	// 然后又有另一个加锁请求，正好创建了key与此相同的一个锁
//	// 这时，自己的锁实际上已经没了，下面操作删掉的，其实别人的锁
//	cnt, err := l.client.Del(ctx, l.key).Result()
//	if err != nil {
//		return err
//	}
//	if cnt != 1 {
//		return ErrLockNotHold
//	}
//	return nil
//}
