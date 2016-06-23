package requests

import (
	"net/http"

	"github.com/corvuscrypto/birdnest/sessions"
)

//Request is a wrapper containing the base request information along with context
type Request struct {
	*http.Request
	Response  http.ResponseWriter
	Context   context
	Session   sessions.Session
	CSRFToken string
}
