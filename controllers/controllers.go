package controllers

import (
	"auth_server/config"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber"
)

func GoogleLogin(c *fiber.Ctx) error {
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")

	fmt.Println(url)

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}

func GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		c.SendString("States don't Match!!")
		return nil
	}

	code := c.Query("code")

	googlecon := config.GoogleConfig()

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		c.SendString("Code-Token Exchange Failed")
		return err
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.SendString("User Data Fetch Failed")
		return err
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.SendString("JSON Parsing Failed")
		return err
	}

	c.SendString(string(userData))
	return nil
}
