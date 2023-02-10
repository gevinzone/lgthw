package internal

import (
	"context"
	"github.com/gevinzone/lgthw/loop1/c07Micro/grpcjson/keyvalue/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
)

type KeyValue struct {
	mutex sync.RWMutex
	m     map[string]string
	gen.UnimplementedKeyValueServer
}

func NewKeyValue() *KeyValue {
	return &KeyValue{
		m: make(map[string]string),
	}
}

func (k *KeyValue) Set(ctx context.Context, r *gen.SetKeyValueRequest) (*gen.KeyValueResponse, error) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	k.m[r.GetKey()] = r.GetValue()
	return &gen.KeyValueResponse{
		Success: "true",
		Value:   r.GetValue(),
	}, nil

	//return &gen.KeyValueResponse{
	//	Success: "true",
	//	Value:   r.GetValue(),
	//}, nil
}

func (k *KeyValue) Get(ctx context.Context, r *gen.GetKeyValueRequest) (*gen.KeyValueResponse, error) {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	if val, ok := k.m[r.GetKey()]; !ok {
		return nil, status.Errorf(codes.NotFound, "key not set")
	} else {
		return &gen.KeyValueResponse{Value: val}, nil
	}
	//return &gen.KeyValueResponse{Value: r.Key}, nil
}
