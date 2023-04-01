package leader

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestCampaign(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()

	key := "/test/leader"
	id := "test-id"
	ttl := 5
	election := NewEtcdLeaderElection(cli, key, id, ttl)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = election.Campaign(ctx)
	if err != nil {
		t.Fatal(err)
	}

	isLeader, err := election.IsLeader(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !isLeader {
		t.Errorf("expected to be leader, but is not")
	}
}

func TestResign(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()

	key := "/test/leader"
	id := "test-id"
	ttl := 5
	election := NewEtcdLeaderElection(cli, key, id, ttl)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = election.Campaign(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = election.Resign(ctx)
	if err != nil {
		t.Fatal(err)
	}

	isLeader, err := election.IsLeader(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if isLeader {
		t.Errorf("expected not to be leader, but is")
	}
}

func TestIsLeader(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()

	key := "/test/leader"
	id := "test-id"
	ttl := 5
	election := NewEtcdLeaderElection(cli, key, id, ttl)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// check if initially not leader
	isLeader, err := election.IsLeader(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if isLeader {
		t.Errorf("expected not to be leader, but is")
	}

	// campaign and check if leader
	err = election.Campaign(ctx)
	if err != nil {
		t.Fatal(err)
	}

	isLeader, err = election.IsLeader(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !isLeader {
		t.Errorf("expected to be leader, but is not")
	}

	// resign and check if not leader
	err = election.Resign(ctx)
	if err != nil {
		t.Fatal(err)
	}

	isLeader, err = election.IsLeader(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if isLeader {
		t.Errorf("expected not to be leader, but is")
	}
}

// TestIsLeaderWithMultipleInstances tests leader election with multiple instances.
func TestIsLeaderWithMultipleInstances(t *testing.T) {
	// create etcd client
	cli1, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer cli1.Close()

	cli2, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer cli2.Close()

	// create etcdLeaderElection instances
	key := "/test/leader"
	id1 := "test-id-1"
	id2 := "test-id-2"
	ttl := 5
	election1 := NewEtcdLeaderElection(cli1, key, id1, ttl)
	election2 := NewEtcdLeaderElection(cli2, key, id2, ttl)

	// test campaign and isLeader
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = election1.Campaign(ctx)
	if err != nil {
		t.Fatal(err)
	}

	isLeader, err := election1.IsLeader(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !isLeader {
		t.Errorf("expected election1 to be leader, but is not")
	}

	isLeader, err = election2.IsLeader(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if isLeader {
		t.Errorf("expected election2 not to be leader, but is")
	}

	// test resign and isLeader
	err = election1.Resign(ctx)
	if err != nil {
		t.Fatal(err)
	}

	isLeader, err = election1.IsLeader(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if isLeader {
		t.Errorf("expected election1 not to be leader, but is")
	}

	isLeader, err = election2.IsLeader(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !isLeader {
		t.Errorf("expected election2 to be leader, but is not")
	}
}
