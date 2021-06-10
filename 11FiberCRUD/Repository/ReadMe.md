# Ví dụ về Model, Repository

Cần chia dự án thành các phần nhỏ Single Reposibility

```
.
├── app.go <-- file chính
├── go.mod
├── go.sum
├── models.go <-- định nghĩa model, entity hay bảng
└── repository.go <-- định nghĩa các hàm truy vấn, thêm, sửa, xoá dữ liệu
```