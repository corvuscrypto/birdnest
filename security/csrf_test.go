package security

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/corvuscrypto/birdnest/requests"
)

func TestCSRF(T *testing.T) {

	//generate a CSRF token
	token := GenerateCSRFToken()

	//check to ensure that the encoded length is the expected 44 characters long
	if len(token) != 44 {
		T.Errorf("Unexpected CSRF token length!")
	}

	//test adding a token to a response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//make a new Request out of it
		testReq := new(requests.Request)
		testReq.Request = r
		testReq.Response = w
		token = AddCSRFTokenToResponse(testReq)
		//check to ensure that the encoded length is the expected 44 characters long
		if len(token) != 44 {
			T.Errorf("Unexpected CSRF token length!")
		}
	}))
	defer ts.Close()

	response, _ := http.Get(ts.URL)

	//check to ensure the cookie value and the token match
	csrfCookie := response.Cookies()[0]
	cookieToken := csrfCookie.Value
	if cookieToken != token {
		T.Errorf("Cookie and generated token values do not match!")
	}

}
