
## ArrayList - Stack
```
go test -bench=. array_lis*.go stack.go bitset.go
```

```
goos: linux
goarch: amd64
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkArrayList_Append-8     1000000000               0.05712 ns/op
BenchmarkArrayList_Get-8        478751671                2.326 ns/op
BenchmarkArrayList_Remove-8     360056191                3.236 ns/op
```