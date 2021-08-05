package cookie_session

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var SessStore *session.Store

func init() {
	SessStore = session.New(session.Config{
		KeyLookup:  "cookie:session_id",
		Expiration: 24 * time.Hour,
	})
}

/*
func InitSession() *redis.Storage {
	redisDB := redis.New(redis.Config{
		Host:     "localhost",
		Port:     6379,
		Username: "",
		Password: "123",
		Database: 0,
		Reset:    false,
	})

	Sess.Storage = redisDB
	return redisDB
}*/
