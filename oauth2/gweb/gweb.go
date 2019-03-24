package gweb

import (
	"context"
	"fmt"
	"gRPC-Remote-Procedure-Call-/oauth2/goauth"
	"gRPC-Remote-Procedure-Call-/oauth2/proto"
	"gRPC-Remote-Procedure-Call-/oauth2/utils/gtoken"
	"log"

	"net/http"
)

const (
	CODE  = "code"
	STATE = "state"
)

var sc oauthpb.OAuthServiceClient

func WebApisForOAuth(temp oauthpb.OAuthServiceClient) {
	sc = temp
	goauth.GoogleAuthInit()
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/callback", callback)
	http.ListenAndServe(":7007", nil)
	log.Println("Server start with port:7007")
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html>
    <body>
	<a href="/auth">Google Log In</a>
    </body>
	</html>`
	fmt.Fprintf(w, htmlIndex)
}

func auth(w http.ResponseWriter, r *http.Request) {
	url := goauth.GetOAuthConfig().AuthCodeURL(gtoken.OAUTH_STATE)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func callback(w http.ResponseWriter, r *http.Request) {
	CheckOAuthCodeAndState(r.FormValue(STATE), r.FormValue(CODE))
}

func CheckOAuthCodeAndState(state string, code string) {
	req := &oauthpb.OAuthCodeRequest{
		Code:  code,
		State: state,
	}
	res, err := sc.GetCodeState(context.Background(), req)
	if err == nil {
		log.Println("Is Auth Success: ", res.GetIsAuthcode())
	} else {
		log.Println("Error: ", err)
	}
}
