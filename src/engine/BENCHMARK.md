
## ArrayList - Stack

```bash
go test -bench=. array_lis*.go stack.go bitset.go
```

```bash
goos: linux
goarch: amd64
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkArrayList_Append-8     1000000000               0.05712 ns/op
BenchmarkArrayList_Get-8        478751671                2.326 ns/op
BenchmarkArrayList_Remove-8     360056191                3.236 ns/op
```

```bash
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i5-8400 CPU @ 2.80GHz
BenchmarkArrayList_Append-6     1000000000               0.06215 ns/op
BenchmarkArrayList_Get-6        631588254                1.941 ns/op
BenchmarkArrayList_Remove-6     234760795                4.342 ns/op
```
