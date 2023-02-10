package cache

import (
	"context"
	"fmt"
	"github.com/gevinzone/lgthw/advanced/cache/mocks"
	"github.com/go-redis/redis/v9"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//go:generate mockgen -destination=./mocks/redis_mock.go -package=mocks github.com/go-redis/redis/v9 Cmdable

func TestRedisCache_Set(t *testing.T) {
	testCases := []struct {
		name       string
		key        string
		val        string
		expiration time.Duration
		mock       func(ctrl *gomock.Controller) redis.Cmdable
		wantErr    error
	}{
		{
			name:       "normal",
			key:        "key1",
			val:        "value1",
			expiration: time.Second,
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := mocks.NewMockCmdable(ctrl)
				//status := redis.NewStatusCmd(context.Background())
				//status.SetVal("OK")
				cmd.EXPECT().
					Set(gomock.Any(), "key1", "value1", time.Second).
					Return(redis.NewStatusResult("OK", nil))
				return cmd
			},
		},
		{
			name:       "timeout",
			key:        "key1",
			val:        "value1",
			expiration: time.Second,
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := mocks.NewMockCmdable(ctrl)
				cmd.EXPECT().
					Set(gomock.Any(), "key1", "value1", time.Second).
					Return(redis.NewStatusResult("", context.DeadlineExceeded))
				return cmd
			},
			wantErr: context.DeadlineExceeded,
		},
		{
			name:       "unexpected msg",
			key:        "key1",
			val:        "value1",
			expiration: time.Second,
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := mocks.NewMockCmdable(ctrl)
				cmd.EXPECT().
					Set(gomock.Any(), "key1", "value1", time.Second).
					Return(redis.NewStatusResult("something", nil))
				return cmd
			},
			wantErr: fmt.Errorf("%w, 返回信息 %s", errFailedToSetCache, "something"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := NewRedisCache(tc.mock(ctrl))
			err := client.Set(context.Background(), tc.key, tc.val, tc.expiration)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestRedisCache_Get(t *testing.T) {
	testCases := []struct {
		name    string
		key     string
		wantVal string
		wantErr error
		mock    func(ctrl *gomock.Controller) redis.Cmdable
	}{
		{
			name:    "normal",
			key:     "key1",
			wantVal: "value1",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := mocks.NewMockCmdable(ctrl)
				str := redis.NewStringCmd(context.Background())
				str.SetVal("value1")
				cmd.EXPECT().
					Get(gomock.Any(), "key1").
					Return(str)
				return cmd
			},
		},
		{
			name:    "timeout",
			key:     "key1",
			wantVal: "value1",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := mocks.NewMockCmdable(ctrl)
				str := redis.NewStringCmd(context.Background())
				str.SetErr(context.DeadlineExceeded)
				cmd.EXPECT().
					Get(gomock.Any(), "key1").
					Return(str)
				return cmd
			},
			wantErr: context.DeadlineExceeded,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cache := NewRedisCache(tc.mock(ctrl))
			value, err := cache.Get(context.Background(), tc.key)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVal, value)
		})
	}
}

func TestRedisCache_Delete(t *testing.T) {
}

func TestRedisCache_LoadAndDelete(t *testing.T) {

}
