package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func main() {
	// 创建一个etcd客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 创建一个session，用于获取锁的租约
	sess, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	// 创建一个分布式锁，指定锁的键名为"my-lock"
	lock := concurrency.NewMutex(sess, "my-lock")

	// 尝试获取锁，如果获取成功，则执行业务逻辑，否则等待锁释放
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := lock.Lock(ctx); err != nil {
		log.Println("lock failed:", err)
		return
	}
	log.Println("lock acquired")

	// 模拟业务逻辑，打印当前时间
	fmt.Println("do something...", time.Now())

	// 释放锁，让其他客户端可以获取锁
	if err := lock.Unlock(ctx); err != nil {
		log.Println("unlock failed:", err)
		return
	}
	log.Println("lock released")
}
