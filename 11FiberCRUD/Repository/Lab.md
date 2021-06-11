1- Pattern Single Resposibility

Chia dự án ra thành nhiều sub package, trong mỗi sub package có thể có vài file *.go

Mỗi sub package có một chức năng cụ thể, gọi tên ra được

- Routes
- Controller
- Models
- Repository: thêm, liệt kê, tìm theo trường, xem, sửa, xoá

Vào xem file app.go so sánh với file app.go buổi trước
```go
package main

import (
	"github.com/gofiber/fiber/v2"
)

var repository Repository

func buildRestV1() (appREST *fiber.App) {
	appREST = fiber.New()
	appREST.Get("/books", func(c *fiber.Ctx) error {
		var books []Book
		books = append(books, Book{Title: "Dế Mèn Phiêu Lưu Ký", Author: "Tô Hoài", Rating: 4.5})
		books = append(books, Book{Title: "Clean Code", Author: "Robert Cecil Martin", Rating: 4.4})
		return c.JSON(books)
	})
	return
}

func buildRestV2() (appREST *fiber.App) {
	appREST = fiber.New()
	appREST.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(repository.GetAllBooks())
	})

	appREST.Get("/books/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		book, err := repository.FindBookById(id)
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		return c.JSON(book)
	})

	return
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})
	app.Mount("/v1", buildRestV1()) //http://localhost:3000/v1/books
	app.Mount("/v2", buildRestV2()) //http://localhost:3000/v2/books

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	repository = Repository{}
	repository.InitData("sql:45312")

	app.Listen(":3000")
}
```

## Code khởi tạo dữ liệu ở bài trước

```go
func (r *Repository) InitData(connection string) {
	fmt.Println("Connect to ", connection)
	r.books = append(r.books, Book2{
		Id:    1,
		Title: "Dế Mèn Phiêu Lưu Ký",
		Authors: []Author{
			{FullName: "Tô Hoài", Country: "Vietnam"},
			{FullName: "Hames", Country: "Turkey"},
		},
		Rating: 4.5})
```

## Lệnh UpSert thực chất gồm 2 lệnh: Insert và Update

Nếu Update không tìm thấy ID thì tạo mới bản ghi tương đương Create