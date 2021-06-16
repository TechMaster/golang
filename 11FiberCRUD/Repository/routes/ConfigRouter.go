package routes

import (
	"github.com/TechMaster/golang/11FiberCRUD/Repository/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigBookRouter(router *fiber.Router) {
	//Return all books
	(*router).Get("/", controller.GetAllBook) //Liệt kê

	(*router).Get("/:id", controller.GetBookById) //Xem chi tiết một bản ghi

	(*router).Delete("/:id", controller.DeleteBookById) //Xoá một bản ghi

	(*router).Post("", controller.CreateBook) //INSERT: Tạo một bản ghi

	(*router).Patch("", controller.UpdateBook) //UPDATE: Cập nhật một bản ghi

	(*router).Put("", controller.UpsertBook) //UPSERT: Cập nhật một bản ghi nếu tìm thấy còn không tạo mới
}

func ConfigReviewRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllReviews) //Liệt kê

	(*router).Get("/:id", controller.GetReviewById) //Xem chi tiết một bản ghi

	(*router).Delete("/:id", controller.DeleteReviewById) //Xoá một bản ghi

	(*router).Post("", controller.CreateReview) //INSERT: Tạo một bản ghi

	(*router).Patch("", controller.UpdateReview) //UPDATE: Cập nhật một bản ghi

	(*router).Put("", controller.UpsertReview)
}
