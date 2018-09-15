package oauthclient3dcart

import (
	//"fmt"
	"net/http"
	cm "oauthclient3dcart/common"
)

//Oauth3dc Oauth3dc
type Oauth3dc struct {
	OauthURL    string
	ClientID    string
	RedirectURL string
	State       string
	Secret      string
}

//Response Response
type Response struct {
	Token     string `json:"access_token"`
	UserID    int    `json:"user_id"`
	TokenType string `json:"token_type"`
	Code      int
}

//Authorize Authorize
func (o *Oauth3dc) Authorize(w http.ResponseWriter, r *http.Request, storeURL string) {
	var authURL = o.OauthURL + "/oauth/authorize?client_id=" + o.ClientID + "&redirect_uri=" + o.RedirectURL +
		"&state=" + o.State + "&response_type=code&store_url=" + storeURL
	//fmt.Println("authUrl: ", authURL)
	http.Redirect(w, r, authURL, http.StatusFound)
}

//Token Token
func (o *Oauth3dc) Token(code string) *Response {
	var rtn = new(Response)
	var tokenURL = o.OauthURL + "/oauth/token?grant_type=authorization_code&client_id=" + o.ClientID +
		"&client_secret=" + o.Secret + "&redirect_uri=" + o.RedirectURL + "&code=" + code
	//fmt.Println("tokenURL: ", tokenURL)
	req, fail := cm.GetRequest(tokenURL, http.MethodPost, nil)
	if !fail {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		code := cm.ProcessServiceCall(req, &rtn)
		rtn.Code = code
	}
	return rtn
}
