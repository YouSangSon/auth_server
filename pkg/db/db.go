package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
)

func DBInit(dbType string) (*gorm.DB, *redis.Client, error) {
	var db *gorm.DB
	var err error

	switch dbType {
	case "postgres":
		postgresDsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			os.Getenv("HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))

		newDBLogger := dbLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			dbLogger.Config{
				SlowThreshold: time.Second,
				LogLevel:      dbLogger.Error,
				Colorful:      true,
			},
		)

		db, err = gorm.Open(postgres.Open(postgresDsn), &gorm.Config{
			Logger: newDBLogger,
		})
		if err != nil {
			return nil, nil, err
		}
	}

	rdb, err := CacheDBInit()
	if err != nil {
		return nil, nil, err
	}

	return db, rdb, nil
}
