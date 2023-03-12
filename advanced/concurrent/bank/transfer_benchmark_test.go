package bank

import "testing"

var defaultBalance int64 = 1000

/*
goos: darwin
goarch: arm64
pkg: github.com/igevin/account-transfer-go-example/bank
BenchmarkAccountV1_Transfer-10             10000              1743 ns/op            1155 B/op          6 allocs/op
BenchmarkAccountV2_Transfer-10             10000              2503 ns/op            1299 B/op          6 allocs/op
BenchmarkAccountV3_Transfer-10             10000              3721 ns/op            1852 B/op         14 allocs/op

goos: darwin
goarch: arm64
pkg: github.com/igevin/account-transfer-go-example/bank
BenchmarkAccountV1_Transfer-10             50000              1796 ns/op            1168 B/op          6 allocs/op
BenchmarkAccountV2_Transfer-10             50000              4585 ns/op            1242 B/op          7 allocs/op
BenchmarkAccountV3_Transfer-10             50000              3132 ns/op            1764 B/op         14 allocs/op
*/

func BenchmarkAccountV1_Transfer(b *testing.B) {
	x, y, z := newAccountV1(1), newAccountV1(2), newAccountV1(3)
	benchmarkAccountTransfer(b, x, y, z)
}

func BenchmarkAccountV2_Transfer(b *testing.B) {
	x, y, z := newAccountV2(1), newAccountV2(2), newAccountV2(3)
	benchmarkAccountTransfer(b, x, y, z)
}

func BenchmarkAccountV3_Transfer(b *testing.B) {
	x, y, z := NewAccountV3(1, defaultBalance), NewAccountV3(2, defaultBalance), NewAccountV3(3, defaultBalance)
	benchmarkAccountTransfer(b, x, y, z)
}

func benchmarkAccountTransfer(b *testing.B, x, y, z Accountable) {
	transferTo := func(from, to Accountable) {
		from.Transfer(to, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go transferTo(x, y)
		go transferTo(y, z)
		go transferTo(z, x)
	}
}

func BenchmarkBank_Transfer(b *testing.B) {
	x, y, z := newAccountV1(1), newAccountV1(2), newAccountV1(3)
	benchmarkBankTransfer(b, x, y, z)
}

func benchmarkBankTransfer(b *testing.B, x, y, z Accountable) {
	bank := NewBank()
	transferTo := func(bank *Bank, from, to Accountable) {
		bank.Transfer(from, to, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go transferTo(bank, x, y)
		go transferTo(bank, y, z)
		go transferTo(bank, z, x)
	}
}

func BenchmarkBank_TransferAsync(b *testing.B) {
	x, y, z := newAccountV1(1), newAccountV1(2), newAccountV1(3)
	benchmarkBankTransferAsync(b, x, y, z)
}

func benchmarkBankTransferAsync(b *testing.B, x, y, z Accountable) {
	bank := NewBank()
	transferTo := func(bank *Bank, from, to Accountable) {
		bank.TransferAsync(from, to, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go transferTo(bank, x, y)
		go transferTo(bank, y, z)
		go transferTo(bank, z, x)
	}
}

func newAccountV1(id int64) *AccountV1 {
	return &AccountV1{
		Account: Account{
			Id:      id,
			Balance: defaultBalance,
		},
	}
}
