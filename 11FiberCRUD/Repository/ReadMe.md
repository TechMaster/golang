# Ví dụ về Model, Repository

Bài này cải tiến từ bài trước. Cấu trúc thư mục bài trước rất đơn giản, tất cả đều nằm trong package main

```
.
├── app.go <-- file chính
├── go.mod
├── go.sum
├── models.go <-- định nghĩa model, entity hay bảng
└── repository.go <-- định nghĩa các hàm truy vấn, thêm, sửa, xoá dữ liệu
```

Còn bài này, chúng ta bắt đầu xây dựng cấu trúc thư mục để dễ nhìn, dễ bảo trì hơn.
```
.
├── .vscode
│   └── launch.json <-- cấu hình cho VSCode
├── controller <-- Chứa các file handler xử lý các request
│   └── BookController.go <-- Handler dành riêng cho Book
├── model <-- Thư mục chứa các file định nghĩa Entity
│   ├── Author.go
│   └── Book.go
├── public
│   └── dog.jpeg
├── repository <-- Chứa các logic thao tác dữ liệu, trước mặt là lưu in memory đã
│   └── BookRepository.go
├── routes <-- Chứa các cấu hình router
│   └── ConfigRouter.go
├── app.go
├── go.mod
├── go.sum
└── ReadMe.md <-- Dự án nào cũng phải viết ReadMe.md cho thật rõ nhé
```