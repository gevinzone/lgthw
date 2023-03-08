package redis_lock

import "time"

type RetryStrategy interface {
	// Next 返回重试的时间间隔，和要不要继续重试
	Next() (time.Duration, bool)
}

type FixedIntervalRetryStrategy struct {
	Interval time.Duration
	MaxCnt   int
	cnt      int
}

func (f *FixedIntervalRetryStrategy) Next() (time.Duration, bool) {
	if f.cnt >= f.MaxCnt {
		return 0, false
	}
	f.cnt++
	return f.Interval, true
}
