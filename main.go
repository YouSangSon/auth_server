package main

import (
	"auth_server/app"
	"auth_server/config"
	"auth_server/pkg/db"
	"auth_server/pkg/logger"
	"auth_server/pkg/web"
	"context"

	"github.com/gofiber/fiber/v2"
)

//type User struct {
//	ID       int    `json:"id"`
//	Username string `json:"username"`
//	Password string `json:"password"`
//	Email    string `json:"email"`
//}

func init() {
	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

	if err := logger.InitLogger(); err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()

	logger.Logger.Info("Auth Server started...")

	db, rdb, err := db.DBInit("postgres")
	if err != nil {
		panic(err)
	}

	web := web.WebInit()

	api := web.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	v2 := api.Group("/v2", func(c *fiber.Ctx) error {
		c.Set("Version", "v2")
		return c.Next()
	})

	newApp := app.NewApp(ctx, db, rdb, config.Configs.GoogleConf, config.Configs.GithubConf, v1, v2)

	newApp.Start()
}

// func signUpHandler(c fiber.Ctx) error {
// 	user := User{}
// 	err := c.Bind().Body(&user)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).SendString("Invalid data")
// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).SendString("Error hashing password")
// 	}

// 	user.Password = string(hashedPassword)

// 	stmt, err := db.Prepare("INSERT INTO users (username, password, email) VALUES (?, ?, ?)")
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).SendString("Error preparing statement")
// 	}

// 	_, err = stmt.Exec(user.Username, user.Password, user.Email)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).SendString("Error saving user")
// 	}

// 	return c.SendString("Success")
// }

// func loginHandler(c fiber.Ctx) error {
// 	user := User{}
// 	err := c.Bind().Body(&user)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).SendString("Invalid data")
// 	}

// 	var dbUser User
// 	err = db.QueryRow("SELECT * FROM users WHERE username = ?", user.Username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password, &dbUser.Email)
// 	if err != nil {
// 		return c.Status(http.StatusNotFound).SendString("User not found")
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
// 	if err != nil {
// 		return c.Status(http.StatusUnauthorized).SendString("Incorrect password")
// 	}

// 	// 토큰 생성 및 발급
// 	// ...

// 	return c.SendString("Success")
// }
