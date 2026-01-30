package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type (
	GoogleUserInfo struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		Picture       string `json:"picture"`
	}

	Oauth struct {
		Config *oauth2.Config
	}
)

func New() Oauth {
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  os.Getenv("SERVER_URL") + "/api/v1/auth/google/callback",
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	return Oauth{
		Config: googleOauthConfig,
	}
}

func (o Oauth) GetUserInfo(token *oauth2.Token) (*GoogleUserInfo, error) {
	client := o.Config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get user info, status: %d", resp.StatusCode)
	}

	var userInfo GoogleUserInfo
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func GetConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  os.Getenv("SERVER_URL") + "/api/v1/auth/google/callback",
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func RandomState() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "default-state"
	}
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
}
