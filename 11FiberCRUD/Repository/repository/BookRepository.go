package repository

import (
	"errors"
	"fmt"

	"github.com/TechMaster/golang/11FiberCRUD/Repository/model"
)

type BookRepository struct {
	books  map[int64]*model.Book
	autoID int64 //đây là biến đếm tự tăng gán giá trị cho id của Book
}

var BookRepo BookRepository //Khai báo biến toàn cục, global variable

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	BookRepo = BookRepository{autoID: 0}
	BookRepo.books = make(map[int64]*model.Book)
	BookRepo.InitData("sql:45312")
}

//Pointer receiver ~ method trong Java. Đối tượng chủ thể là *BookRepo
func (r *BookRepository) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *BookRepository) CreateNewBook(book *model.Book) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	book.Id = nextID
	r.books[nextID] = book //tạo mới một phần tử trong map, gán key bằng nextID
	return nextID
}

func (r *BookRepository) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewBook(&model.Book{
		Title: "Dế Mèn Phiêu Lưu Ký",
		Authors: []model.Author{
			{FullName: "Tô Hoài", Country: "Vietnam"},
			{FullName: "Hames", Country: "Turkey"},
		},
		Rating: 4.5})

	r.CreateNewBook(&model.Book{
		Title: "100 năm cô đơn",
		Authors: []model.Author{
			{FullName: "Gabriel Garcia Marquez", Country: "Columbia"},
			{FullName: "Ivan", Country: "Russia"},
		},
		Rating: 4.5})
}

func (r *BookRepository) GetAllBooks() map[int64]*model.Book {
	return r.books
}

func (r *BookRepository) FindBookById(Id int64) (*model.Book, error) {
	if book, ok := r.books[Id]; ok {
		return book, nil //tìm được
	} else {
		return nil, errors.New("book not found")
	}
}

func (r *BookRepository) DeleteBookById(Id int64) error {
	if _, ok := r.books[Id]; ok {
		delete(r.books, Id)
		return nil
	} else {
		return errors.New("book not found")
	}
}

func (r *BookRepository) UpdateBook(book *model.Book) error {
	if _, ok := r.books[book.Id]; ok {
		r.books[book.Id] = book
		return nil //tìm được
	} else {
		return errors.New("book not found")
	}
}

func (r *BookRepository) Upsert(book *model.Book) int64 {
	if _, ok := r.books[book.Id]; ok {
		r.books[book.Id] = book //tìm thấy thì update
		return book.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateNewBook(book)
	}
}

//Cập nhật average rating của Book
func (r *BookRepository) Update(bookId int64, averageRating float32) error {
	//TODO: cập nhật dữ liệu ở đây
}
