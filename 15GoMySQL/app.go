package main

import (
	"github.com/TechMaster/golang/15GoMySQL/config"
	"github.com/TechMaster/golang/15GoMySQL/repo"
	"github.com/TechMaster/golang/15GoMySQL/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}
	repo.Connect(config.Config)
	repo.InitMasterData()

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	routes.ConfigRouter(app)

	app.Listen(":3000")
}
