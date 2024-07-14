package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/madhwan-codes/go-oauth2/config"
	"io/ioutil"
	"net/http"
)

func FacebookLogin(c echo.Context) error {
	url := config.AppConfig.FacebookLoginConfig.AuthCodeURL("randomstate")
	log.Info("Facebook Login URL: ", url)
	return c.Redirect(http.StatusSeeOther, url)
}

func FacebookCallback(c echo.Context) error {
	state := c.QueryParam("state")
	if state != "randomstate" {
		return c.String(http.StatusBadRequest, "States don't Match!!")
	}

	code := c.QueryParam("code")

	oauthConfig := config.FacebookConfig()

	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Code-Token Exchange Failed")
	}

	resp, err := http.Get("https://graph.facebook.com/v13.0/me?fields=name&access_token=" + token.AccessToken)
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
