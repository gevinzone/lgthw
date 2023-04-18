package v1

import (
	"context"
	"testing"
	"time"
)

func TestEtcdLeader_ElectLeader(t *testing.T) {
	endpoints := []string{"localhost:2379"}
	timeout := 5 * time.Second
	etcdLeader, err := NewEtcdLeader(endpoints, timeout)
	if err != nil {
		t.Fatalf("failed to create etcd leader: %v", err)
	}
	defer etcdLeader.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := etcdLeader.ElectLeader(ctx); err != nil {
		t.Fatalf("failed to elect leader: %v", err)
	}
}

func TestEtcdLeader_ResignLeadership(t *testing.T) {
	endpoints := []string{"localhost:2379"}
	timeout := 5 * time.Second
	etcdLeader, err := NewEtcdLeader(endpoints, timeout)
	if err != nil {
		t.Fatalf("failed to create etcd leader: %v", err)
	}
	defer etcdLeader.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := etcdLeader.ElectLeader(ctx); err != nil {
		t.Fatalf("failed to elect leader: %v", err)
	}

	if err := etcdLeader.ResignLeadership(ctx); err != nil {
		t.Fatalf("failed to resign leadership: %v", err)
	}
}

func TestEtcdLeader_GetLeader(t *testing.T) {
	endpoints := []string{"localhost:2379"}
	timeout := 5 * time.Second
	etcdLeader, err := NewEtcdLeader(endpoints, timeout)
	if err != nil {
		t.Fatalf("failed to create etcd leader: %v", err)
	}
	defer etcdLeader.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := etcdLeader.ElectLeader(ctx); err != nil {
		t.Fatalf("failed to elect leader: %v", err)
	}

	leader, err := etcdLeader.GetLeader(ctx)
	if err != nil {
		t.Fatalf("failed to get leader: %v", err)
	}

	if leader != "leader" {
		t.Fatalf("expected leader to be 'leader', but got '%s'", leader)
	}
}
