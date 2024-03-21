package config

// type Config struct {
// 	GoogleLoginConfig oauth2.Config
// 	NaverLoginConfig  oauth2.Config
// 	KaKaoLoginConfig  oauth2.Config
// 	GithubLoginConfig oauth2.Config
// }

// var AppConfig Config

// func ConfigLoad() {
// 	err := godotenv.Load("auth_server.env")
// 	if err != nil {
// 		log.Fatalf("Some error occurred. Err: %s", err)
// 	}

// 	AppConfig.GoogleLoginConfig = GoogleConfig()
// }

// func GoogleConfig() oauth2.Config {
// 	AppConfig.GoogleLoginConfig = oauth2.Config{
// 		RedirectURL:  os.Getenv("GOOGLE_REDIRECT"),
// 		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
// 		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
// 		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
// 		Endpoint:     google.Endpoint,
// 	}

// 	return AppConfig.GoogleLoginConfig
// }

// func githubConfig() oauth2.Config {
// 	AppConfig.GithubLoginConfig = oauth2.Config{
// 		RedirectURL:  os.Getenv("GITHUB_REDIRECT"),
// 		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
// 		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
// 		Scopes:       []string{"user:email"},
// 		Endpoint:     google.Endpoint,
// 	}

// 	return AppConfig.GithubLoginConfig
// }
