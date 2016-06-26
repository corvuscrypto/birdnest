package server

import "github.com/corvuscrypto/birdnest/requests"

var pipeline []RequestHandler

//RequestHandler type is the function signature required in
//order to register a new middleware request handler
type RequestHandler func(*requests.Request)

//RegisterMiddleware registers middleware. Middleware request handlers are utilized in FIFO order.
func RegisterMiddleware(rh ...RequestHandler) {
	if rh == nil {
		return
	}
	pipeline = append(pipeline, rh...)
}

func applyMiddleware(r *requests.Request) {
	for _, f := range pipeline {
		f(r)
	}
}
