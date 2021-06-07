# Để hiểu về Router trong Fiber


## 1. Router trong Fiber là một interface

Nếu bạn chưa hiểu về interface trong ngôn ngữ Golang, [hãy xem link này](https://tour.golang.org/methods/9). Có 2 điểm chính

1. An interface type is defined as a set of method signatures. Kiểu interface định nghĩa một tập các mẫu hàm. Mẫu hàm chỉ có phần khai báo, nhưng không chứa phần thực hiện chi tiết trong hàm.

2. A value of interface type can hold any value that implements those methods. Một biến có kiểu interface có thể được gán bởi bất kỳ đối tượng nào implement các hàm đó.

```go
// Router defines all router handle interface includes app and group router.
type Router interface {
	Use(args ...interface{}) Router

	Get(path string, handlers ...Handler) Router
	Head(path string, handlers ...Handler) Router
	Post(path string, handlers ...Handler) Router
	Put(path string, handlers ...Handler) Router
	Delete(path string, handlers ...Handler) Router
	Connect(path string, handlers ...Handler) Router
	Options(path string, handlers ...Handler) Router
	Trace(path string, handlers ...Handler) Router
	Patch(path string, handlers ...Handler) Router

	Add(method, path string, handlers ...Handler) Router
	Static(prefix, root string, config ...Static) Router
	All(path string, handlers ...Handler) Router

	Group(prefix string, handlers ...Handler) Router

	Mount(prefix string, fiber *App) Router
}
```

# 2. Tìm hiểu ```func (app *App) Get(path string, handlers ...Handler) Router```

Bên trong hàm này là 2 hàm ```app.Add(MethodHead)``` nối tiếp bởi ```Add(MethodGet)```

```go
func (app *App) Get(path string, handlers ...Handler) Router {
	return app.Add(MethodHead, path, handlers...).Add(MethodGet, path, handlers...)
}
```

Tiếp tục truy sâu hơn đến hàm ```func (app *App) Add```
```go
func (app *App) Add(method, path string, handlers ...Handler) Router {
	return app.register(method, path, handlers...)
}
```

Truy sâu hơn nữa bạn sẽ thấy register trả về chính đối tượng app.
```go
func (app *App) register(method, pathRaw string, handlers ...Handler) Router {
  // Tạm bỏ qua một đống code phía trên
  // dòng cuối cùng lại trả về chính app.
	return app  
}
```

Kiểu viết này để linh hoạt nối chuỗi các hàm xử lý request.

## 3. Router khác gì Route?
Router là interface định nghĩa các mẫu hàm. Còn Route là struct, cấu trúc dữ liệu thực sự lưu đường dẫn sau khi ghép nối hoàn chỉnh và mảng các hàm xử lý request ```Handlers []Handler```

Mỗi lần bạn gọi một hàm trong Router interface, thực chất bạn đã bổ xung thêm một biến Route vào một slice 2 chiều
```stack [][]*Route``` trong ```App struct```

```go
type Route struct {
	// Data for routing
	pos         uint32      // Position in stack -> important for the sort of the matched routes
	use         bool        // USE matches path prefixes
	star        bool        // Path equals '*'
	root        bool        // Path equals '/'
	path        string      // Prettified path
	routeParser routeParser // Parameter parser

	// Public fields
	Method   string    `json:"method"` // HTTP method
	Path     string    `json:"path"`   // Original registered route path
	Params   []string  `json:"params"` // Case sensitive param keys
	Handlers []Handler `json:"-"`      // Ctx handlers
}
```

## 4. Cấu trúc của slice 2 chiều ```stack [][]*Route```
Định nghĩa của nó như thế này. Đọc chú thích Route stack divided by HTTP methods, bạn có thể hiểu chiều thứ nhất của slide lưu theo danh sách các HTTP methods. Tương ứng với mỗi HTTP method cụ thể sẽ là một slice lưu tất cả Route xử lý request khác đường dẫn nhưng cùng tính chất HTTP method. Việc này để làm gì? Để tăng tốc độ lấy handler phù hợp theo 2 tiêu chí: HTTP method và path

```go
type App struct {
	// Route stack divided by HTTP methods
	stack [][]*Route
```

Danh sách các HTTP methods
```go
const (
	MethodGet     = "GET"     // RFC 7231, 4.3.1
	MethodHead    = "HEAD"    // RFC 7231, 4.3.2
	MethodPost    = "POST"    // RFC 7231, 4.3.3
	MethodPut     = "PUT"     // RFC 7231, 4.3.4
	MethodPatch   = "PATCH"   // RFC 5789
	MethodDelete  = "DELETE"  // RFC 7231, 4.3.5
	MethodConnect = "CONNECT" // RFC 7231, 4.3.6
	MethodOptions = "OPTIONS" // RFC 7231, 4.3.7
	MethodTrace   = "TRACE"   // RFC 7231, 4.3.8
	methodUse     = "USE"
)
```

