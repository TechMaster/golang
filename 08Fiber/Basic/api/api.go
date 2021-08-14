package api

import (
	"errors"

	"github.com/TechMaster/eris"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	data := map[string]interface{}{
		"host":  "192.168.1.1",
		"port":  8008,
		"roles": []string{"admin", "editor", "user"},
	}

	return eris.New("Unable connect to login").SetData(data).StatusCode(fiber.StatusUnauthorized).EnableJSON()
}

func Logout1(ctx *fiber.Ctx) error {
	return eris.NewFromMsg(logout2(ctx), "logout1 error").StatusCode(401).EnableJSON()
}

func logout2(ctx *fiber.Ctx) error {
	return logout3(ctx)
}

func logout3(ctx *fiber.Ctx) error {
	return errors.New("Cannot logout")
}

func Query(ctx *fiber.Ctx) error {
	if err := connectDB(); err != nil {
		if eris.IsPanic(err) { //Hãy dùng hàm có sẵn trong eris
			panic(err.Error())
		} else {
			return err
		}
	} else {
		return nil
	}
}

func connectDB() error {
	return eris.Panic("Failed to connect to Postgresql").
		SetData(
			map[string]interface{}{
				"host": "localhost",
				"port": "5432",
			},
		)
}
