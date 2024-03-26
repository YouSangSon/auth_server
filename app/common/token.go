package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"auth_server/app/response"
	"auth_server/model"
	"auth_server/pkg/logger"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("welcome to the yousang's auth server")

type Claims struct {
	Email string `json:"email"`
	jwt.Claims
}

func GenerateJWT(id string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24시간 후 만료

	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)

	return token.SignedString(jwtSecret)
}

// VerifyToken checks if the token is valid and returns the claims
func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}

// GetTokenByRequest 요청으로 부터 유저 토큰 얻어오기
func GetTokenByRequest(session *redis.Client, c *fiber.Ctx) (*model.Token, error) {
	tokenHeader := c.Cookies("__token")

	// 올바른 토큰인지 확인
	var token *model.Token

	path := c.Request().URI().String()

	if strings.Contains(path, "/api/v1") {
		// 쿠키나 헤더에서 토큰 확인
		if tokenHeader == "" {
			tokenHeaders := c.GetReqHeaders()
			authentication := tokenHeaders["authentication"]
			tokenHeader = authentication[0]

			if tokenHeader == "" {
				return nil, errors.New("authentication required")
			}
		} else {
			c.Cookie(&fiber.Cookie{
				Name:     "__token",
				Value:    tokenHeader,
				Expires:  time.Now().Add(60 * 60 * 24 * 3 * time.Second),
				Path:     "/",
				Secure:   true,
				SameSite: "None",
				HTTPOnly: true,
			})
		}
		token, err = u.GetV1TokenByAuthentication(tokenHeader)

	} else if strings.Contains(path, "/api/v2") {
		tokenHeader = c.GetHeader("authentication")
		if tokenHeader == "" {
			logger.Logger.Info(c.Request.Header)
			return nil, errors.New("authentication required")
		}
		token, err = u.GetV2TokenByAuthentication(tokenHeader, withApiKey(c.Query("key")))
	} else {
		token, err = u.GetV1TokenByAuthentication(tokenHeader)
	}

	if err != nil {
		return nil, err
	}

	// 토큰에서 유저 가져오기
	user, err := u.GetUserByToken(token)

	if err != nil {
		return nil, err
	}

	// 유효한 토큰인지 확인
	if user.ServiceUser != 1 && token.Token_type != model.SERVICE && token.Token_type != model.EMS && strings.HasPrefix(c.Request.URL.String(), "/api/v1") {
		buffer := bytes.Buffer{}
		buffer.WriteString(response.BAD_REQUEST)
		//buffer.WriteString("invalid access")
		return nil, errors.New(buffer.String())
	}

	return token, nil
}

func GetV1TokenByAuthentication(tokenHeader string, options ...Option) (*model.Token, error) {
	// 토큰 디코딩
	strToken, err := Decrypt(tokenHeader)
	if err != nil {
		buffer := bytes.Buffer{}
		buffer.WriteString(response.INVALID_TOKEN)
		return nil, errors.New(buffer.String())
	}
	token := model.Token{}
	json.Unmarshal([]byte(strToken), &token)

	sessionKey := u.GetV1TokenSessionKey(&token)

	// redis에서 토큰 가져오기
	exists, err := u.Session.SIsMember(u.Ctx, sessionKey, tokenHeader).Result()

	// 이미 존재하는 유효하는 key라면 유지시간 증가
	if exists && token.Token_type == model.EMS {
		if err != nil || !exists {
			return nil, fmt.Errorf(response.INVALID_TOKEN)
		}

		u.Session.Expire(u.Ctx, sessionKey, time.Hour*24*3)
	}

	return &token, nil
}

// 토큰으로 부터 유저 모델 얻어오기
func GetUserByToken(token *model.Token) (*model.User, error) {
	var user = model.User{}

	u.ORM.Where("email = ?", token.User_email).First(&user)
	if user.Email == "" {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
