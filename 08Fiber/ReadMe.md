1- Ưu tiên dùng VSCode launch package để debug

2- Dùng VSCode mở đúng thư mục, chứ không không mở thư mục cha gồm nhiều dự án golang bên trong

3- Lệnh go run thì chạy được, nhưng khi vào VSCode ấn nút launch package thì fail. Cách sửa
+ Uninstall Golang extension trong VSCode rồi cài lại.

4- Để lấy code mẫu của giảng viên có 2 cách.
+ git pull https://github.com/TechMaster/golang.git
+ sử dụng công cụ https://downgit.github.io/ để tải về một thư mục cụ thể dạng file zip, mà không cần pull cả thư mục git

5- Cần tổ chức thư mục cho hợp lý, chia theo ngày, trong mỗi folder cần gõ file ReadMe.md nốt lại những ý chính.

6- Trên Windows, thì cần chuyển Prefork: true --> Prefork: false

Những kiểu interfacr và struct trong Fiber cần nhớ

- interface Router
- fiber.App là struct
- fiber.Ctx
- Route
- Handler --> func(*Ctx) error
- Stack stack [][]*Route