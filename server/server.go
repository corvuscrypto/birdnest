package server

import (
	"net/http"
	"strconv"

	"github.com/corvuscrypto/birdnest/config"
)

//NewServer initializes a server to listen for http requests and returns it for user-initiated listening.
func NewServer(router *Router) *http.Server {
	if router == nil {
		router = NewRouter(nil)
	}
	serverPort := config.Config.GetInt("serverPort", 8080)
	server := new(http.Server)
	server.Handler = router
	server.Addr = ":" + strconv.Itoa(serverPort)
	return server
}
