package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

func Setup(addr string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: "",
		Password: "",
		DB:       0,
	})
	res, err := client.Ping(context.Background()).Result()
	fmt.Println(res, err)
	return client, err
}

func Exec(conn redis.Cmdable) error {
	conn.Set(context.Background(), "key", "value", 5*time.Second)
	res, err := conn.Get(context.Background(), "key").Result()
	if err != nil {
		return err
	}
	fmt.Println(res)
	var result string
	if err = conn.Get(context.Background(), "key").Scan(&result); err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func Sort(conn redis.Cmdable) error {
	ctx := context.Background()
	key := "list"
	if err := conn.LPush(ctx, key, 1).Err(); err != nil {
		return err
	}
	if err := conn.LPush(ctx, key, 3).Err(); err != nil {
		return err
	}
	if err := conn.LPush(ctx, key, 2).Err(); err != nil {
		return err
	}
	res, err := conn.Sort(ctx, key, &redis.Sort{Order: "ASC"}).Result()
	if err != nil {
		return err
	}
	fmt.Println(res)
	conn.Del(ctx, key)
	return nil
}
