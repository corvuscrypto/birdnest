package config

import "github.com/corvuscrypto/birdnest/requests"

//RegisterMiddleware registers a request middleware handler. These middle-tier handlers are executed in sequence.
func RegisterMiddleware(f func(*requests.Request)) {
	pipeline = append(pipeline, f)
}
