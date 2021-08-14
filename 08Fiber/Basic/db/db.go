package db

import (
	"errors"

	"github.com/TechMaster/eris"
)

func Query(query string) error {
	err := eris.New("query string is bad")
	//err := errors.New("query string is bad")
	return err
}

func Bar() error {
	return eris.Warning("Không tìm thấy bản ghi trong CSDL").StatusCode(404)
}

func Query2() error {
	return errors.New("Unable to query")
}
