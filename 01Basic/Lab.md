go run app.go util.go
go build app.go util.go
./app

go mod init github.com/ocg.com/go/01/lab

go mod tidy: để cập nhật, dọn dẹp các package. Cũ thì nâng lên mới, không dùng đến thì bỏ, import mà chưa có thư viện thì tải về. Cập nhật lại cả file go.sum

Create Go Launch Package tạo ra thư mục .vscode chứa file launch.json