package app

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type App struct {
	ctx               context.Context
	postgresDB        *gorm.DB
	redisDB           *redis.Client
	googleLoginConfig oauth2.Config
	githubLoginConfig oauth2.Config
	Base              fiber.Router
	Cache             fiber.Router
	// NaverLoginConfig  oauth2.Config
	// KaKaoLoginConfig  oauth2.Config

}

// var App app

func NewApp(ctx context.Context, db *gorm.DB, rdb *redis.Client, google oauth2.Config, github oauth2.Config, v1, v2 fiber.Router) App {
	return App{
		ctx:               ctx,
		postgresDB:        db,
		redisDB:           rdb,
		googleLoginConfig: google,
		githubLoginConfig: github,
		Base:              v1,
		Cache:             v2,
	}
}

func (app *App) Start() {

}
