package controllers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/madhwan-codes/go-oauth2/config"
	"io/ioutil"
	"net/http"
)

func GithubLogin(c echo.Context) error {

	url := config.AppConfig.GitHubLoginConfig.AuthCodeURL("randomstate")
	log.Info("Github Login URL: ", url)
	return c.Redirect(http.StatusSeeOther, url)
}

func GithubCallback(c echo.Context) error {

	state := c.QueryParam("state")
	if state != "randomstate" {
		return c.String(http.StatusBadRequest, "States don't Match!!")
	}

	code := c.QueryParam("code")

	githubcon := config.GithubConfig()
	fmt.Println(code)

	token, err := githubcon.Exchange(context.Background(), code)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Code-Token Exchange Failed")
	}
	fmt.Println(token)

	//resp, err := http.Get("https://api.github.com/user/repo?access_token=" + token.AccessToken)
	resp, err := func() (*http.Response, error) {
		req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)
		return http.DefaultClient.Do(req)
	}()

	if err != nil {
		return c.String(http.StatusInternalServerError, "User Data Fetch Failed")
	}
	fmt.Println(resp)

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, "JSON Parsing Failed")
	}
	fmt.Println(userData)

	return c.String(http.StatusOK, string(userData))

}
