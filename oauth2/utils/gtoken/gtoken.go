package gtoken

import (
	"fmt"
	"gRPC-Remote-Procedure-Call-/oauth2/goauth"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	VERIFY_ACCESS_TOKEN = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	OAUTH_STATE         = "pseudo-random"
)

func IsTokenValid(access_token string) bool {
	_, err := http.Get(VERIFY_ACCESS_TOKEN + access_token)
	if err != nil {
		return false
	} else {
		return true
	}
}

func FetchDataFromAccessToken(access_token string) ([]byte, error) {
	resp, err := http.Get(VERIFY_ACCESS_TOKEN + access_token)
	if err != nil {
		return nil, fmt.Errorf("Error in user info:", err.Error())
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response:", err.Error())
	}
	return data, nil
}

func GetAccessToken(state string, code string) (*oauth2.Token, error) {
	if state != OAUTH_STATE {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := goauth.GetOAuthConfig().Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("Error getting the token: %s", err.Error())
	}
	return token, err
}
