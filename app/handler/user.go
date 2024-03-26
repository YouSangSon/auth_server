package handler

import (
	"auth_server/app/common"
	"auth_server/app/response"
	"auth_server/model"
	"net/http"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// ValidateID는 주어진 ID가 이메일 형식이 아니고 최소 5글자 이상인지 검사합니다.
func ValidateID(id string) bool {
	// 이메일 형식을 검사하는 정규 표현식
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	// ID가 이메일 형식이 아니며 최소 5글자 이상인지 검사
	return !emailRegex.MatchString(id) && len(id) >= 5
}

func (hc *HandlerContext) UserHandler() {
	// 회원가입 신청
	hc.Router.Post("/sign-up", hc.signUp)

	// 회원가입 승낙, 거절
	hc.Router.Post("/accept", hc.accept)

	// 로그인
	hc.Router.Post("/sign-in", hc.signIn)
}

func (hc *HandlerContext) accept(c *fiber.Ctx) error {
	a, err := common.GetTokenByRequest(hc.RedisDB, c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{response.RESULT: response.ERROR, response.MESSAGE: response.INVALID_TOKEN})
	}

	user := model.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{response.RESULT: response.ERROR, response.MESSAGE: response.INVALID_DATA})
	}

	dbUser := model.User{}

}

func (hc *HandlerContext) signUp(c *fiber.Ctx) error {
	user := model.User{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{response.RESULT: response.ERROR, response.MESSAGE: response.INVALID_DATA})
	}

	if !ValidateID(user.ID) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{response.RESULT: response.ERROR, response.MESSAGE: response.INVALID_DATA})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MaxCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{response.RESULT: response.ERROR, response.MESSAGE: response.SERVER_ERROR})
	}

	dbUser := model.User{
		ID:       user.ID,
		Password: string(hashedPassword[:]),
		Email:    user.Email,
		Username: user.Username,
		Accepted: false,
	}

	if err := hc.PostgresDB.Raw("INSERT INTO users (id, username, password, email) VALUES (?, ?, ?)", dbUser.ID, dbUser.Username, dbUser.Password, dbUser.Email).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{response.RESULT: response.ERROR, response.MESSAGE: response.SERVER_ERROR})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{response.RESULT: "ok", response.MESSAGE: "success"})
}

func (hc *HandlerContext) signIn(c *fiber.Ctx) error {
	user := model.User{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{response.RESULT: response.ERROR, response.MESSAGE: response.INVALID_DATA})
	}

	dbUser := model.User{}

	if err := hc.PostgresDB.Raw("SELECT * FROM users WHERE id = ?", user.ID).Scan(&dbUser).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{response.RESULT: response.ERROR, response.MESSAGE: response.SERVER_ERROR})
	}

	if !dbUser.Accepted {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{response.RESULT: response.ERROR, response.MESSAGE: response.SERVER_ERROR})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{response.RESULT: response.ERROR, response.MESSAGE: response.SERVER_ERROR})
	}

	_token, err := common.GenerateJWT(dbUser.ID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{response.RESULT: response.ERROR, response.MESSAGE: response.SERVER_ERROR})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{response.RESULT: "ok", response.MESSAGE: "success", response.DATA: _token})
}

// this is what the data would look like, after the convert
// you can use this reference to get the data and save to db for example.
// type GooglePayload struct {
// 	SUB           string `json:"sub"`
// 	Name          string `json:"name"`
// 	GivenName     string `json:"given_name"`
// 	FamilyName    string `json:"family_name"`
// 	Picture       string `json:"picture"`
// 	Email         string `json:"email"`
// 	EmailVerified bool   `json:"email_verified"`
// 	Locale        string `json:"locale"`
// }

// func ConvertToken(accessToken string) (*GooglePayload, error) {
// 	// call http request to this URL, this is a valid
// 	// URL provided from google to convert access token into
// 	// valid user data
// 	resp, httpErr := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", accessToken))
// 	if httpErr != nil {
// 		return nil, httpErr
// 	}

// 	// clean up when this function returns (destroyed)
// 	defer resp.Body.Close()

// 	// Reads the entire HTTP body from resp.Body using ioutil.ReadAll.
// 	// If any error occurs during the read operation, it is
// 	// returned as bodyErr. Otherwise it is stored in the respBody variable.
// 	respBody, bodyErr := ioutil.ReadAll(resp.Body)
// 	if bodyErr != nil {
// 		return nil, bodyErr
// 	}

// 	// Unmarshal raw response body to a map
// 	var body map[string]interface{}
// 	if err := json.Unmarshal(respBody, &body); err != nil {
// 		return nil, err
// 	}

// 	// If json body containing error,
// 	// then the token is indeed invalid. return invalid token err
// 	if body[response.ERROR] != nil {
// 		return nil, errors.New("Invalid token")
// 	}

// 	// Bind JSON into struct
// 	var data GooglePayload
// 	err := json.Unmarshal(respBody, &data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &data, nil
// }

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
