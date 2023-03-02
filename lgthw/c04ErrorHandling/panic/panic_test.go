package main

import (
	"fmt"
	"testing"
	"time"
)

// goroutinePanic 无论在哪个 Goroutine 中发生未被恢复的 panic，整个程序都将崩溃退出。
func goroutinePanic() {
	fmt.Println("start")
	go func() {
		//defer func() {
		//	recover()
		//}()
		fmt.Println("goroutine panic")
		panic("goroutine panic")
	}()
	time.Sleep(time.Second)
	fmt.Println("Complete")
}

func TestGoroutinePanic(t *testing.T) {
	goroutinePanic()
}
