package extend

import (
	"math/rand"
	"time"
)

func ReverseString(s string) string {
	ss := []rune(s)
	left, right := 0, len(ss)-1
	for left < right {
		ss[left], ss[right] = ss[right], ss[left]
		left++
		right--
	}
	return string(ss)
}

func createRandomString(length int, randIntn func(int) int) string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	res := make([]byte, length)
	for i := 0; i < length; i++ {
		res[i] = bytes[randIntn(len(bytes))]
	}
	return string(res)
}

func CreateRandomString(length int) string {
	// 不用种子，亦即种子为1
	return createRandomString(length, rand.Intn)
}

func CreateRandomStringV2(length int) string {
	// 用种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return createRandomString(length, r.Intn)
}
