# So sánh giữa truyền pointer struct vs value struct

## Khi muốn thay đổi thuộc tính trong struct cần phải truyền pointer struct
```go
func PassStructAsValue(acc Account) {
	acc.name = "John"  //Chỉ có tác dụng bên trong hàm này
}

func PassStructAsPointer(acc *Account) {
	acc.name = "John"  //Thuộc tính giữ giá trị thay đổi khi hàm kết thúc.
}
```
## Benchmark tốc độ: PassStructAsValue chạy nhanh hơn PassStructAsPointer khoảng 4 lần !
```
go test -bench=.
goos: darwin
goarch: amd64
pkg: pointer_value
cpu: Intel(R) Core(TM) i7-4800MQ CPU @ 2.70GHz
BenchmarkPassStructAsValue-8            1000000000               0.2876 ns/op
BenchmarkPassStructAsPointer-8          1000000000               0.8367 ns/op
```
## Tóm lại là

1. Cần phải trả về thay đổi thuộc tính trong struct khi hàm kết thúc dùng truyền pointer struct
2. Cần tốc độ và đơn giản truyền value struct



## Tham khảo
- [Go: Should I Use a Pointer instead of a Copy of my Struct?](https://medium.com/a-journey-with-go/go-should-i-use-a-pointer-instead-of-a-copy-of-my-struct-44b43b104963)
- [Design Philosophy On Data And Semantics](https://www.ardanlabs.com/blog/2017/06/design-philosophy-on-data-and-semantics.html)