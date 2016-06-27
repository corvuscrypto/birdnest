package requests

import (
	"net/http"

	"github.com/corvuscrypto/birdnest/security"
	"github.com/corvuscrypto/birdnest/sessions"
	"github.com/julienschmidt/httprouter"
)

//Request is a wrapper containing the base request information along with context
type Request struct {
	*http.Request
	Response  http.ResponseWriter
	Ctx       Context
	Session   *sessions.Session
	CSRFToken string
	Params    httprouter.Params
}

//AddCSRFToken generates a CSRF token and adds it into a response's cookie headers using the
//default name of CSRFToken. This function returns the value of the token for convenience.
func (request Request) AddCSRFToken() string {
	csrfToken := security.GenerateCSRFToken()

	request.Ctx.Set("CSRFToken", csrfToken)

	csrfCookie := new(http.Cookie)
	csrfCookie.Name = "CSRFToken"
	csrfCookie.Value = csrfToken //todo
	http.SetCookie(request.Response, csrfCookie)
	return csrfToken
}
