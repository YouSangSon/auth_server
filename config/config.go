package config

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Conf struct {
	GoogleConf oauth2.Config
	GithubConf oauth2.Config
}

var Configs Conf

func LoadConfig() error {
	err := godotenv.Load("auth_server.env")
	if err != nil {
		return err
	}

	Configs.GoogleConf = googleConfig()
	Configs.GithubConf = githubConfig()

	return nil
}

func googleConfig() oauth2.Config {
	return oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func githubConfig() oauth2.Config {
	return oauth2.Config{
		RedirectURL:  os.Getenv("GITHUB_REDIRECT"),
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"user:email"},
		Endpoint:     google.Endpoint,
	}
}
