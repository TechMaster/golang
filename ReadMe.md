# Golang 

## Mô tả khoá học
Đây là khoá học đào tạo lập trình viên Golang xây dựng REST API trong hệ thống microservice
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
- Hàm variadic params
- Closure (anonymous function)
- Recursion
- Debug

#### Bài tập
- Tìm số nguyên tố
- Tính giai thừa
- In chuỗi fibonacci

## 02. Golang căn bản
- Chữa Bài
- Public vs Private
- package và import package
- Con trỏ và lấy địa chỉ biến
- Struct: khởi tạo, truyền vào hàm
- Array
- Slice và thao tác trên Slice
- Map và thao tác trên Map
#### Bài tập
- Phân tích dữ liệu trên file json

## 03. Interface - String
- Chữa bài tập buổi 02
- Interfaces: khác biệt giữa Java và Golang
- Pointer Receiver vs Value Receiver
- Đọc file
- Các thao tác xử lý chuỗi: biến đổi, tìm kiếm, lấy một phần

#### Bài tập
- Đọc file đếm từ
- Mô phỏng

## 04. Regular Expression - Package regexp
- Giải đáp bài tập lần trước
- Quy tắc viết Regular Expression
- Package regexp

#### Bài tập
- Giải thích regex cho IPv4
- Giải thích regex cho domain
- Viết regex cho string chứa tham số cân nặng (kg), chiều cao (m)

## 05. Thao tác file / folder
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

## 06. Concurrency
- Go routine

## 07. Async package

## Unit Test

[gotests](https://github.com/cweill/gotests): gotests makes writing Go tests easy. It's a Golang commandline tool that generates table driven tests based on its target source files' function and method signatures. Any new dependencies in the test files are automatically imported.

## Resty Client
[Resty Client](https://github.com/go-resty/resty): Simple HTTP and REST client library for Go (inspired by Ruby rest-client)

## SOLID Pattern
- [SOLID principle in GO](https://s8sg.medium.com/solid-principle-in-go-e1a624290346)
- [SOLID Go Design](https://dave.cheney.net/2016/08/20/solid-go-design)
- [SOLID : Interface Segregation Principle in Golang](https://medium.com/@felipedutratine/solid-interface-segregation-principle-in-golang-49d4bbb4d3f7)
- [SOLID principles in Golang](https://github.com/ammorteza/SOLID-principles-in-Golang)
- [Design Patterns-Seven Design Principles for Golang](https://www.programmersought.com/article/46554309204/)


## Gửi nhận với Kafka


## Reading List
https://github.com/ardanlabs/gotraining/tree/master/reading

Thread Safe Counting
https://brunocalza.me/there-are-many-ways-to-safely-count/
