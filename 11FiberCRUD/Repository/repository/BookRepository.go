package repository

import (
	"errors"
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
)

type BookRepo struct {
	books  []*model.Book
	autoID int64
}

var Books BookRepo

func init() {
	Books = BookRepo{autoID: 0}
	Books.InitData("sql:45312")
}

func (r *BookRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *BookRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)
	r.books = append(r.books, &model.Book{
		Id:    r.getAutoID(),
		Title: "Dế Mèn Phiêu Lưu Ký",
		Authors: []model.Author{
			{FullName: "Tô Hoài", Country: "Vietnam"},
			{FullName: "Hames", Country: "Turkey"},
		},
		Rating: 4.5})

	r.books = append(r.books, &model.Book{
		Id:    r.getAutoID(),
		Title: "100 năm cô đơn",
		Authors: []model.Author{
			{FullName: "Gabriel Garcia Marquez", Country: "Columbia"},
			{FullName: "Ivan", Country: "Russia"},
		},
		Rating: 4.5})
}

func (r *BookRepo) GetAllBooks() []*model.Book {
	return r.books
}

func (r *BookRepo) FindBookById(Id int64) (*model.Book, error) {
	for _, book := range r.books {
		if book.Id == Id {
			return book, nil
		}
	}
	return nil, errors.New("book not found")
}
