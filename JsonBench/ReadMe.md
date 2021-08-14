

## JSON2Struct nhanh hơn nhiều JSON2Map Và Goccy tốc độ tốt nhất !
```
Benchmark_Goccy_JSON2Struct-8            1442637               826.7 ns/op
Benchmark_Goccy_JSON2Map-8                488211              2304 ns/op

Benchmark_Default_JSON2Struct-8           361782              3146 ns/op
Benchmark_Default_JSON2Map-8              222560              5188 ns/op

Benchmark_Jsoniter_JSON2Struct-8          812524              1355 ns/op
Benchmark_Jsoniter_JSON2Map-8             387609              2794 ns/op
```

Chưa kể github.com/goccy/go-json mặc định tương thích với thư viện chuẩn "encoding/json"
**Vậy từ bây giờ hãy luôn chọn github.com/goccy/go-json cho dự án của bạn nhé**