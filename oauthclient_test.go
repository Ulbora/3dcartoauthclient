package oauthclient3dcart

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"
)

func TestOauth3dc_Authorize(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com", nil) // new(http.Request)
	if err != nil {
		log.Fatal(err)
	}
	var cid = os.Args[2]
	var store = os.Args[5]
	res := httptest.NewRecorder()
	var oc Oauth3dc
	oc.OauthURL = "https://apirest.3dcart.com"
	oc.ClientID = cid
	oc.RedirectURL = "https://google.com"
	oc.State = "ffl"
	oc.Authorize(res, req, store)
}

func TestOauth3dc_Token(t *testing.T) {
	// go test -coverprofile=coverage.out -args clientid secrete code storeurl
	fmt.Println("args: ", os.Args)
	fmt.Println("args len: ", len(os.Args))
	var cid = os.Args[2]
	var sec = os.Args[3]
	var code = os.Args[4]
	//var cid = os.Environ()
	fmt.Println("cid:", cid)
	var oc Oauth3dc
	oc.OauthURL = "https://apirest.3dcart.com"
	oc.ClientID = cid
	oc.Secret = sec
	oc.RedirectURL = "https://google.com"
	res := oc.Token(code)
	fmt.Println("token res: ", res)
	if res.Token == "" {
		t.Fail()
	}
}
