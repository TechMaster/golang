# Ví dụ về func (app *App) Mount(prefix string, fiber *App)
Hàm Mount cho phép chúng ta tạo ra nhiều sub fiber.App rồi gắn lên main fiber.App với mục đích phân cụm.


>Mount attaches another app instance as a sub-router along a routing path.
It's very useful to split up a large API as many independent routers and
compose them as a single service using Mount.

![](diagram.jpg)

```go
app := fiber.New()
app.Mount("/v1", buildRestV1()) //http://localhost:3000/v1/books
app.Mount("/v2", buildRestV2()) //http://localhost:3000/v2/books
```

## 1. Để phục vụ static file, bạn luôn phải sử dụng main fiber.App

```go
app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})
app.Listen(":3000")
```

Nếu bạn cấu hình phục vụ file tĩnh trong sub fiber.App thì cũng không có hiệu lực.

## 2. Luôn cấu hình main fiber.App chứ không cấu hình sub fiber.App

```go
app := fiber.New(fiber.Config{
  Prefork:       true,
  CaseSensitive: true,
  StrictRouting: true,
})
app.Mount("/v1", buildRestV1()) //http://localhost:3000/v1/books
app.Mount("/v2", buildRestV2()) //http://localhost:3000/v2/books


// Luôn cấu hình ở main fiber.App
app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
  Compress:  true,
  ByteRange: true,
  Browse:    true,
  MaxAge:    3600,
})
app.Listen(":3000")
```