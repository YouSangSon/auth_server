package service

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserContext struct {
	Ctx        context.Context
	PostgresDB *gorm.DB
	RedisDB    *redis.Client
	Router     fiber.Router
}
