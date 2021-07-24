# Golang 

## Mô tả khoá học
Đây là khoá học đào tạo lập trình viên Golang xây dựng REST API trong hệ thống.

Đồ án mẫu là xây dựng một 
## Yêu cầu đầu vào
1. Sinh viên cần sử dụng hệ điều hành Linux hoặc Mac để lập trình Golang
2. Cần biết hoặc sẽ phải tự học Docker, Docker Compose, Docker Swarm


## 01. Cài đặt môi trường - công cụ và nhập môn
- Cài đặt Golang
- Sử dụng VSCode viết ứng dụng Golang đầu tiên
- In ra màn hình
- Tạo go module
- Khai báo biến: kiểu biến, gán, gán biến tự đoán kiểu
- Các kiểu căn bản. Ép kiểu
- Biến toàn cục vs cục bộ
- Khai báo hàm
- Hàm có nhiều tham số
- Hàm trả về nhiều giá trị, đặt tên biến trả về
- [Hàm khác gì method](https://tutorialedge.net/golang/go-methods-tutorial/)
- Hàm variadic params
- Closure (anonymous function)
- Recursion
- Debug

#### Bài tập
- Tìm số nguyên tố
- Tính giai thừa
- In chuỗi fibonacci

## 02. Golang căn bản
- Public vs Private
- Module, tạo go module
- Con trỏ và lấy địa chỉ biến
- Composite Types: Array, Slice, Map, Struct
- Struct: khởi tạo, truyền vào hàm
- Array
- Slice và thao tác trên Slice
- Map và thao tác trên Map
#### Bài tập
- Phân tích dữ liệu trên file json

## 03. Interface - String
- Interfaces: khác biệt giữa Java và Golang
- Pointer Receiver vs Value Receiver
- Đọc file
- Các thao tác xử lý chuỗi: biến đổi, tìm kiếm, lấy một phần

#### Bài tập
- Đọc file đếm từ
- Mô phỏng

## 04. Thao tác file / folder
- Thao tác folder: thư mục hiện thời, di chuyển thư mục, tạo thư mục, xoá thư mục
- Thao tác file: tạo file, sửa tên file, copy file
- Quét file trong thư mục

#### Bài tập
Hãy sử dụng thư viện này https://github.com/xlab/treeprint và bài viết https://flaviocopes.com/go-list-files/ hãy in ra cấu trúc cây thư mục kiểu như sau:

```
.
├── config
│   └── config.go
├── controller
│   ├── ManufactureController.go
│   └── ProductController.go
```

## 05. Concurrency
- Go routine
- Go channel

## 06. Sync package
- [Wait Group](https://gobyexample.com/waitgroups)
- [Rate Limiting](https://gobyexample.com/rate-limiting)
- [Atomic Counter](https://gobyexample.com/atomic-counters)

## 07. Unit Test (2 buổi)
- [An Introduction to Testing in Go](https://tutorialedge.net/golang/intro-testing-in-go/)
- [testify](https://github.com/stretchr/testify)
- [Improving Your Go Tests and Mocks With Testify](https://tutorialedge.net/golang/improving-your-tests-with-testify-go/)
- [gotests](https://github.com/cweill/gotests): gotests makes writing Go tests easy. It's a Golang commandline tool that generates table driven tests based on its target source files' function and method signatures. Any new dependencies in the test files are automatically imported.
- [Benchmark Golang Code](https://tutorialedge.net/golang/benchmarking-your-go-programs/)

## 08. Package phổ biến trong Go (2 buổi)
- [Go cron](https://github.com/go-co-op/gocron) định thời chạy tác vụ
- [Uber Zap](https://github.com/uber-go/zap) logging
- [Viber configuration](https://github.com/spf13/viper) cấu hình
- [Cobra](https://github.com/spf13/cobra)
- [Pterm](https://github.com/pterm/pterm)

Bài tập thực hành

## 09. Resty Client
- HTTP Verbs: GET, POST, PUT, DELETE
- Status Code
- Header vs Body
- Retry
- [Resty Client](https://github.com/go-resty/resty): Simple HTTP and REST client library for Go (inspired by Ruby rest-client)
- [An introduction to REST API testing in Go with Resty](https://www.ontestautomation.com/an-introduction-to-rest-api-testing-in-go-with-resty/)


## 10. Fiber (3 buổi)
- Application
- Application Context
- Router - Routing - Group
- Middle ware
- Custom middle ware

## 11. Go Swagger (1 buổi)

## 12. GORM (2 buổi)
- Định nghĩa Model
- Quan hệ 1:nhiều, nhiều:nhiều
- Thêm sửa xoá truy vấn
- Transaction

## 13. JWT (1 buổi)
- [Securing Your Go REST APIs With JWTs](https://tutorialedge.net/golang/authenticating-golang-rest-api-with-jwts/)


## 14. SOLID Pattern
- [SOLID principle in GO](https://s8sg.medium.com/solid-principle-in-go-e1a624290346)
- [SOLID Go Design](https://dave.cheney.net/2016/08/20/solid-go-design)
- [SOLID : Interface Segregation Principle in Golang](https://medium.com/@felipedutratine/solid-interface-segregation-principle-in-golang-49d4bbb4d3f7)
- [SOLID principles in Golang](https://github.com/ammorteza/SOLID-principles-in-Golang)
- [Design Patterns-Seven Design Principles for Golang](https://www.programmersought.com/article/46554309204/)

## 15. Design Patterns In Go (2 buổi)
- Creational Patterns
- Structural Patterns
- Behavioral Patterns



## Reading List
https://github.com/ardanlabs/gotraining/tree/master/reading

Thread Safe Counting
https://brunocalza.me/there-are-many-ways-to-safely-count/
