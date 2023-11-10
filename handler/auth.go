package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type GoogleUserInfo struct {
	// The 'sub' field represents the unique identifier for the user
	Sub           string `json:"sub"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

var googleOauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Endpoint:     google.Endpoint,
	RedirectURL:  "http://localhost:1323/auth/google/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
}

func (h *Handler) GoogleLogin(c echo.Context) error {
	url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *Handler) GoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")

	token, err := googleOauthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		return err
	}

	client := googleOauthConfig.Client(c.Request().Context(), token)

	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	userInfo := GoogleUserInfo{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return err
	}

	c.Logger().Debug(fmt.Sprintf("%+v", userInfo))

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
