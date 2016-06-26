package server

import (
	"net/http"
	"sync"
	"testing"

	"github.com/corvuscrypto/birdnest/requests"
)

func TestServer(T *testing.T) {
	var success bool
	//start a server with a single route
	router := NewRouter(nil)
	router.GET("/", func(r *requests.Request) { success = true })
	server := NewServer(router)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		go server.ListenAndServe()
		http.Get("http://localhost" + server.Addr + "/")
		if !success {
			T.Errorf("The request was not served as expected!")
		}
		wg.Done()
	}()
	wg.Wait()

	//now just test server creation in nil value for parameter
	server = NewServer(nil)
}
