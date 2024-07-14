package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/madhwan-codes/go-oauth2/config"
	"io/ioutil"
	"net/http"
)

func GoogleLogin(c echo.Context) error {
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")
	log.Info("Google Login URL: ", url)
	return c.Redirect(http.StatusSeeOther, url)
}

func GoogleCallback(c echo.Context) error {
	state := c.QueryParam("state")
	if state != "randomstate" {
		return c.String(http.StatusBadRequest, "States don't Match!!")
	}

	code := c.QueryParam("code")

	googlecon := config.GoogleConfig()

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Code-Token Exchange Failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.String(http.StatusInternalServerError, "User Data Fetch Failed")
	}
	defer resp.Body.Close()

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, "JSON Parsing Failed")
	}

	return c.String(http.StatusOK, string(userData))
}
