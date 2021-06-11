# 1. Tái cấu trúc dự án

Bài này cải tiến từ bài trước. Cấu trúc thư mục bài trước rất đơn giản, tất cả đều nằm trong package main

```
.
├── app.go <-- file chính
├── go.mod
├── go.sum
├── models.go <-- định nghĩa model, entity hay bảng
└── repository.go <-- định nghĩa các hàm truy vấn, thêm, sửa, xoá dữ liệu
```

Còn bài này, chúng ta bắt đầu xây dựng cấu trúc thư mục để dễ nhìn, dễ bảo trì hơn. Mỗi thư mục là một sub package
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
## 1.1 Cấu trúc mới có những cải tiến nào?

Code ở file [app.go](app.go) chỉ tập trung cấu hình quan trọng cho ứng dụng, chuyển logic sau đây vào những sub package để dễ tìm, dễ bảo trì hơn:
- [routes/ConfigRouter.go](routes/ConfigRouter.go) cấu hình định tuyến (configure routes) theo HTTP Verb (GET, POST, PUT, DELETE...) và đường dẫn tương đối
- [controller/BookController.go](controller/BookController.go) nơi chuyên để viết các hàm xử lý request. Nên tạo thành nhiều file xxxController.go khác nhau để phân các hàm xử lý cho một entity hay một nghiệp vụ nhỏ vào một file.
- [model/Book.go](model/Book.go) định nghĩa Entity. Entity thể hiện một cấu trúc dữ liệu đặc thù của ứng dụng. Nó có thể là Author, Book, Person, Film, Product... Mỗi Entity thường tương ứng với một bảng trong CSDL.
- [repository/BookRepository.go](repository/BookRepository.go) là nơi chuyên viết những hàm xử lý dữ liệu : tạo mới (create new), liệt kê (list all), tìm kiếm theo (find by), sửa (edit), xoá (delete)

 ## 1.2 Cấu trúc thư mục đầy đủ một ứng dụng Golang backend sẽ cần thêm sub package nào nữa?
- cronjob: chuyên chứa logic để chạy lặp lại theo thời gian
- security: logic authentication / authorization
- config: load các cấu hình
- public: lưu các file static như ảnh, css, javascript
...

Bạn có thể tham khảo một [đề xuất cấu trúc thư mục khác cho Fiber ở đây](https://sujit-baniya.gitbook.io/fiber-boilerplate/project-structure)
# 2. Chuyển đổi từ cấu trúc dữ liệu Slice sang Map

Slice (dạng Array động) phù hợp khi truy xuất phần tử theo đánh số vị trí. Ngược lại Map phù hợp với truy xuất phần tử theo key.
### Đầu tiên tôi cấu trúc slice để lưu các book
```go

type BookRepo struct {
	books  []*model.Book 
	autoID int64
}
```

Mỗi lần tìm một book theo id, sẽ phải duyệt qua từng phần tử trong slice

```go
func (r *BookRepo) FindBookById(Id int64) (*model.Book, error) {
	for _, book := range r.books {
		if book.Id == Id {
			return book, nil
		}
	}
	return nil, errors.New("book not found")
}
```

### Chuyển sang cấu trúc map, tìm kiếm Book theo key của map nhanh hơn nhiều
```go
type BookRepo struct {
	books  map[int64]*model.Book
	autoID int64
}

var Books BookRepo

func init() {
	Books = BookRepo{autoID: 0}
	Books.books = make(map[int64]*model.Book)
	Books.InitData("sql:45312")
}

func (r *BookRepo) FindBookById(Id int64) (*model.Book, error) {
	if book, ok := r.books[Id]; ok {
		return book, nil
	} else {
		return nil, errors.New("book not found")
	}
}

func (r *BookRepo) DeleteBookById(Id int64) error {
	if _, ok := r.books[Id]; ok {
		delete(r.books, Id)
		return nil
	} else {
		return errors.New("book not found")
	}
}
```

## 3. Bài tập thực hành

Hãy bổ xung chức năng Review để người dùng có thể đánh giá sách.
Rating có giá trị từ 0 đến 5. Sau khi bổ xung một review thì cần cập nhật lại Rating trung bình ở Book
```go
type Review struct {
	Id      int64   `json:"Id"`
	BookId  int64   `json:"BookId"`
	Comment string  `json:"comment"`
	Rating  int  		`json:"rating"`
}
```