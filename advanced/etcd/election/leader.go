package election

import (
	"context"
	clientV3 "go.etcd.io/etcd/client/v3"
)

type LeaderElection interface {
	Campaign(ctx context.Context) error
	Resign(ctx context.Context) error
	IsLeader(ctx context.Context) (bool, error)
}

// etcdLeaderElection is a struct that implements the LeaderElection interface
// using etcd as the backend for leader election.
type etcdLeaderElection struct {
	client   *clientV3.Client
	key      string
	id       string
	ttl      int
	stopChan chan struct{}
}

// NewEtcdLeaderElection creates a new instance of etcdLeaderElection.
func NewEtcdLeaderElection(client *clientV3.Client, key, id string, ttl int) LeaderElection {
	return &etcdLeaderElection{
		client:   client,
		key:      key,
		id:       id,
		ttl:      ttl,
		stopChan: make(chan struct{}),
	}
}

// Campaign starts a new leader election campaign.
func (e *etcdLeaderElection) Campaign(ctx context.Context) error {
	lease, err := e.client.Grant(ctx, int64(e.ttl))
	if err != nil {
		return err
	}

	_, err = e.client.Put(ctx, e.key, e.id, clientV3.WithLease(lease.ID))
	if err != nil {
		return err
	}

	keepAliveChan, err := e.client.KeepAlive(ctx, lease.ID)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-e.stopChan:
				return
			case _, ok := <-keepAliveChan:
				if !ok {
					return
				}
			}
		}
	}()

	return nil
}

// Resign resigns from the current leader election campaign.
func (e *etcdLeaderElection) Resign(ctx context.Context) error {
	close(e.stopChan)

	_, err := e.client.Delete(ctx, e.key)
	if err != nil {
		return err
	}

	return nil
}

// IsLeader returns true if the current instance is the leader.
func (e *etcdLeaderElection) IsLeader(ctx context.Context) (bool, error) {
	resp, err := e.client.Get(ctx, e.key)
	if err != nil {
		return false, err
	}

	if len(resp.Kvs) == 0 {
		return false, nil
	}

	return string(resp.Kvs[0].Value) == e.id, nil
}
