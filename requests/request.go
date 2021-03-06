package requests

import (
	"net/http"

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
	Rendered  bool
}
