package v1

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"time"
)

type EtcdLeader struct {
	client   *clientv3.Client
	session  *concurrency.Session
	election *concurrency.Election
}

func NewEtcdLeader(endpoints []string, timeout time.Duration) (*EtcdLeader, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: timeout,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create etcd client: %v", err)
	}

	session, err := concurrency.NewSession(client)
	if err != nil {
		return nil, fmt.Errorf("failed to create etcd session: %v", err)
	}

	election := concurrency.NewElection(session, "/election/leader")

	return &EtcdLeader{
		client:   client,
		session:  session,
		election: election,
	}, nil
}

func (e *EtcdLeader) ElectLeader(ctx context.Context) error {
	if err := e.election.Campaign(ctx, "leader"); err != nil {
		return fmt.Errorf("failed to campaign for leader: %v", err)
	}

	return nil
}

func (e *EtcdLeader) ResignLeadership(ctx context.Context) error {
	if err := e.election.Resign(ctx); err != nil {
		return fmt.Errorf("failed to resign leadership: %v", err)
	}

	return nil
}

func (e *EtcdLeader) GetLeader(ctx context.Context) (string, error) {
	resp, err := e.election.Leader(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get leader: %v", err)
	}

	return string(resp.Kvs[0].Value), nil
}

func (e *EtcdLeader) Close() error {
	e.session.Close()
	return e.client.Close()
}
