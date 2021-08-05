package db

import (
	"errors"
	"fmt"
)

func Query(query string) error {
	fmt.Println("Làm một số thao tác truy vấn")
	return errors.New("query string is bad") //Ở đây dùng package chuẩn error
}
