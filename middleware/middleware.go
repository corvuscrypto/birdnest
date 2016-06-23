package middleware

import "github.com/corvuscrypto/birdnest/requests"

var pipeline []RequestHandler

//RequestHandler type is the function signature required in
//order to register a new middleware request handler
type RequestHandler func(*requests.Request)

//RegisterRequestHandler registers middleware. Middleware request handlers are utilized in FIFO order.
func RegisterRequestHandler(rh RequestHandler) {
	pipeline = append(pipeline, rh)
}
