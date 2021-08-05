package controller

import (
	"demofiber/db"

	"github.com/rotisserie/eris"
)

func GetAllPosts() (posts []string, err error) {
	if err := db.Query("SELECT * FROM post"); err != nil {
		return nil, eris.Wrap(err, "Cannot get posts")
		//return nil, err
	} else {
		return []string{"post1", "post2", "post3"}, nil
	}
}
