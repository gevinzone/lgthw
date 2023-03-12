package bank

import "runtime"

// Account 是一个线程不安全的结构体，实现了Accountable接口
type Account struct {
	Id      int64
	Balance int64
}

var _ Accountable = &Account{}

func (a *Account) GetId() int64 {
	return a.Id
}

func (a *Account) GetBalance() int64 {
	return a.Balance
}

func (a *Account) SetBalance(amount int64) {
	a.Balance = amount
}

func (a *Account) transfer(to Accountable, amount int64) {
	a.Balance -= amount
	// Gosched yields the processor, allowing other goroutines to run. It does not
	// suspend the current goroutine, so execution resumes automatically.
	runtime.Gosched()
	to.SetBalance(to.GetBalance() + amount)
}

func (a *Account) Transfer(to Accountable, amount int64) {
	a.transfer(to, amount)
}

type Accountable interface {
	GetId() int64
	GetBalance() int64
	SetBalance(amount int64)
	Transfer(to Accountable, amount int64)
}
