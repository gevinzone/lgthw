package ch8

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func prepareData(size int) <-chan int {
	ch := make(chan int, size)
	for i := 0; i < size; i++ {
		ch <- i + 1
	}
	return ch
}

func TestGetDataNoWait(t *testing.T) {
	size := 100
	dataCh := prepareData(size)
	for i := 0; i < size; i++ {
		go mockBiz(dataCh)
	}
	fmt.Println("Complete")
}

func TestGetDataWaitWithChannel(t *testing.T) {
	size := 100
	dataCh := prepareData(size)
	signal := make(chan struct{})
	for i := 0; i < size; i++ {
		go func() {
			mockBiz(dataCh)
			signal <- struct{}{}
		}()
	}
	for i := 0; i < size; i++ {
		<-signal
	}
	fmt.Println("Complete")
}

func TestGetDataWaitWithTimeout(t *testing.T) {
	size := 100
	dataCh := prepareData(size)
	signal := make(chan struct{}, size)
	for i := 0; i < size; i++ {
		go func() {
			mockBiz(dataCh)
			signal <- struct{}{}
		}()
	}
Loop:
	for {
		select {
		case <-signal:
		case <-time.After(time.Second * 2):
			close(signal)
			break Loop
		}
	}

	fmt.Println("Complete")
}

func TestGetDataWaitWithContext(t *testing.T) {
	size := 100
	dataCh := prepareData(size)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	for i := 0; i < size; i++ {
		go func() {
			mockBiz(dataCh)
		}()
	}
Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		}
	}

	fmt.Println("Complete")
}

func TestGetDataWaitWithWg(t *testing.T) {
	size := 100
	dataCh := prepareData(size)
	wg := sync.WaitGroup{}
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func() {
			mockBiz(dataCh)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Complete")
}

func mockBiz(ch <-chan int) {
	// 模拟耗时
	time.Sleep(time.Microsecond * 10)
	fmt.Println(<-ch)
}
