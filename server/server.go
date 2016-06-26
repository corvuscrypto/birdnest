package server

import (
	"net/http"
	"strconv"

	"github.com/corvuscrypto/birdnest/config"
)

//Serve initializes a server to listen for http requests. It automatically handles transformation of an http.Request to
//the birdnest Request format.
func Serve(router *Router) {
	if router == nil {
		router = NewRouter(nil)
	}
	serverPort := config.Config.GetInt("serverPort")
	http.ListenAndServe(":"+strconv.Itoa(serverPort), router)
}
