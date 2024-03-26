package app

import (
	"auth_server/app/handler"
	"context"
	"time"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func BaseServer(ctx context.Context, db *gorm.DB, rdb *redis.Client) *fiber.App {
	web := fiber.New(fiber.Config{
		AppName:       "Auth Server v1.0",
		CaseSensitive: true,
		ReadTimeout:   15 * time.Second,
		WriteTimeout:  15 * time.Second,
	})

	web.Use(logger.New())
	web.Use(recover.New())

	api := web.Group("/api")
	v1 := api.Group("/v1")

	handlerContext := &handler.HandlerContext{
		Ctx:        ctx,
		PostgresDB: db,
		RedisDB:    rdb,
		Router:     v1,
	}

	handlerContext.UserHandler()

	return web
}
