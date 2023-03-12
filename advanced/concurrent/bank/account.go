package bank

import (
	"context"
	"golang.org/x/sync/semaphore"
	"sync"
)

// AccountV1 组合了Account，通过锁，保证线程安全
type AccountV1 struct {
	Account
	lock sync.Mutex
}

var _ Accountable = &AccountV1{}

func (a *AccountV1) Transfer(to Accountable, amount int64) {
	too, ok := to.(*AccountV1)
	if !ok {
		return
	}
	left, right := a, too
	if left.Id > right.Id {
		left, right = too, a
	}

	left.lock.Lock()
	defer left.lock.Unlock()
	right.lock.Lock()
	defer right.lock.Unlock()
	a.transfer(to, amount)
}

// AccountV2 组合了Account，通过借助外部的Allocator，保证线程安全
type AccountV2 struct {
	Account
}

var _ Accountable = &AccountV2{}

func (a *AccountV2) Transfer(to Accountable, amount int64) {
	al := getAllocator()
	al.Apply(a, to)
	a.transfer(to, amount)
	al.Free(a, to)
}

var allocator *Allocator
var allocatorOnce sync.Once

type Allocator struct {
	occupied map[any]bool
	mu       *sync.Mutex
	cond     *sync.Cond
}

func (a *Allocator) Apply(from, to any) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for a.occupied[from] || a.occupied[to] {
		a.cond.Wait()
	}
	a.occupied[from] = true
	a.occupied[to] = true
}

func (a *Allocator) Free(from, to any) {
	a.mu.Lock()
	defer a.mu.Unlock()
	delete(a.occupied, from)
	delete(a.occupied, to)
	a.cond.Broadcast()
}

func getAllocator() *Allocator {
	allocatorOnce.Do(func() {
		mu := &sync.Mutex{}
		allocator = &Allocator{
			occupied: make(map[any]bool, 1000),
			mu:       mu,
			cond:     sync.NewCond(mu),
		}
	})
	return allocator
}

// AccountV3 组合了Account，通过信号量，保证线程安全
type AccountV3 struct {
	Account
	sema *semaphore.Weighted
}

func NewAccountV3(id int64, balance int64) *AccountV3 {
	return &AccountV3{
		Account: Account{
			Id:      id,
			Balance: balance,
		},
		sema: semaphore.NewWeighted(1),
	}
}

var _ Accountable = &AccountV3{}

func (a *AccountV3) Transfer(to Accountable, amount int64) {
	too, ok := to.(*AccountV3)
	if !ok {
		return
	}
	left, right := a, too
	if left.Id > right.Id {
		left, right = too, a
	}

	if err := left.sema.Acquire(context.Background(), 1); err != nil {
		return
	}
	defer left.sema.Release(1)
	if err := right.sema.Acquire(context.Background(), 1); err != nil {
		return
	}
	defer right.sema.Release(1)
	a.transfer(to, amount)
}
