package benchmark

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkExample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond)
		_ = fmt.Sprintf("something, %s", "something")
	}
}

func BenchmarkAnother(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond)
		_ = fmt.Sprintf("something, %s", "something")
	}
}

func BenchmarkExample2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Microsecond)
		_ = fmt.Sprintf("something, %s", "something")
	}
}

func BenchmarkExample3(b *testing.B) {
	// 启动性能测试前，可以在这里写与性能测试无关的准备代码
	time.Sleep(time.Millisecond * 400)

	b.ResetTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Microsecond)
		_ = fmt.Sprintf("something, %s", "something")
	}
	b.StopTimer()

	// 可以在这里写与性能测试无关的扫尾代码
	time.Sleep(time.Millisecond * 400)
}

func BenchmarkExample4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 性能测试过程中，有准备工作的开销，可以通过b.StopTimer()和b.StartTime()处理掉
		b.StopTimer()
		time.Sleep(time.Microsecond)
		b.StartTimer()
		time.Sleep(time.Microsecond)
		_ = fmt.Sprintf("something, %s", "something")
	}
}
