package config

import (
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"log"
	"os"
)

type Config struct {
	GoogleLoginConfig   oauth2.Config
	GitHubLoginConfig   oauth2.Config
	FacebookLoginConfig oauth2.Config
}

var AppConfig Config

func GoogleConfig() oauth2.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	AppConfig.GoogleLoginConfig = oauth2.Config{
		RedirectURL:  "http://localhost:8080/google_callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	return AppConfig.GoogleLoginConfig
}

func GithubConfig() oauth2.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	AppConfig.GitHubLoginConfig = oauth2.Config{
		RedirectURL:  "http://localhost:8080/github_callback",
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"user", "repo"},
		Endpoint:     github.Endpoint,
	}

	return AppConfig.GitHubLoginConfig
}

func FacebookConfig() oauth2.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	AppConfig.FacebookLoginConfig = oauth2.Config{
		RedirectURL:  "http://localhost:8080/facebook_callback",
		ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
		ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
		Scopes:       []string{"email", "public_profile"},
		Endpoint:     facebook.Endpoint,
	}

	return AppConfig.FacebookLoginConfig
}
