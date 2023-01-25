go test -bench=^BenchmarkExample$ -benchtime=1000x -benchmem -cpu=2,3,4 .
go test -bench=Concat -benchtime=1s -benchmem .
