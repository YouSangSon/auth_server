package main

import (
	"auth_server/config"
	"auth_server/controller"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	config.GoogleConfig()

	app.Get("/google_login", func(c *fiber.Ctx) {
		controller.GoogleLogin(c)
	})

	app.Listen(":8000")
}
