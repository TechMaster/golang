package controller

import (
	"demofiber/db"

	"github.com/gofiber/fiber/v2"
)

func GetAllPosts(c *fiber.Ctx) error {
	if posts, err := getAllPosts(); err != nil {
		return err
	} else {
		return c.Render("post", fiber.Map{
			"Posts": posts,
		})
	}
}

func getAllPosts() (posts []string, err error) {
	if err := db.Query("SELECT * FROM post"); err != nil {
		return nil, err
	} else {
		return []string{"post1", "post2", "post3"}, nil
	}
}
