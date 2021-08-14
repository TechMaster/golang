package cookie_session

import (
	"encoding/gob"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var Sess *session.Store

func init() {
	Sess = session.New(session.Config{
		KeyLookup:  "cookie:session_id",
		Expiration: 24 * time.Hour,
	})
	gob.Register(Authenticate{})
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
