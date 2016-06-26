package server

import (
	"net/http"
	"testing"

	"github.com/corvuscrypto/birdnest/requests"
)

func TestServer(T *testing.T) {
	var success bool
	//start a server with a single route
	router := NewRouter(nil)
	router.GET("/", func(r *requests.Request) { success = true })
	server := NewServer(nil)
	go func() {
		server.ListenAndServe()
		http.Get(server.Addr + "/")
		if !success {
			T.Errorf("The request was not served as expected!")
		}
	}()
}
