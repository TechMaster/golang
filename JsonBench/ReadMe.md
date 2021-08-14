# JSON Unmarshal

Unmarshal là các decode một string kiểu JSON sang một đối tượng. Đối tượng này có thể là kiểu `map[string]interface{}` hoặc một kiểu struct cụ thể.

Ưu điểm của `map[string]interface{}` là nó tuỳ biến động. Nhưng nó có hai nhược điểm lớn:

1. Chậm -> hãy xem phần benchmark dưới đây
2. Không kiểm tra kiểu chặt chẽ. Có thể gây lỗi panic khi truy cập vào một thuộc tính nil

Trong bài viết này chúng ta sẽ chọn ra thư viện JSON nào chạy nhanh nhất trong số

1. encoding/json
2. github.com/goccy/go-json
3. github.com/json-iterator/go

## JSON2Struct nhanh hơn nhiều JSON2Map Và Goccy tốc độ tốt nhất !

Chuyển vào thư mục demo và chạy lệnh `go test -bench .`
```
cd demo       
go test -bench .
```

Kết quả là
```
Benchmark_Goccy_JSON2Struct-8            1442637               826.7 ns/op
Benchmark_Goccy_JSON2Map-8                488211              2304 ns/op

Benchmark_Default_JSON2Struct-8           361782              3146 ns/op
Benchmark_Default_JSON2Map-8              222560              5188 ns/op

Benchmark_Jsoniter_JSON2Struct-8          812524              1355 ns/op
Benchmark_Jsoniter_JSON2Map-8             387609              2794 ns/op
```

Chưa kể github.com/goccy/go-json mặc định tương thích với thư viện chuẩn "encoding/json"
**Hãy chọn github.com/goccy/go-json và Unmarshall ra một struct cụ thể để có tốc độ tốt nhất**