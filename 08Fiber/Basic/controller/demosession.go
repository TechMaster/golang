package controller

import (
	"demofiber/cookie_session"
	"demofiber/eris"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SessionCounter(c *fiber.Ctx) error {
	if session, err := cookie_session.SessStore.Get(c); err == nil {
		fmt.Println(session.ID())
		var counter int
		var ok bool

		val := session.Get("counter")
		if counter, ok = val.(int); ok {
			counter += 1
		} else {
			counter = 0
		}
		session.Set("counter", counter)

		if err := session.Save(); err != nil {

			return eris.Wrap(err, "Save Session")
		}
		return c.SendString(fmt.Sprintf("%d", counter))
	} else {
		return err
	}
}

func GetSessionID(c *fiber.Ctx) error {
	if session, err := cookie_session.SessStore.Get(c); err == nil {
		id := session.ID()
		if err := session.Destroy(); err != nil {
			return eris.Cause(err)
		}
		return c.SendString(id)

	} else {
		return c.SendString("invalid session id")
	}
}
