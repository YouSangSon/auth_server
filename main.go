package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	// config.GoogleConfig()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// app.Get("/google_login", func(c *fiber.Ctx) {
	// 	controllers.GoogleLogin(c)go
	// })

	// app.Get("/google_callback", func(c *fiber.Ctx) {
	// 	controllers.GoogleCallback(c)
	// })

	app.Listen(":8080")
}
