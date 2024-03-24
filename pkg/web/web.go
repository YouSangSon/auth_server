package web

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func WebInit() *fiber.App {
	web := fiber.New(fiber.Config{
		AppName:       "Auth Server v1.0",
		CaseSensitive: true,
		ReadTimeout:   15 * time.Second,
		WriteTimeout:  15 * time.Second,
	})

	web.Use(logger.New())
	web.Use(recover.New())

	return web
}
