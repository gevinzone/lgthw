package math

import "math/big"

var cache = make(map[int]*big.Int)

func Fib(n int) *big.Int {
	if n < 0 {
		return nil
	}
	if n < 2 {
		cache[n] = big.NewInt(1)
	}
	val, ok := cache[n]
	if ok {
		return val
	}
	cache[n] = big.NewInt(0)
	cache[n].Add(cache[n], Fib(n-1))
	cache[n].Add(cache[n], Fib(n-2))
	return cache[n]
}
