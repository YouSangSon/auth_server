package main

// godotenv.Load()

// server := fiber.New()

// server.Get("/", func(c *fiber.Ctx) error {
// 	return c.SendString("Hello from Fiber!")
// })

// server.Listen(":8000")

// server.Get("/google", func(c *fiber.Ctx) error {
// 	return c.SendString("Hello from google!")
// })

// server.Get("/login/oauth2/code/google", func(c *fiber.Ctx) error {
// 	return c.SendString("Hello from google callback!")
// })

// // create a config for google config
// conf := &oauth2.Config{
// 	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
// 	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
// 	RedirectURL:  os.Getenv("GOOGLE_REDIRECT"),
// 	Endpoint:     google.Endpoint,
// 	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
// }

// server.Get("/google", func(c *fiber.Ctx) error {
// 	// create url for auth process.
// 	// we can pass state as someway to identify
// 	// and validate the login process.
// 	URL := conf.AuthCodeURL("not-implemented-yet")

// 	// redirect to the google authentication URL
// 	return c.Redirect(URL)
// })

// server.Get("/login/oauth2/code/google", func(c *fiber.Ctx) error {
// 	// get auth code from the query
// 	code := c.Query("code")

// 	// exchange the auth code that retrieved from google via
// 	// URL query parameter into an access token.
// 	token, err := conf.Exchange(c.Context(), code)
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusInternalServerError)
// 	}

// 	// convert token to user data
// 	profile, err := libs.ConvertToken(token.AccessToken)
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusInternalServerError)
// 	}

// 	return c.JSON(profile)
// })
