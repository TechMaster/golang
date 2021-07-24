# Sử dụng Go module

Trong bài thực hành này chúng ta tìm hiểu về go module.
Ý nghĩa của go module là để đóng gói chức năng và tái sử dụng.

Tôi sẽ tạo 3 module khác nhau:
1. greeting
2. foo
3. mygomodule

greeting và foo là hai module tôi không muốn công khai. Còn mygomodule là module tôi sẽ công khai mã nguồn trên github.

## Tạo mới một module

tạo thư mục greeting
```
$ mkdir greeting
$ cd greeting



cấu trúc ứng dụng
```
├── .vscode
│   └── launch.json
├── foo
│   ├── foo.go
│   └── go.mod
├── greeting
│   ├── go.mod
│   └── greeting.go
├── app
├── app.go
├── go.mod
├── go.sum
└── ReadMe.md
```



