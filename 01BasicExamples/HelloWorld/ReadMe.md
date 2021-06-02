# Tạo Go module

Thường chúng ta ít khi chỉ tạo duy nhất một file *.go rồi dùng lệnh go run. Chúng ta sẽ tạo ra một ứng dụng có cấu trúc thư mục, và import nhiều package từ bên ngoài.

Lúc này cần đến Go mode. Đặc điểm của go module có thể dùng theo 4 cách:

1. Mã nguồn mở. Mã nguồn module lưu git repo công khai
2. Mã nguồn đóng, sử dụng trong dự án nội bộ sử dụng GOPRIVATE, hướng dẫn sử dụng [](https://johanlejdung.medium.com/a-mini-guide-go-modules-and-private-repositories-fa94c3726cf1)
3. Biên dịch Go module thành binary plugin.

Trong bài này chúng ta tập trung vào phương pháp số 1, phổ biến nhất
## 1. Tạo module
```
$ go mod init github.com/techmaster.vn/helloworld
```

file go.mod sẽ được tạo ra với nội dung như sau
```
module github.com/techmaster.vn/helloworld

go 1.16
```
Chú ý: đường dẫn truyền vào ```go mod init``` có 2 khả năng:
1. Thực sự là đường dẫn một github repo
2. Không thực sự là đường dẫn một github repo

Khả năng 1 khi bạn push code lên github, rồi thực hiện lệnh go get thi

Starting with Go 1.13, Go modules are the standard package manager in Golang, automatically enabled on installation along with a default GOPROXY. 
https://jfrog.com/blog/why-goproxy-matters-and-which-to-pick/

https://arslan.io/2019/08/02/why-you-should-use-a-go-module-proxy/




