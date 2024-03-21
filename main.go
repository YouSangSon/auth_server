package main

import (
	"auth_server/config"
	"auth_server/controllers"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	config.GoogleConfig()

	app.Get("/google_login", func(c *fiber.Ctx) {
		controllers.GoogleLogin(c)
	})

	app.Get("/google_callback", func(c *fiber.Ctx) {
		controllers.GoogleCallback(c)
	})

	app.Listen(":8080")
}
