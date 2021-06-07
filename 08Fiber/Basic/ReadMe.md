# Căn bản với Fiber

## 1. Khởi tạo dự án

```
$ mkdir BASIC
$ cd BASIC
$ go mod init github.com/TechMaster/golang/08Fiber/Basic
$ go get github.com/gofiber/fiber/v2
```

## 2. Tạo ứng dụng Fiber đơn giản app.go
```go
package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
```

## 3. Refactor code bằng cách khai báo hàm xử lý request

Thay vì viết function anonymous kiểu này
```go
app.Get("/", func(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
})
```

chuyển sang viết thành 
```go
app.Get("/", hello)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
```

## 4. Tham số đường dẫn ```:/name```
Bổ xung ```app.Get("/:name", sayName) // GET /john```

```go
func main() {
	app := fiber.New()

	app.Get("/", hello)

	app.Get("/:name", sayName) // GET /john

	app.Listen(":3000")
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func sayName(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
	return c.SendString(msg) // => Hello john 👋!
}
```
Thử vào http://localhost:3000/John, kết quả in ra

Hello, John 👋!

## 5. Escape và Unescale URL

Thử vào http://localhost:3000/Cường, kết quả in ra

Hello, C%C6%B0%E1%BB%9Dng 👋!

Tham số đường dẫn đã bị escape các ký tự unicode

Sửa lại như sau
```go
func sayName(c *fiber.Ctx) error {
	name, err := url.PathUnescape(c.Params("name"))
	fmt.Println(err)
	fmt.Println(name)

	msg := fmt.Sprintf("Hello, %s 👋!", name)
	return c.SendString(msg) // => Hello john 👋!
}
```
Hello, Cường 👋!

## 6. Đăng ký nhiều hàm xử lý request cho cùng một đường dẫn

Định nghĩa của hàm ```func (app *App) Get``` như sau
```go
func (app *App) Get(path string, handlers ...Handler) Router {
	return app.Add(MethodHead, path, handlers...).Add(MethodGet, path, handlers...)
}
```

Chúng ta thấy ```handlers ...Handler``` có nghĩa hàm này sẽ nhận 1 hoặc nhiều handler (variadic function)

```go
app.Get("/bye/:name", log, bye)

func log(c *fiber.Ctx) error {
	fmt.Println("Log: " + c.Params("name"))
	return c.Next()
}

func bye(c *fiber.Ctx) error {
	msg := fmt.Sprintf("good bye %s 👋!", c.Params("name"))
	return c.SendString(msg) // => good bye john 👋!
}
```

