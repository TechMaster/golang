# So sánh giữa truyền pointer struct vs value struct

## 1. Khi muốn thay đổi thuộc tính trong struct cần phải truyền pointer struct
```go
func PassStructAsValue(acc Account) {
	acc.name = "John"  //Chỉ có tác dụng bên trong hàm này
}

func PassStructAsPointer(acc *Account) {
	acc.name = "John"  //Thuộc tính giữ giá trị thay đổi khi hàm kết thúc.
}
```
## 2. Benchmark tốc độ: PassStructAsValue chạy nhanh hơn PassStructAsPointer khoảng 4 lần !
```
$ go test -bench=.
```

Kết quả
```
goos: darwin
goarch: amd64
pkg: pointer_value
cpu: Intel(R) Core(TM) i7-4800MQ CPU @ 2.70GHz
BenchmarkPassStructAsValue-8            1000000000               0.2876 ns/op
BenchmarkPassStructAsPointer-8          1000000000               0.8367 ns/op
```

## 3. Truyền map bằng pointer khác gì value?
```go
func PassMapAsValue(data map[string]interface{}) {
	data["name"] = "John"
	data["email"] = "john@gmail.com"
}

func PassMapAsPointer(data *map[string]interface{}) {
	(*data)["name"] = "Hann"
	(*data)["email"] = "hann@gmail.com"
}
```
Tác dụng thay đổi dữ liệu cả hai đều giống nhau. Cú pháp truyền con trỏ map phức tạp hơn.
Tốc độ chậm hơn một chút
```
Benchmark_PassMapAsValue-8              36649004                30.32 ns/op
Benchmark_PassMapAsPointer-8            36922010                30.36 ns/op
```

## 4. Struct to và phức tạp hơn

Xem thư mục [largestruct.go](large/largestruct.go)

```
cd large
go test -bench .
```

Kết quả
```
BenchmarkPassPersonAsValue-8            1000000000               0.2854 ns/op
BenchmarkPassPersonAsPointer-8          26682351                40.49 ns/op
```

Với struct to, pass pointer struct còn chậm hơn rất nhiều

## 5. Pointer hay Value Receiver
Tôi viết một Repository theo 2 cách: 
1. [pointer struct, pointer receiver](pointer/pointer_struct.go)
2. [value struct, value receiver](value/value_struct.go)

Trong [value_struct.go](value/value_struct.go) có hàm lưu tài khoản mới.
Tôi sẽ phải append tài khoản mới vào slice bên trong `accRepo *AccountRepo`
```go
func (accRepo *AccountRepo) Save(acc AccountNew) (Id string, err error) {
```

Trong trường hợp này, Go static check khuyến cáo phải dùng pointer receiver. Mà như vậy mới đúng vì tôi thay đổi thuộc tính bên trong của `accRepo`

Chuyển vào thư mục pv và chạy benchmark

```
cd pv
go test -bench .
```

Kết quả
```
Benchmark_GetAllPointer-8       1000000000               0.4498 ns/op
Benchmark_GetAllValue-8         1000000000               0.2793 ns/op
Benchmark_ValidatePointer-8        40315             28190 ns/op
Benchmark_ValidateValue-8          41931             27839 ns/op
Benchmark_SavePointer-8               16          67663546 ns/op
Benchmark_SaveValue-8                 16          67766001 ns/op
Benchmark_GetIdPointer-8         9466819               126.1 ns/op
Benchmark_GetIdValue-8           8473074               142.5 ns/op
```

Kết quả cho thấy Value Receiver không chậm hơn so với Pointer Receiver khi struct nhỏ, thậm chí còn chạy nhanh ở 2 trong số 4 bài test !
## 6. Tóm lại là

1. Cần phải trả về thay đổi thuộc tính trong struct khi hàm kết thúc dùng truyền pointer struct
2. **Cần tốc độ và đơn giản truyền value struct**



## 7. Tham khảo
- [Go: Should I Use a Pointer instead of a Copy of my Struct?](https://medium.com/a-journey-with-go/go-should-i-use-a-pointer-instead-of-a-copy-of-my-struct-44b43b104963)
- [Design Philosophy On Data And Semantics](https://www.ardanlabs.com/blog/2017/06/design-philosophy-on-data-and-semantics.html)