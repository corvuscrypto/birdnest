package server

import (
	"net/http"
	"time"

	"github.com/corvuscrypto/birdnest/config"
	"github.com/corvuscrypto/birdnest/logging"
	"github.com/corvuscrypto/birdnest/requests"
	"github.com/corvuscrypto/birdnest/security"
)

var pipeline []RequestHandler

//RegisterMiddleware registers middleware. Middleware request handlers are utilized in FIFO order.
func RegisterMiddleware(rh ...RequestHandler) {
	if rh == nil {
		return
	}
	pipeline = append(pipeline, rh...)
}

func applyMiddleware(r *requests.Request) {

	//if we have timing enabled we initiate the time here
	if config.Config.GetBool("enableTiming", false) {
		start := time.Now()
		defer func() { logging.GetLogger().Log(logging.DEBUG, time.Since(start)) }()
	}

	for _, f := range pipeline {
		f(r)
	}

}

//AddCSRFToken generates a CSRF token and adds it into a response's cookie headers using the
//default name of CSRFToken. This function returns the value of the token for convenience.
func AddCSRFToken(request *requests.Request) {
	csrfToken := security.GenerateCSRFToken()

	request.Ctx.Set("CSRFToken", csrfToken)

	csrfCookie := new(http.Cookie)
	csrfCookie.Name = "CSRFToken"
	csrfCookie.Value = csrfToken //todo
	http.SetCookie(request.Response, csrfCookie)
}

func mandatoryMiddleware(r *requests.Request) {
	//for now only check security setting for CSRF inclusion
	//Set the default behavior to have CSRF protection enabled
	if config.Config.GetBool("enableCSRFProtection", true) {
		AddCSRFToken(r)
	}
}

func init() {
	pipeline = make([]RequestHandler, 0)
	//Add in the security CSRF check by default
	pipeline = append(pipeline, mandatoryMiddleware)
}
