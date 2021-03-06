package requests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequest(T *testing.T) {
	var token string
	//test adding a token to a response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//make a new Request out of it
		testReq := new(Request)
		testReq.Request = r
		testReq.Response = w
		token = testReq.AddCSRFToken()
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
