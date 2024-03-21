package controllers

// func GoogleLogin(c *fiber.Ctx) error {
// 	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomState")

// 	fmt.Println(url)

// 	c.Status(fiber.StatusSeeOther)
// 	c.Redirect(url)
// 	return c.JSON(url)
// }

// func GoogleCallback(c *fiber.Ctx) error {
// 	state := c.Query("state")
// 	if state != "randomState" {
// 		c.SendString("States don't Match!!")
// 		return nil
// 	}

// 	code := c.Query("code")

// 	googleConfig := config.GoogleConfig()

// 	token, err := googleConfig.Exchange(context.Background(), code)
// 	if err != nil {
// 		c.SendString("Code-Token Exchange Failed")
// 		return err
// 	}

// 	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
// 	if err != nil {
// 		c.SendString("User Data Fetch Failed")
// 		return err
// 	}

// 	userData, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		c.SendString("JSON Parsing Failed")
// 		return err
// 	}

// 	return c.JSON(string(userData))
// 	// return c.JSON(string(userData)
// }
