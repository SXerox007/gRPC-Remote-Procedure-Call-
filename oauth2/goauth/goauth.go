package goauth

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	GOOGLE_AUTH = "https://www.googleapis.com/auth/userinfo.email"
)

var googleOauthConfig *oauth2.Config

func GoogleAuthInit() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "",
		ClientID:     os.Getenv("OAUTH_GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		Scopes:       []string{GOOGLE_AUTH},
		Endpoint:     google.Endpoint,
	}
}

func GetOAuthConfig() *oauth2.Config {
	return googleOauthConfig
}
