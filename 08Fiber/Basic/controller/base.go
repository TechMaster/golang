package controller

import (
	"demofiber/email"
	"errors"

	"demofiber/eris"

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
	return c.Status(fiber.StatusUnauthorized).JSON("Bad Request")
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
	if err := email.SendEmail("cuong@techmaster.vn", "Apply job", "My CV"); err != nil {
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
