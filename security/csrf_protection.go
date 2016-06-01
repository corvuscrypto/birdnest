package security

import (
	"net/http"

	"github.com/corvuscrypto/birdnest/requests"
)

func generateCSRFToken(request requests.Request) {
	csrfToken := ""
	request.Context.Set("csrftoken", csrfToken)
}

func addCSRFTokenToResponse(request requests.Request) {
	csrfCookie := new(http.Cookie)
	csrfCookie.Name = "csrftoken"
	csrfCookie.Value = "" //todo
	http.SetCookie(request.Response, csrfCookie)
}
