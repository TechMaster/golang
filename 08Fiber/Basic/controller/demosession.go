package controller

import (
	"demofiber/cookie_session"
	"fmt"
	"time"

	"github.com/TechMaster/eris"

	"github.com/gofiber/fiber/v2"
)

func GetSessionID(c *fiber.Ctx) error {
	if session, err := cookie_session.Sess.Get(c); err == nil {
		id := session.ID()
		if err := session.Destroy(); err != nil {
			return eris.Cause(err)
		}
		return c.SendString(id)

	} else {
		return c.SendString("invalid session id")
	}
}

func SessionCounter(c *fiber.Ctx) error {
	if session, err := cookie_session.Sess.Get(c); err == nil {
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

func SessionAuthenticate(c *fiber.Ctx) error {
	if session, err := cookie_session.Sess.Get(c); err == nil {
		fmt.Println(session.ID())
		var authen cookie_session.Authenticate
		var ok bool

		valAuthen := session.Get("authenticate")

		if authen, ok = valAuthen.(cookie_session.Authenticate); ok {
			fmt.Println(authen.UserId)
		} else {
			authen = cookie_session.Authenticate{
				UserId:      "OX-13",
				Keys:        []string{"Cuong", "Dzung", "Linh"},
				ExpiredTime: time.Now().Add(48 * time.Hour),
			}
		}
		session.Set("authenticate", authen)
		if err := session.Save(); err != nil {
			return eris.Wrap(err, "Save Session")
		}
		session.Destroy().Error()
		return c.SendString(authen.ExpiredTime.String())
	} else {
		return err
	}
}
