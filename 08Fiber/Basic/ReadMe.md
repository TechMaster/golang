# Hướng dẫn xử lý lỗi trong Iris và Fiber
## Trả về lỗi xuất ra giao diện người dùng cuối

Ý tưởng chung là các hàm trong Controller, Package chỉ cần trả về err




file main.go sẽ có hàm Custom Error Handler xử lý
```go
app := fiber.New(
  fiber.Config{
    Views:        html.New("./views", ".html"),
    ErrorHandler: CustomErrorHandler, //Đăng ký hàm xử lý lỗi ở đây
  },
)
```

Hàm xử lý `CustomErrorHandler`
```
func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	var statusCode = 500

	if e, ok := err.(*fiber.Error); ok { //Thử kiểm tra xem có phải là kiểu fiber.Error không
		statusCode = e.Code
	} else if e, ok := err.(*errors.Error); ok { //Thử kiểm tra xem có phải là kiểu errors.Error
		statusCode = e.Code
	}

	if err = ctx.Render("error", fiber.Map{
		"ErrorMessage": err.Error(),
		"StatusCode":   statusCode,
	}); err != nil {
		return ctx.Status(500).SendString("Internal Server Error")
	}

	return nil
}
```
