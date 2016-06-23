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

func addCSRFTokenToResponse(request *requests.Request) {
	csrfToken := GenerateCSRFToken()

	request.CSRFToken = csrfToken

	csrfCookie := new(http.Cookie)
	csrfCookie.Name = "CSRFToken"
	csrfCookie.Value = csrfToken //todo
	http.SetCookie(request.Response, csrfCookie)
}
