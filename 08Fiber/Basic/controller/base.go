package controller

import (
	"demofiber/api"
	"demofiber/email"
	"errors"
	"fmt"

	"github.com/TechMaster/eris"

	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	return c.Render("index", nil)
}

func GenericError(c *fiber.Ctx) error {
	return eris.New("Generic Error")
}

func DemoUnAuthorized(c *fiber.Ctx) error {
	return eris.New("Không thể xác định danh tính")
}
func DemoRESTAPIError(c *fiber.Ctx) error {
	_ = c.Status(fiber.StatusUnauthorized).JSON("Bad Request")
	return eris.New("Bad request").StatusCode(fiber.StatusUnauthorized)
}

func LoginAPI(c *fiber.Ctx) error {
	if err := api.Login(c); err != nil {
		return err
	}
	fmt.Println("Tiếp tục làm")
	return nil
}

func LogOutAPI(c *fiber.Ctx) error {
	if err := api.Logout1(c); err != nil {
		return err
	}
	return nil
}

func DivideZero(c *fiber.Ctx) error {
	return errors.New("Chia cho 0")
}

func DemoPage(c *fiber.Ctx) error {
	viewData := make(fiber.Map)
	viewData["Title"] = "Demo"
	return c.Render("khoahoc/index", viewData)
}

func DemoPanicError(c *fiber.Ctx) error {
	if err := api.Query(c); err != nil {
		return err
	} else {
		viewData := make(fiber.Map)
		viewData["Title"] = "Send Email Success"
		return c.Render("khoahoc/index", viewData)
	}
}

func DemoWarning(c *fiber.Ctx) error {
	if err := email.Foo(); err != nil {
		return err
	} else {
		return nil
	}
}

func DemoCause(c *fiber.Ctx) error {
	if err := email.GotMail(); err != nil {
		return err
	} else {
		return nil
	}
}
