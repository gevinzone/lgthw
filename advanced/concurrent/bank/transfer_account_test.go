package bank

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestAccountV1Transfer(t *testing.T) {
	a := newAccountV1(1)
	b := newAccountV1(2)
	c := newAccountV1(3)
	testAccountableTransfer(t, a, b, c)
}

func TestAccountV2Transfer(t *testing.T) {
	a := newAccountV2(1)
	b := newAccountV2(2)
	c := newAccountV2(3)
	testAccountableTransfer(t, a, b, c)
}

func TestAccountV3Transfer(t *testing.T) {
	a := NewAccountV3(1, defaultBalance)
	b := NewAccountV3(2, defaultBalance)
	c := NewAccountV3(3, defaultBalance)
	testAccountableTransfer(t, a, b, c)
}

func testAccountableTransfer(t *testing.T, a, b, c Accountable) {
	wg := sync.WaitGroup{}

	start := time.Now()

	concurrentAccountableTransfer(c, a, 1, &wg)
	concurrentAccountableTransfer(a, b, 1, &wg)
	concurrentAccountableTransfer(b, c, 1, &wg)

	wg.Wait()
	t.Logf("time: %v", time.Since(start))

	assert.Equal(t, defaultBalance, a.GetBalance())
	assert.Equal(t, defaultBalance, b.GetBalance())
	assert.Equal(t, defaultBalance, c.GetBalance())
}

func concurrentAccountableTransfer(from, to Accountable, amount int64, wg *sync.WaitGroup) {
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			from.Transfer(to, amount)
			wg.Done()
		}()
	}
}

func newAccountV2(id int64) *AccountV2 {
	return &AccountV2{
		Account: Account{
			Id:      id,
			Balance: defaultBalance,
		},
	}
}
