package main

import (
	"errors"
	"fmt"
)

type Repository struct {
	books   []Book2
	authors []Author
}

func (r *Repository) InitData(connection string) {
	fmt.Println("Connect to ", connection)
	r.books = append(r.books, Book2{
		Title: "Dế Mèn Phiêu Lưu Ký",
		Authors: []Author{
			Author{FullName: "Tô Hoài", Country: "Vietnam"},
			Author{FullName: "Hames", Country: "Turkey"},
		},
		Rating: 4.5})

	r.books = append(r.books, Book2{
		Title: "100 năm cô đơn",
		Authors: []Author{
			Author{FullName: "Gabriel Garcia Marquez", Country: "Columbia"},
			Author{FullName: "Ivan", Country: "Russia"},
		},
		Rating: 4.5})
}

func (r *Repository) GetAllBooks() []Book2 {
	return r.books
}

func (r *Repository) FindBookById(Id int) (Book2, error) {
	for _, book := range r.books {
		if book.Id == Id {
			return book, nil
		}
	}
	return Book2{}, errors.New("Book not found")
}
