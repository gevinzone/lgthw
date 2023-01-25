package benchmark

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

// 性能测试说明：
// 以Benchmark开头
// 命令为：go test -bench=Example$ -benchtime=1000x -benchmem -cpu=2,3,4
// 测试结果为：
// goos: darwin
// goarch: arm64
// pkg: github.com/gevinzone/lgthw/loop1/c08Test/benchmark
// BenchmarkExample-2          1000           1157728 ns/op              24 B/op          1 allocs/op
// BenchmarkExample-3          1000           1167774 ns/op              24 B/op          1 allocs/op
// BenchmarkExample-4          1000           1188357 ns/op              25 B/op          1 allocs/op

// 其中，各参数含义分别为：
// -benchtime参数表示执行时间（如-benchtime=1s）或执行次数（如-benchtime=1000x），默认值为-benchtime=1s
// -benchmem 表示显示内存分配情况
// -cpu 修改GOMAXPROCS数值，默认为CPU核数

// 性能测试时，通常要固定执行时间或执行次数再进行测试，以方便对比

// 上面的测试结果中
// - BenchmarkExample-2 的2 为`GOMAXPROCS`的值，通过-cpu参数可以修改，同理BenchmarkExample-3和BenchmarkExample-4分别表示3和4时的情况
// - 1000为执行次数
// - 1157728 ns/op为每次执行花费的时间
// - 24 B/op为每次操作总共分配的内存大小
// - 1 allocs/op表示每次操作分配过几次内存

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

	// 然后这里重新计时
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Microsecond)
		_ = fmt.Sprintf("something, %s", "something")
	}

	// 如果测试之后还有扫尾代码，可以先停止计时
	b.StopTimer()
	// 然后可以在这里写与性能测试无关的扫尾代码
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

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
		_ = buf.String()
	}

}
