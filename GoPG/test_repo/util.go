package test_repo

import "github.com/go-pg/pg/v10"

func int_in_array(a int, int_arr []int) bool {
	for _, b := range int_arr {
		if b == a {
			return true
		}
	}
	return false
}

func check_err(err error, trans *pg.Tx) bool {
	if err != nil {
		_ = trans.Rollback()
		return false
	}
	return true
}
