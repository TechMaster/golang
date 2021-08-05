package db

import "github.com/rotisserie/eris"

func Query(query string) error {
	err := eris.New("query string is bad")
	//err := errors.New("query string is bad")
	return err
}
