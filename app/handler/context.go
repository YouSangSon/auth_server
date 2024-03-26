package handler

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HandlerContext struct {
	Ctx        context.Context
	PostgresDB *gorm.DB
	RedisDB    *redis.Client
	Router     fiber.Router
	// GoogleLoginConfig oauth2.Config
	// GithubLoginConfig oauth2.Config
}
