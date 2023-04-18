package main

import (
	"context"
	"errors"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"time"
)

func main() {
	endpoints := []string{"localhost:2379"}
	key := "/my-election"
	//pid := os.Getegid()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	size := 10
	elections := make([]*concurrency.Election, size)
	for i := 0; i < 10; i++ {
		session, err := concurrency.NewSession(client)
		if err != nil {
			log.Fatal(err)
		}

		election := concurrency.NewElection(session, key)
		elections[i] = election
	}
	myNode := "myNode1"
	var leader *concurrency.Election
	for i := 0; i < size; i++ {
		go func(i int) {
			election := elections[i]
			err := election.Campaign(ctx, myNode)
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					fmt.Println("lose in campaign, ", i)
					return
				}
				fmt.Println(err)
				return
			}
			fmt.Println("I am the leader, ", i)
			leader = election
		}(i)
	}
	time.Sleep(time.Second * 1)
	_ = leader.Resign(context.Background())
}
