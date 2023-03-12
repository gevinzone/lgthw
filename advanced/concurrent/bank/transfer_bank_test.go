package bank

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestBank_TransferAccountV1(t *testing.T) {
	a := newAccountV1(1)
	b := newAccountV1(2)
	c := newAccountV1(3)

	testBankTransfer(t, a, b, c)
}

func testBankTransfer(t *testing.T, a, b, c Accountable) {

	bank := NewBank()

	//t.Logf("Before -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
	wg := sync.WaitGroup{}

	start := time.Now()

	concurrentBankTransfer(bank, c, a, 1, &wg)
	concurrentBankTransfer(bank, a, b, 1, &wg)
	concurrentBankTransfer(bank, b, c, 1, &wg)

	wg.Wait()
	//t.Logf("After -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
	t.Logf("time: %v", time.Since(start))
	assert.Equal(t, defaultBalance, a.GetBalance())
	assert.Equal(t, defaultBalance, b.GetBalance())
	assert.Equal(t, defaultBalance, c.GetBalance())
	bank.Close()
}

func TestBank_TransferAccountV1Async(t *testing.T) {
	a := newAccountV1(1)
	b := newAccountV1(2)
	c := newAccountV1(3)

	testBankTransferAsync(t, a, b, c)
}

func testBankTransferAsync(t *testing.T, a, b, c Accountable) {

	bank := NewBank()

	//t.Logf("Before -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)

	//start := time.Now()

	concurrentBankTransferAsync(bank, c, a, 1)
	concurrentBankTransferAsync(bank, a, b, 1)
	concurrentBankTransferAsync(bank, b, c, 1)

	time.Sleep(time.Millisecond * 2)
	//t.Logf("After -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
	//t.Logf("time: %v", time.Since(start))

	assert.Equal(t, defaultBalance, a.GetBalance())
	assert.Equal(t, defaultBalance, b.GetBalance())
	assert.Equal(t, defaultBalance, c.GetBalance())
	bank.Close()
}

func concurrentBankTransfer(bank *Bank, from, to Accountable, amount int64, wg *sync.WaitGroup) {
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			bank.Transfer(from, to, amount)
			wg.Done()
		}()
	}
}

func concurrentBankTransferAsync(bank *Bank, from, to Accountable, amount int64) {
	for i := 0; i < 500; i++ {
		go func() {
			bank.TransferAsync(from, to, amount)
		}()
	}
}
