package security

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/corvuscrypto/birdnest/requests"
)

//GenerateCSRFToken generates and returns a cryptographically random token for use in cross-site request forgery protection
func GenerateCSRFToken() string {
	//256 bit entropy
	token := make([]byte, 32)
	rand.Read(token)
	return base64.StdEncoding.EncodeToString(token)
}

//AddCSRFTokenToResponse generates a CSRF token and adds it into a response's cookie headers using the
//default name of CSRFToken. This function returns the value of the token for convenience.
func AddCSRFTokenToResponse(request *requests.Request) string {
	csrfToken := GenerateCSRFToken()

	request.CSRFToken = csrfToken

	csrfCookie := new(http.Cookie)
	csrfCookie.Name = "CSRFToken"
	csrfCookie.Value = csrfToken //todo
	http.SetCookie(request.Response, csrfCookie)
	return csrfToken
}
