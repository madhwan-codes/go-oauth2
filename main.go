package main

import (
	"github.com/labstack/echo/v4"
	"github.com/madhwan-codes/go-oauth2/config"
	"github.com/madhwan-codes/go-oauth2/controllers"
)

func main() {
	e := echo.New()

	config.GoogleConfig()
	config.GithubConfig()
	config.FacebookConfig()

	e.GET("/", controllers.Ping)
	e.GET("/google_login", controllers.GoogleLogin)
	e.GET("/google_callback", controllers.GoogleCallback)
	e.GET("/github_login", controllers.GithubLogin)
	e.GET("/github_callback", controllers.GithubCallback)
	e.GET("/facebook_login", controllers.FacebookLogin)
	e.GET("/facebook_callback", controllers.FacebookCallback)

	e.Logger.Fatal(e.Start(":8080"))
}
